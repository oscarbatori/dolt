// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dolt/services/eventsapi/v1alpha1/client_event.proto

package eventsapi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Platform int32

const (
	Platform_PLATFORM_UNSPECIFIED Platform = 0
	Platform_LINUX                Platform = 1
	Platform_WINDOWS              Platform = 2
	Platform_DARWIN               Platform = 3
)

var Platform_name = map[int32]string{
	0: "PLATFORM_UNSPECIFIED",
	1: "LINUX",
	2: "WINDOWS",
	3: "DARWIN",
}

var Platform_value = map[string]int32{
	"PLATFORM_UNSPECIFIED": 0,
	"LINUX":                1,
	"WINDOWS":              2,
	"DARWIN":               3,
}

func (x Platform) String() string {
	return proto.EnumName(Platform_name, int32(x))
}

func (Platform) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2059ca7f76f8dca6, []int{0}
}

type ClientEventType int32

const (
	ClientEventType_TYPE_UNSPECIFIED                 ClientEventType = 0
	ClientEventType_INIT                             ClientEventType = 1
	ClientEventType_STATUS                           ClientEventType = 2
	ClientEventType_ADD                              ClientEventType = 3
	ClientEventType_RESET                            ClientEventType = 4
	ClientEventType_COMMIT                           ClientEventType = 5
	ClientEventType_SQL                              ClientEventType = 6
	ClientEventType_SQL_SERVER                       ClientEventType = 7
	ClientEventType_LOG                              ClientEventType = 8
	ClientEventType_DIFF                             ClientEventType = 9
	ClientEventType_MERGE                            ClientEventType = 10
	ClientEventType_BRANCH                           ClientEventType = 11
	ClientEventType_CHECKOUT                         ClientEventType = 12
	ClientEventType_REMOTE                           ClientEventType = 13
	ClientEventType_PUSH                             ClientEventType = 14
	ClientEventType_PULL                             ClientEventType = 15
	ClientEventType_FETCH                            ClientEventType = 16
	ClientEventType_CLONE                            ClientEventType = 17
	ClientEventType_LOGIN                            ClientEventType = 18
	ClientEventType_VERSION                          ClientEventType = 19
	ClientEventType_CONFIG                           ClientEventType = 20
	ClientEventType_LS                               ClientEventType = 21
	ClientEventType_SCHEMA                           ClientEventType = 22
	ClientEventType_TABLE_IMPORT                     ClientEventType = 23
	ClientEventType_TABLE_EXPORT                     ClientEventType = 24
	ClientEventType_TABLE_CREATE                     ClientEventType = 25
	ClientEventType_TABLE_RM                         ClientEventType = 26
	ClientEventType_TABLE_MV                         ClientEventType = 27
	ClientEventType_TABLE_CP                         ClientEventType = 28
	ClientEventType_TABLE_SELECT                     ClientEventType = 29
	ClientEventType_TABLE_PUT_ROW                    ClientEventType = 30
	ClientEventType_TABLE_RM_ROW                     ClientEventType = 31
	ClientEventType_CREDS_NEW                        ClientEventType = 32
	ClientEventType_CREDS_RM                         ClientEventType = 33
	ClientEventType_CREDS_LS                         ClientEventType = 34
	ClientEventType_CONF_CAT                         ClientEventType = 35
	ClientEventType_CONF_RESOLVE                     ClientEventType = 36
	ClientEventType_REMOTEAPI_GET_REPO_METADATA      ClientEventType = 37
	ClientEventType_REMOTEAPI_HAS_CHUNKS             ClientEventType = 38
	ClientEventType_REMOTEAPI_GET_DOWNLOAD_LOCATIONS ClientEventType = 39
	ClientEventType_REMOTEAPI_GET_UPLOAD_LOCATIONS   ClientEventType = 40
	ClientEventType_REMOTEAPI_REBASE                 ClientEventType = 41
	ClientEventType_REMOTEAPI_ROOT                   ClientEventType = 42
	ClientEventType_REMOTEAPI_COMMIT                 ClientEventType = 43
	ClientEventType_REMOTEAPI_LIST_TABLE_FILES       ClientEventType = 44
)

