#!/usr/bin/env bats

setup() {
    export PATH=$PATH:~/go/bin
    export NOMS_VERSION_NEXT=1
    load $BATS_TEST_DIRNAME/helper/common.bash
    cd $BATS_TMPDIR
    mkdir "dolt-repo-$$"
    cd "dolt-repo-$$"
    dolt init
    dolt table create -s=`batshelper 1pk5col-ints.schema` test
}

teardown() {
    rm -rf "$BATS_TMPDIR/dolt-repo-$$"
}

@test "start a sql shell and exit using exit" {
    skiponwindows "Works on Windows command prompt but not the WSL terminal used during bats"
    run bash -c "echo exit | dolt sql" 
    [ $status -eq 0 ]
    [[ "$output" =~ "# Welcome to the DoltSQL shell." ]] || false
    [[ "$output" =~ "Bye" ]] || false
}

@test "start a sql shell and exit using quit" {
    skiponwindows "Works on Windows command prompt but not the WSL terminal used during bats"
    run bash -c "echo quit | dolt sql"
    [ $status -eq 0 ]
    [[ "$output" =~ "# Welcome to the DoltSQL shell." ]] || false
    [[ "$output" =~ "Bye" ]] || false
}

@test "run a query in sql shell" {
    skiponwindows "Works on Windows command prompt but not the WSL terminal used during bats"
    run bash -c "echo 'select * from test;' | dolt sql"
    [ $status -eq 0 ]
    [[ "$output" =~ "# Welcome to the DoltSQL shell." ]] || false
    [[ "$output" =~ "pk" ]] || false
    [[ "$output" =~ "Bye" ]] || false
}