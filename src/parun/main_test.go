package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"testing"
)

func assert(t *testing.T, b bool, msg string) {
	if !b {
		t.Error(msg)
	}
}

func Test_cmdString(t *testing.T) {
	s, _ := cmdString("main.go")
	assert(t, len(s) > 0, "File open eror")
}

func getArgs() Args {
	file := "../../input.txt"
	proc := 1
	pholder := false

	args := Args{
		File:    &file,
		Proc:    &proc,
		Pholder: &pholder,
		Cmd:     []string{"echo"},
	}
	return args
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

	result := strings.Trim(buff.String(), "\n")
	assert(t, result == TESTSTR, result)
}

func Test_Program(t *testing.T) {
	cmd := exec.Command("../../bin/parun.exe", "-f", "../../input.txt", "echo")

	out, err := cmd.Output()
	assert(t, err == nil, "exec.Command | cmd.Output()")
	assert(t, strings.Trim(string(out), "\n") == "input.txt", strings.Trim(string(out), "\n"))
}

func Test_Code(t *testing.T) {
	buff := new(bytes.Buffer)

	fmt.Fprintln(buff, "test")
	assert(t, strings.Trim(buff.String(), "\n") == "test", buff.String())
}