var ClientEventType_name = map[int32]string{
	0:  "TYPE_UNSPECIFIED",
	1:  "INIT",
	2:  "STATUS",
	3:  "ADD",
	4:  "RESET",
	5:  "COMMIT",
	6:  "SQL",
	7:  "SQL_SERVER",
	8:  "LOG",
	9:  "DIFF",
	10: "MERGE",
	11: "BRANCH",
	12: "CHECKOUT",
	13: "REMOTE",
	14: "PUSH",
	15: "PULL",
	16: "FETCH",
	17: "CLONE",
	18: "LOGIN",
	19: "VERSION",
	20: "CONFIG",
	21: "LS",
	22: "SCHEMA",
	23: "TABLE_IMPORT",
	24: "TABLE_EXPORT",
	25: "TABLE_CREATE",
	26: "TABLE_RM",
	27: "TABLE_MV",
	28: "TABLE_CP",
	29: "TABLE_SELECT",
	30: "TABLE_PUT_ROW",
	31: "TABLE_RM_ROW",
	32: "CREDS_NEW",
	33: "CREDS_RM",
	34: "CREDS_LS",
	35: "CONF_CAT",
	36: "CONF_RESOLVE",
	37: "REMOTEAPI_GET_REPO_METADATA",
	38: "REMOTEAPI_HAS_CHUNKS",
	39: "REMOTEAPI_GET_DOWNLOAD_LOCATIONS",
	40: "REMOTEAPI_GET_UPLOAD_LOCATIONS",
	41: "REMOTEAPI_REBASE",
	42: "REMOTEAPI_ROOT",
	43: "REMOTEAPI_COMMIT",
	44: "REMOTEAPI_LIST_TABLE_FILES",
}

var ClientEventType_value = map[string]int32{
	"TYPE_UNSPECIFIED":                 0,
	"INIT":                             1,
	"STATUS":                           2,
	"ADD":                              3,
	"RESET":                            4,
	"COMMIT":                           5,
	"SQL":                              6,
	"SQL_SERVER":                       7,
	"LOG":                              8,
	"DIFF":                             9,
	"MERGE":                            10,
	"BRANCH":                           11,
	"CHECKOUT":                         12,
	"REMOTE":                           13,
	"PUSH":                             14,
	"PULL":                             15,
	"FETCH":                            16,
	"CLONE":                            17,
	"LOGIN":                            18,
	"VERSION":                          19,
	"CONFIG":                           20,
	"LS":                               21,
	"SCHEMA":                           22,
	"TABLE_IMPORT":                     23,
	"TABLE_EXPORT":                     24,
	"TABLE_CREATE":                     25,
	"TABLE_RM":                         26,
	"TABLE_MV":                         27,
	"TABLE_CP":                         28,
	"TABLE_SELECT":                     29,
	"TABLE_PUT_ROW":                    30,
	"TABLE_RM_ROW":                     31,
	"CREDS_NEW":                        32,
	"CREDS_RM":                         33,
	"CREDS_LS":                         34,
	"CONF_CAT":                         35,
	"CONF_RESOLVE":                     36,
	"REMOTEAPI_GET_REPO_METADATA":      37,
	"REMOTEAPI_HAS_CHUNKS":             38,
	"REMOTEAPI_GET_DOWNLOAD_LOCATIONS": 39,
	"REMOTEAPI_GET_UPLOAD_LOCATIONS":   40,
	"REMOTEAPI_REBASE":                 41,
	"REMOTEAPI_ROOT":                   42,
	"REMOTEAPI_COMMIT":                 43,
	"REMOTEAPI_LIST_TABLE_FILES":       44,
}

func (x ClientEventType) String() string {
	return proto.EnumName(ClientEventType_name, int32(x))
}

func (ClientEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2059ca7f76f8dca6, []int{1}
}

type MetricID int32

const (
	MetricID_METRIC_UNSPECIFIED  MetricID = 0
	MetricID_BYTES_DOWNLOADED    MetricID = 1
	MetricID_BYTES_UPLOADED      MetricID = 2
	MetricID_DOWNLOAD_MS_ELAPSED MetricID = 3
	MetricID_UPLOAD_MS_ELAPSED   MetricID = 4
	MetricID_REMOTEAPI_RPC_ERROR MetricID = 5
)

var MetricID_name = map[int32]string{
	0: "METRIC_UNSPECIFIED",
	1: "BYTES_DOWNLOADED",
	2: "BYTES_UPLOADED",
	3: "DOWNLOAD_MS_ELAPSED",
	4: "UPLOAD_MS_ELAPSED",
	5: "REMOTEAPI_RPC_ERROR",
}

var MetricID_value = map[string]int32{
	"METRIC_UNSPECIFIED":  0,
	"BYTES_DOWNLOADED":    1,
	"BYTES_UPLOADED":      2,
	"DOWNLOAD_MS_ELAPSED": 3,
	"UPLOAD_MS_ELAPSED":   4,
	"REMOTEAPI_RPC_ERROR": 5,
}

func (x MetricID) String() string {
	return proto.EnumName(MetricID_name, int32(x))
}

