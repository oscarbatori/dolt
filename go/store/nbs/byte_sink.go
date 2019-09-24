package nbs

import (
	"errors"
	"github.com/google/uuid"
	"github.com/liquidata-inc/dolt/go/libraries/utils/iohelp"
	"github.com/liquidata-inc/dolt/go/store/atomicerr"
	"io"
	"os"
	"path/filepath"
	"sync"
)

func flushSinkToFile(sink ByteSink, path string) (err error) {
	var f *os.File
	f, err = os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)

	if err != nil {
		return err
	}

	defer func() {
		closeErr := f.Close()

		if err == nil {
			err = closeErr
		}
	}()

	err = sink.Flush(f)
	return err
}

// A ByteSink is an interface for writing bytes which can later be flushed to a writer
type ByteSink interface {
	io.Writer

	// Flush writes all the data that was written to the ByteSink to the supplied writer
	Flush(wr io.Writer) error

	// FlushToFile writes all the data that was written to the ByteSink to a file at the given path
	FlushToFile(path string) error
}

// ErrBuffFull used by the FixedBufferSink when the data written is larger than the buffer allocated.
var ErrBufferFull = errors.New("buffer full")

// FixedBufferByteSink is a ByteSink implementation with a buffer whose size will not change.  Writing more
// data than the fixed buffer can hold will result in an error
type FixedBufferByteSink struct {
	buff []byte
	pos  uint64
}

// NewFixedBufferTableSink creates a FixedBufferTableSink which will use the supplied buffer
func NewFixedBufferTableSink(buff []byte) *FixedBufferByteSink {
	if len(buff) == 0 {
		panic("must provide a buffer")
	}

	return &FixedBufferByteSink{buff: buff}
}

// Write writes a byte array to the sink.
func (sink *FixedBufferByteSink) Write(src []byte) (int, error) {
	dest := sink.buff[sink.pos:]
	destLen := len(dest)
	srcLen := len(src)

	if destLen < srcLen {
		return 0, ErrBufferFull
	}

	copy(dest, src)

	sink.pos += uint64(srcLen)
	return srcLen, nil
}

// Flush writes all the data that was written to the ByteSink to the supplied writer
func (sink *FixedBufferByteSink) Flush(wr io.Writer) error {
	return iohelp.WriteAll(wr, sink.buff[:sink.pos])
}

func (sink *FixedBufferByteSink) FlushToFile(path string) (err error) {
	return flushSinkToFile(sink, path)
}

// BlockBufferByteSink allocates blocks of data with a given block size to store the bytes written to the sink. New
// blocks are allocated as needed in order to handle all the data of the Write calls.
type BlockBufferByteSink struct {
	blockSize int
	pos       uint64
	blocks    [][]byte
}

// NewBlockBufferTableSink creates a BlockBufferByteSink with the provided block size.
func NewBlockBufferTableSink(blockSize int) *BlockBufferByteSink {
	block := make([]byte, 0, blockSize)
	return &BlockBufferByteSink{blockSize, 0, [][]byte{block}}
}

// Write writes a byte array to the sink.
func (sink *BlockBufferByteSink) Write(src []byte) (int, error) {
	srcLen := len(src)
	currBlockIdx := len(sink.blocks) - 1
	currBlock := sink.blocks[currBlockIdx]
	remaining := cap(currBlock) - len(currBlock)

	if remaining >= srcLen {
		currBlock = append(currBlock, src...)
		sink.blocks[currBlockIdx] = currBlock
	} else {
		if remaining > 0 {
			currBlock = append(currBlock, src[:remaining]...)
			sink.blocks[currBlockIdx] = currBlock
		}

		newBlock := make([]byte, 0, sink.blockSize)
		newBlock = append(newBlock, src[remaining:]...)
		sink.blocks = append(sink.blocks, newBlock)
	}

	sink.pos += uint64(srcLen)
	return srcLen, nil
}

// Flush writes all the data that was written to the ByteSink to the supplied writer
func (sink *BlockBufferByteSink) Flush(wr io.Writer) (err error) {
	return iohelp.WriteAll(wr, sink.blocks...)
}

func (sink *BlockBufferByteSink) FlushToFile(path string) (err error) {
	return flushSinkToFile(sink, path)
}

type BufferedFileByteSink struct {
	blockSize 		int
	pos       		uint64
	currentBlock    []byte

	writeCh 		chan []byte
	ae              *atomicerr.AtomicError
	wg				*sync.WaitGroup

	wr				io.WriteCloser
	path			string
}

func NewBufferedFileByteSink(blockSize, chBufferSize int) (*BufferedFileByteSink, error) {
	path := filepath.Join(os.TempDir(), uuid.New().String() + ".tf")
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)

	if err != nil {
		return nil, err
	}

	sink := &BufferedFileByteSink{
		blockSize:    blockSize,
		currentBlock: make([]byte, blockSize),
		writeCh:      make(chan []byte, chBufferSize),
		ae:           atomicerr.New(),
		wg:           &sync.WaitGroup{},
		wr:			  f,
		path:		  path,
	}

	sink.wg.Add(1)
	go func() {
		defer sink.wg.Done()
		sink.backgroundWrite()
	}()

	return sink, nil
}

// Write writes a byte array to the sink.
func (sink *BufferedFileByteSink) Write(src []byte) (int, error) {
	srcLen := len(src)
	remaining := cap(sink.currentBlock) - len(sink.currentBlock)

	if remaining >= srcLen {
		sink.currentBlock = append(sink.currentBlock, src...)

		if remaining == srcLen {
			sink.writeCh <- sink.currentBlock
			sink.currentBlock = nil
		}
	} else {
		if remaining > 0 {
			sink.currentBlock = append(sink.currentBlock, src[:remaining]...)
			sink.writeCh <- sink.currentBlock
		}

		newBlock := make([]byte, 0, sink.blockSize)
		newBlock = append(newBlock, src[remaining:]...)
		sink.currentBlock = newBlock
	}

	sink.pos += uint64(srcLen)
	return srcLen, nil
}

func (sink *BufferedFileByteSink) backgroundWrite() {
	var err error
	for buff := range sink.writeCh {
		if err != nil {
			continue // drain
		}

		err = iohelp.WriteAll(sink.wr, buff)
		sink.ae.SetIfError(err)
	}

	err = sink.wr.Close()
	sink.ae.SetIfError(err)
}

// Flush writes all the data that was written to the ByteSink to the supplied writer
func (sink *BufferedFileByteSink) Flush(wr io.Writer) (err error) {
	toWrite := len(sink.currentBlock)
	if toWrite > 0 {
		sink.writeCh <- sink.currentBlock[:toWrite]
	}

	close(sink.writeCh)
	sink.wg.Wait()

	if err := sink.ae.Get(); err != nil {
		return err
	}

	var f *os.File
	f, err = os.Open(sink.path)

	if err != nil {
		return err
	}

	defer func() {
		closeErr := f.Close()

		if err == nil {
			err = closeErr
		}
	}()

	_, err = io.Copy(wr, f)

	return err
}

func (sink *BufferedFileByteSink) FlushToFile(path string) (err error) {
	toWrite := len(sink.currentBlock)
	if toWrite > 0 {
		sink.writeCh <- sink.currentBlock[:toWrite]
	}

	close(sink.writeCh)
	sink.wg.Wait()

	if err := sink.ae.Get(); err != nil {
		return err
	}

	return os.Rename(sink.path, path)
}
