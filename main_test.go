package main

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
	"testing"
)

func assert(t *testing.T, b bool, msg string) {
	if !b {
		t.Error(msg)
	}
}

func getArgs() Flags {
	file := "input.txt"
	proc := 1
	pholder := false

	args := Flags{
		File:    &file,
		Proc:    &proc,
		Pholder: &pholder,
		VarArgs: []string{"echo"},
	}
	return args
}

func Test_ReadArgs(t *testing.T) {
	s, _ := ReadArgs("input.txt")
	assert(t, len(s) > 0, "File Open Error")
}

func Test_ExecTest(t *testing.T) {
	INSTR := "1234"
	output, err := ExecCommand("echo", INSTR)
	assert(t, err == nil, "ExecCommand")
	_ = output
	// assert(t, output == INSTR, "Test ExecTest Error")
}

func Test_Worker(t *testing.T) {
	const TESTSTR = "input.txt"
	args := getArgs()
	buff := &bytes.Buffer{}

	ch := make(chan bool, 4)
	defer close(ch)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go Worker(TESTSTR, args, &wg, ch, buff)
	wg.Wait()

	assert(t, buff.String() == TESTSTR, TESTSTR)
}

func Test_Code(t *testing.T) {
	buff := new(bytes.Buffer)

	fmt.Fprintln(buff, "test")
	assert(t, strings.Trim(buff.String(), "\n") == "test", buff.String())
}