func (MetricID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2059ca7f76f8dca6, []int{2}
}

type AttributeID int32

const (
	AttributeID_ATTRIBUTE_UNSPECIFIED AttributeID = 0
	AttributeID_ACTIVE_REMOTE_URL     AttributeID = 1
)

var AttributeID_name = map[int32]string{
	0: "ATTRIBUTE_UNSPECIFIED",
	1: "ACTIVE_REMOTE_URL",
}

var AttributeID_value = map[string]int32{
	"ATTRIBUTE_UNSPECIFIED": 0,
	"ACTIVE_REMOTE_URL":     1,
}

func (x AttributeID) String() string {
	return proto.EnumName(AttributeID_name, int32(x))
}

func (AttributeID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2059ca7f76f8dca6, []int{3}
}

type ClientEventAttribute struct {
	Id                   AttributeID `protobuf:"varint,1,opt,name=id,proto3,enum=dolt.services.eventsapi.v1alpha1.AttributeID" json:"id,omitempty"`
	Value                string      `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ClientEventAttribute) Reset()         { *m = ClientEventAttribute{} }
func (m *ClientEventAttribute) String() string { return proto.CompactTextString(m) }
func (*ClientEventAttribute) ProtoMessage()    {}
func (*ClientEventAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_2059ca7f76f8dca6, []int{0}
}

func (m *ClientEventAttribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientEventAttribute.Unmarshal(m, b)
}
func (m *ClientEventAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientEventAttribute.Marshal(b, m, deterministic)
}
func (m *ClientEventAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientEventAttribute.Merge(m, src)
}
func (m *ClientEventAttribute) XXX_Size() int {
	return xxx_messageInfo_ClientEventAttribute.Size(m)
}
func (m *ClientEventAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientEventAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_ClientEventAttribute proto.InternalMessageInfo

func (m *ClientEventAttribute) GetId() AttributeID {
	if m != nil {
		return m.Id
	}
	return AttributeID_ATTRIBUTE_UNSPECIFIED
}

func (m *ClientEventAttribute) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ClientEventMetric struct {
	// Types that are valid to be assigned to MetricOneof:
	//	*ClientEventMetric_Duration
	//	*ClientEventMetric_Count
	MetricOneof          isClientEventMetric_MetricOneof `protobuf_oneof:"metric_oneof"`
	MetricId             MetricID                        `protobuf:"varint,100,opt,name=metric_id,json=metricId,proto3,enum=dolt.services.eventsapi.v1alpha1.MetricID" json:"metric_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ClientEventMetric) Reset()         { *m = ClientEventMetric{} }
func (m *ClientEventMetric) String() string { return proto.CompactTextString(m) }
func (*ClientEventMetric) ProtoMessage()    {}
func (*ClientEventMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_2059ca7f76f8dca6, []int{1}
}

func (m *ClientEventMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientEventMetric.Unmarshal(m, b)
}
func (m *ClientEventMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientEventMetric.Marshal(b, m, deterministic)
}
func (m *ClientEventMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientEventMetric.Merge(m, src)
}
func (m *ClientEventMetric) XXX_Size() int {
	return xxx_messageInfo_ClientEventMetric.Size(m)
}
func (m *ClientEventMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientEventMetric.DiscardUnknown(m)
}

var xxx_messageInfo_ClientEventMetric proto.InternalMessageInfo

type isClientEventMetric_MetricOneof interface {
	isClientEventMetric_MetricOneof()
}

type ClientEventMetric_Duration struct {
	Duration *duration.Duration `protobuf:"bytes,1,opt,name=duration,proto3,oneof"`
}

type ClientEventMetric_Count struct {
	Count int32 `protobuf:"varint,2,opt,name=count,proto3,oneof"`
}

func (*ClientEventMetric_Duration) isClientEventMetric_MetricOneof() {}

func (*ClientEventMetric_Count) isClientEventMetric_MetricOneof() {}

func (m *ClientEventMetric) GetMetricOneof() isClientEventMetric_MetricOneof {
	if m != nil {
		return m.MetricOneof
	}
	return nil
}

func (m *ClientEventMetric) GetDuration() *duration.Duration {
	if x, ok := m.GetMetricOneof().(*ClientEventMetric_Duration); ok {
		return x.Duration
	}
	return nil
}

func (m *ClientEventMetric) GetCount() int32 {
	if x, ok := m.GetMetricOneof().(*ClientEventMetric_Count); ok {
		return x.Count
	}
	return 0
}

func (m *ClientEventMetric) GetMetricId() MetricID {
	if m != nil {
		return m.MetricId
	}
	return MetricID_METRIC_UNSPECIFIED
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ClientEventMetric) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ClientEventMetric_Duration)(nil),
		(*ClientEventMetric_Count)(nil),
	}
}

type ClientEvent struct {
	Id                   string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StartTime            *timestamp.Timestamp    `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp    `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Type                 ClientEventType         `protobuf:"varint,4,opt,name=type,proto3,enum=dolt.services.eventsapi.v1alpha1.ClientEventType" json:"type,omitempty"`
	Attributes           []*ClientEventAttribute `protobuf:"bytes,5,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Metrics              []*ClientEventMetric    `protobuf:"bytes,6,rep,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClientEvent) Reset()         { *m = ClientEvent{} }
func (m *ClientEvent) String() string { return proto.CompactTextString(m) }
func (*ClientEvent) ProtoMessage()    {}
func (*ClientEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_2059ca7f76f8dca6, []int{2}
}

func (m *ClientEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientEvent.Unmarshal(m, b)
}
func (m *ClientEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientEvent.Marshal(b, m, deterministic)
}
func (m *ClientEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientEvent.Merge(m, src)
}
func (m *ClientEvent) XXX_Size() int {
	return xxx_messageInfo_ClientEvent.Size(m)
}
func (m *ClientEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ClientEvent proto.InternalMessageInfo

func (m *ClientEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ClientEvent) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ClientEvent) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *ClientEvent) GetType() ClientEventType {
	if m != nil {
		return m.Type
	}
	return ClientEventType_TYPE_UNSPECIFIED
}

func (m *ClientEvent) GetAttributes() []*ClientEventAttribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *ClientEvent) GetMetrics() []*ClientEventMetric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type LogEventsRequest struct {
	MachineId            string         `protobuf:"bytes,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	Extra                string         `protobuf:"bytes,2,opt,name=extra,proto3" json:"extra,omitempty"`
	Version              string         `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Platform             Platform       `protobuf:"varint,4,opt,name=platform,proto3,enum=dolt.services.eventsapi.v1alpha1.Platform" json:"platform,omitempty"`
	Events               []*ClientEvent `protobuf:"bytes,5,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *LogEventsRequest) Reset()         { *m = LogEventsRequest{} }
func (m *LogEventsRequest) String() string { return proto.CompactTextString(m) }
func (*LogEventsRequest) ProtoMessage()    {}
func (*LogEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2059ca7f76f8dca6, []int{3}
}

func (m *LogEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEventsRequest.Unmarshal(m, b)
}
func (m *LogEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEventsRequest.Marshal(b, m, deterministic)
}
func (m *LogEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEventsRequest.Merge(m, src)
}
func (m *LogEventsRequest) XXX_Size() int {
	return xxx_messageInfo_LogEventsRequest.Size(m)
}
func (m *LogEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogEventsRequest proto.InternalMessageInfo

func (m *LogEventsRequest) GetMachineId() string {
	if m != nil {
		return m.MachineId
	}
	return ""
}

func (m *LogEventsRequest) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

func (m *LogEventsRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *LogEventsRequest) GetPlatform() Platform {
	if m != nil {
		return m.Platform
	}
	return Platform_PLATFORM_UNSPECIFIED
}

func (m *LogEventsRequest) GetEvents() []*ClientEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

type LogEventsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogEventsResponse) Reset()         { *m = LogEventsResponse{} }
func (m *LogEventsResponse) String() string { return proto.CompactTextString(m) }
func (*LogEventsResponse) ProtoMessage()    {}
func (*LogEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2059ca7f76f8dca6, []int{4}
}

func (m *LogEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEventsResponse.Unmarshal(m, b)
}
func (m *LogEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEventsResponse.Marshal(b, m, deterministic)
}
func (m *LogEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEventsResponse.Merge(m, src)
}
func (m *LogEventsResponse) XXX_Size() int {
	return xxx_messageInfo_LogEventsResponse.Size(m)
}
func (m *LogEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogEventsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("dolt.services.eventsapi.v1alpha1.Platform", Platform_name, Platform_value)
	proto.RegisterEnum("dolt.services.eventsapi.v1alpha1.ClientEventType", ClientEventType_name, ClientEventType_value)
	proto.RegisterEnum("dolt.services.eventsapi.v1alpha1.MetricID", MetricID_name, MetricID_value)
	proto.RegisterEnum("dolt.services.eventsapi.v1alpha1.AttributeID", AttributeID_name, AttributeID_value)
	proto.RegisterType((*ClientEventAttribute)(nil), "dolt.services.eventsapi.v1alpha1.ClientEventAttribute")
	proto.RegisterType((*ClientEventMetric)(nil), "dolt.services.eventsapi.v1alpha1.ClientEventMetric")
	proto.RegisterType((*ClientEvent)(nil), "dolt.services.eventsapi.v1alpha1.ClientEvent")
	proto.RegisterType((*LogEventsRequest)(nil), "dolt.services.eventsapi.v1alpha1.LogEventsRequest")
	proto.RegisterType((*LogEventsResponse)(nil), "dolt.services.eventsapi.v1alpha1.LogEventsResponse")
}

func init() {
	proto.RegisterFile("dolt/services/eventsapi/v1alpha1/client_event.proto", fileDescriptor_2059ca7f76f8dca6)
}

var fileDescriptor_2059ca7f76f8dca6 = []byte{
	// 1145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xdd, 0x72, 0xdb, 0x44,
	0x14, 0xae, 0x9d, 0x38, 0xb1, 0x8f, 0xd3, 0xf4, 0x64, 0x93, 0xb4, 0x6e, 0x4a, 0xd3, 0x60, 0x0a,
	0x84, 0x40, 0xed, 0x69, 0x32, 0xc0, 0x30, 0x0c, 0xc3, 0xc8, 0xd2, 0x3a, 0xd6, 0x54, 0x7f, 0xdd,
	0x5d, 0xc7, 0x2d, 0x37, 0x1a, 0xc5, 0x56, 0x53, 0x0d, 0x8e, 0xe5, 0x5a, 0x72, 0x86, 0x3e, 0x03,
	0xd7, 0x3c, 0x0e, 0x4f, 0xc3, 0x0b, 0xf0, 0x02, 0xcc, 0x30, 0xbb, 0xf2, 0x5f, 0xcc, 0x85, 0xc9,
	0x9d, 0xcf, 0x39, 0xdf, 0xf7, 0xed, 0xf9, 0x8e, 0x8e, 0xd6, 0x82, 0xb3, 0x5e, 0xdc, 0x4f, 0xeb,
	0x49, 0x38, 0xba, 0x89, 0xba, 0x61, 0x52, 0x0f, 0x6f, 0xc2, 0x41, 0x9a, 0x04, 0xc3, 0xa8, 0x7e,
	0xf3, 0x32, 0xe8, 0x0f, 0xdf, 0x07, 0x2f, 0xeb, 0xdd, 0x7e, 0x14, 0x0e, 0x52, 0x5f, 0x55, 0x6a,
	0xc3, 0x51, 0x9c, 0xc6, 0xe4, 0x48, 0x92, 0x6a, 0x53, 0x52, 0x6d, 0x46, 0xaa, 0x4d, 0x49, 0x07,
	0x87, 0x57, 0x71, 0x7c, 0xd5, 0x0f, 0xeb, 0x0a, 0x7f, 0x39, 0x7e, 0x57, 0xef, 0x8d, 0x47, 0x41,
	0x1a, 0xc5, 0x83, 0x4c, 0xe1, 0xe0, 0xd9, 0x72, 0x3d, 0x8d, 0xae, 0xc3, 0x24, 0x0d, 0xae, 0x87,
	0x19, 0xa0, 0xfa, 0x2b, 0xec, 0xe9, 0xea, 0x60, 0x2a, 0xc5, 0xb5, 0x34, 0x1d, 0x45, 0x97, 0xe3,
	0x34, 0x24, 0x3f, 0x41, 0x3e, 0xea, 0x55, 0x72, 0x47, 0xb9, 0xe3, 0xed, 0xd3, 0x17, 0xb5, 0x55,
	0x7d, 0xd4, 0x66, 0x44, 0xd3, 0x60, 0xf9, 0xa8, 0x47, 0xf6, 0xa0, 0x70, 0x13, 0xf4, 0xc7, 0x61,
	0x25, 0x7f, 0x94, 0x3b, 0x2e, 0xb1, 0x2c, 0xa8, 0xfe, 0x99, 0x83, 0x9d, 0x85, 0xd3, 0xec, 0x30,
	0x1d, 0x45, 0x5d, 0xf2, 0x3d, 0x14, 0xa7, 0x5d, 0xab, 0x03, 0xcb, 0xa7, 0x8f, 0x6b, 0x59, 0xdb,
	0xb5, 0x69, 0xdb, 0x35, 0x63, 0x02, 0x68, 0xdd, 0x63, 0x33, 0x30, 0x79, 0x08, 0x85, 0x6e, 0x3c,
	0x1e, 0xa4, 0xea, 0x90, 0x42, 0xeb, 0x1e, 0xcb, 0x42, 0x72, 0x0e, 0xa5, 0x6b, 0x25, 0xed, 0x47,
	0xbd, 0x4a, 0x4f, 0x59, 0x38, 0x59, 0x6d, 0x21, 0xeb, 0xc6, 0x34, 0x58, 0x31, 0x23, 0x9b, 0xbd,
	0xc6, 0x36, 0x6c, 0x4d, 0x84, 0xe2, 0x41, 0x18, 0xbf, 0xab, 0xfe, 0x93, 0x87, 0xf2, 0x42, 0xff,
	0x64, 0x7b, 0x36, 0xa4, 0x92, 0x72, 0xfd, 0x03, 0x40, 0x92, 0x06, 0xa3, 0xd4, 0x97, 0x53, 0x56,
	0x5d, 0x95, 0x4f, 0x0f, 0xfe, 0xe3, 0x45, 0x4c, 0x1f, 0x01, 0x2b, 0x29, 0xb4, 0x8c, 0xc9, 0xb7,
	0x50, 0x0c, 0x07, 0xbd, 0x8c, 0xb8, 0xb6, 0x92, 0xb8, 0x19, 0x0e, 0x7a, 0x8a, 0x46, 0x61, 0x3d,
	0xfd, 0x38, 0x0c, 0x2b, 0xeb, 0xca, 0xe5, 0xcb, 0xd5, 0x2e, 0x17, 0xda, 0x17, 0x1f, 0x87, 0x21,
	0x53, 0x74, 0x72, 0x01, 0x10, 0x4c, 0x9f, 0x60, 0x52, 0x29, 0x1c, 0xad, 0x1d, 0x97, 0x4f, 0xbf,
	0xbb, 0x93, 0xd8, 0x6c, 0x01, 0xd8, 0x82, 0x12, 0xb1, 0x61, 0x33, 0x1b, 0x60, 0x52, 0xd9, 0x50,
	0xa2, 0x67, 0x77, 0x12, 0xcd, 0x1e, 0x09, 0x9b, 0x6a, 0x54, 0xff, 0xce, 0x01, 0x5a, 0xf1, 0x95,
	0xaa, 0x25, 0x2c, 0xfc, 0x30, 0x0e, 0x93, 0x94, 0x3c, 0x05, 0xb8, 0x0e, 0xba, 0xef, 0xa3, 0x41,
	0xe8, 0xcf, 0x1e, 0x46, 0x69, 0x92, 0x31, 0xd5, 0x26, 0x86, 0xbf, 0xa5, 0xa3, 0x60, 0xba, 0x89,
	0x2a, 0x20, 0x15, 0xd8, 0xbc, 0x09, 0x47, 0x89, 0x5c, 0xb9, 0x35, 0x95, 0x9f, 0x86, 0xa4, 0x09,
	0xc5, 0x61, 0x3f, 0x48, 0xdf, 0xc5, 0xa3, 0xeb, 0xc9, 0x54, 0xff, 0xc7, 0xee, 0x78, 0x13, 0x06,
	0x9b, 0x71, 0x09, 0x85, 0x8d, 0x0c, 0x38, 0x19, 0xe7, 0x8b, 0x3b, 0x39, 0x67, 0x13, 0x72, 0x75,
	0x17, 0x76, 0x16, 0x1c, 0x27, 0xc3, 0x78, 0x90, 0x84, 0x27, 0x2d, 0x28, 0x4e, 0x4f, 0x24, 0x15,
	0xd8, 0xf3, 0x2c, 0x4d, 0x34, 0x5d, 0x66, 0xfb, 0x6d, 0x87, 0x7b, 0x54, 0x37, 0x9b, 0x26, 0x35,
	0xf0, 0x1e, 0x29, 0x41, 0xc1, 0x32, 0x9d, 0xf6, 0x1b, 0xcc, 0x91, 0x32, 0x6c, 0x76, 0x4c, 0xc7,
	0x70, 0x3b, 0x1c, 0xf3, 0x04, 0x60, 0xc3, 0xd0, 0x58, 0xc7, 0x74, 0x70, 0xed, 0xe4, 0xaf, 0x02,
	0x3c, 0x58, 0x5a, 0x09, 0xb2, 0x07, 0x28, 0xde, 0x7a, 0x74, 0x49, 0xad, 0x08, 0xeb, 0xa6, 0x63,
	0x0a, 0xcc, 0x49, 0x3e, 0x17, 0x9a, 0x68, 0x4b, 0xad, 0x4d, 0x58, 0xd3, 0x0c, 0x03, 0xd7, 0xe4,
	0x61, 0x8c, 0x72, 0x2a, 0x70, 0x5d, 0xd6, 0x75, 0xd7, 0xb6, 0x4d, 0x81, 0x05, 0x59, 0xe7, 0xaf,
	0x2d, 0xdc, 0x20, 0xdb, 0x00, 0xfc, 0xb5, 0xe5, 0x73, 0xca, 0x2e, 0x28, 0xc3, 0x4d, 0x59, 0xb0,
	0xdc, 0x73, 0x2c, 0x4a, 0x5d, 0xc3, 0x6c, 0x36, 0xb1, 0x24, 0x25, 0x6c, 0xca, 0xce, 0x29, 0x82,
	0x94, 0x68, 0x30, 0xcd, 0xd1, 0x5b, 0x58, 0x26, 0x5b, 0x50, 0xd4, 0x5b, 0x54, 0x7f, 0xe5, 0xb6,
	0x05, 0x6e, 0xc9, 0x0a, 0xa3, 0xb6, 0x2b, 0x28, 0xde, 0x97, 0x54, 0xaf, 0xcd, 0x5b, 0xb8, 0x9d,
	0xfd, 0xb2, 0x2c, 0x7c, 0x20, 0x45, 0x9a, 0x54, 0xe8, 0x2d, 0x44, 0xf9, 0x53, 0xb7, 0x5c, 0x87,
	0xe2, 0x8e, 0x1a, 0x85, 0x7b, 0x6e, 0x3a, 0x48, 0xe4, 0x28, 0x2e, 0x28, 0xe3, 0xa6, 0xeb, 0xe0,
	0x6e, 0xd6, 0xaa, 0xd3, 0x34, 0xcf, 0x71, 0x8f, 0x6c, 0x40, 0xde, 0xe2, 0xb8, 0xaf, 0xec, 0xe9,
	0x2d, 0x6a, 0x6b, 0xf8, 0x90, 0x20, 0x6c, 0x09, 0xad, 0x61, 0x51, 0xdf, 0xb4, 0x3d, 0x97, 0x09,
	0x7c, 0x34, 0xcf, 0xd0, 0x37, 0x2a, 0x53, 0x99, 0x67, 0x74, 0x46, 0x35, 0x41, 0xf1, 0xb1, 0xec,
	0x38, 0xcb, 0x30, 0x1b, 0x0f, 0xe6, 0x91, 0x7d, 0x81, 0x4f, 0xe6, 0x91, 0xee, 0xe1, 0x27, 0x73,
	0x2e, 0xa7, 0x16, 0xd5, 0x05, 0x3e, 0x25, 0x3b, 0x70, 0x3f, 0xcb, 0x78, 0x6d, 0xe1, 0x33, 0xb7,
	0x83, 0x87, 0x73, 0x10, 0xb3, 0x55, 0xe6, 0x19, 0xb9, 0x0f, 0x25, 0x9d, 0x51, 0x83, 0xfb, 0x0e,
	0xed, 0xe0, 0x91, 0x9a, 0x90, 0x0a, 0x99, 0x8d, 0x9f, 0xce, 0x23, 0x8b, 0x63, 0x55, 0x45, 0xae,
	0xd3, 0xf4, 0x75, 0x4d, 0xe0, 0x67, 0x52, 0x4a, 0x45, 0x8c, 0x72, 0xd7, 0xba, 0xa0, 0xf8, 0x9c,
	0x3c, 0x83, 0x27, 0xd9, 0x3c, 0x35, 0xcf, 0xf4, 0xcf, 0xa9, 0xf0, 0x19, 0xf5, 0x5c, 0xdf, 0xa6,
	0x42, 0x33, 0x34, 0xa1, 0xe1, 0xe7, 0x72, 0xbf, 0xe6, 0x80, 0x96, 0xc6, 0x7d, 0xbd, 0xd5, 0x76,
	0x5e, 0x71, 0xfc, 0x82, 0x3c, 0x87, 0xa3, 0xdb, 0x54, 0xc3, 0xed, 0x38, 0x96, 0xab, 0x19, 0xbe,
	0xe5, 0xea, 0x9a, 0x30, 0x5d, 0x87, 0xe3, 0x97, 0xa4, 0x0a, 0x87, 0xb7, 0x51, 0x6d, 0x6f, 0x09,
	0x73, 0x2c, 0x37, 0x6e, 0x8e, 0x61, 0xb4, 0xa1, 0x71, 0x8a, 0x5f, 0x11, 0x02, 0xdb, 0x0b, 0x59,
	0xd7, 0x15, 0x78, 0x72, 0x1b, 0x39, 0xd9, 0xb2, 0xaf, 0xc9, 0x21, 0x1c, 0xcc, 0xb3, 0x96, 0xc9,
	0x85, 0x9f, 0x0d, 0xac, 0x69, 0x5a, 0x94, 0xe3, 0x37, 0x27, 0x7f, 0xe4, 0xa0, 0x38, 0xbd, 0xde,
	0xc9, 0x43, 0x20, 0x36, 0x15, 0xcc, 0xd4, 0x97, 0x16, 0x7c, 0x0f, 0xb0, 0xf1, 0x56, 0x50, 0x3e,
	0xb3, 0x41, 0x0d, 0xcc, 0xc9, 0x26, 0xb2, 0x6c, 0xd6, 0x36, 0x35, 0x30, 0x4f, 0x1e, 0xc1, 0xee,
	0xcc, 0xaa, 0xcd, 0x7d, 0x6a, 0x69, 0x1e, 0xa7, 0xf2, 0x25, 0xd8, 0x87, 0x9d, 0x89, 0xbb, 0x85,
	0xf4, 0xba, 0xc4, 0x2f, 0x18, 0xf1, 0x74, 0x9f, 0x32, 0xe6, 0x32, 0x2c, 0x9c, 0xfc, 0x0c, 0xe5,
	0x85, 0x3f, 0x4e, 0xf2, 0x18, 0xf6, 0x35, 0x21, 0x98, 0xd9, 0x68, 0x8b, 0xe5, 0xb7, 0x6f, 0x1f,
	0x76, 0x34, 0x5d, 0x98, 0x17, 0xd4, 0xcf, 0x94, 0xfc, 0x36, 0xb3, 0x30, 0x77, 0xfa, 0x7b, 0x0e,
	0x76, 0x17, 0x5e, 0xdf, 0x84, 0x67, 0xb7, 0x0b, 0x49, 0xa1, 0x34, 0xbb, 0x35, 0xc8, 0xe9, 0xea,
	0x9b, 0x67, 0xf9, 0x52, 0x3d, 0x38, 0xbb, 0x13, 0x27, 0xbb, 0x96, 0x1a, 0x9d, 0x5f, 0xda, 0x57,
	0x51, 0xfa, 0x7e, 0x7c, 0x59, 0xeb, 0xc6, 0xd7, 0xf5, 0x7e, 0xf4, 0x61, 0x1c, 0xf5, 0x82, 0x34,
	0x78, 0x11, 0x0d, 0xba, 0x75, 0xf5, 0xf9, 0x73, 0x15, 0xd7, 0xaf, 0xc2, 0x41, 0xf6, 0x31, 0x52,
	0x5f, 0xf5, 0x41, 0xf4, 0xe3, 0x2c, 0x75, 0xb9, 0xa1, 0x18, 0x67, 0xff, 0x06, 0x00, 0x00, 0xff,
	0xff, 0xa3, 0xf5, 0x0e, 0xaf, 0x45, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientEventsServiceClient is the client API for ClientEventsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientEventsServiceClient interface {
	LogEvents(ctx context.Context, in *LogEventsRequest, opts ...grpc.CallOption) (*LogEventsResponse, error)
}

type clientEventsServiceClient struct {
	cc *grpc.ClientConn
}

func NewClientEventsServiceClient(cc *grpc.ClientConn) ClientEventsServiceClient {
	return &clientEventsServiceClient{cc}
}

func (c *clientEventsServiceClient) LogEvents(ctx context.Context, in *LogEventsRequest, opts ...grpc.CallOption) (*LogEventsResponse, error) {
	out := new(LogEventsResponse)
	err := c.cc.Invoke(ctx, "/dolt.services.eventsapi.v1alpha1.ClientEventsService/LogEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientEventsServiceServer is the server API for ClientEventsService service.
type ClientEventsServiceServer interface {
	LogEvents(context.Context, *LogEventsRequest) (*LogEventsResponse, error)
}

// UnimplementedClientEventsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClientEventsServiceServer struct {
}

func (*UnimplementedClientEventsServiceServer) LogEvents(ctx context.Context, req *LogEventsRequest) (*LogEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogEvents not implemented")
}

func RegisterClientEventsServiceServer(s *grpc.Server, srv ClientEventsServiceServer) {
	s.RegisterService(&_ClientEventsService_serviceDesc, srv)
}

func _ClientEventsService_LogEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientEventsServiceServer).LogEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dolt.services.eventsapi.v1alpha1.ClientEventsService/LogEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientEventsServiceServer).LogEvents(ctx, req.(*LogEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientEventsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dolt.services.eventsapi.v1alpha1.ClientEventsService",
	HandlerType: (*ClientEventsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogEvents",
			Handler:    _ClientEventsService_LogEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dolt/services/eventsapi/v1alpha1/client_event.proto",
}
