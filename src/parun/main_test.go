

package main

import (
	"bytes"
	"testing"
	_ "sync"
	"fmt"
	"strings"
	"os/exec"
)


func assert(t *testing.T, b bool, msg string) {
	if !b { t.Error(msg) }
}

func Test_cmdString(t *testing.T) {
	s, _ := cmdString("main.go")
	assert(t, len(s) > 0, "File open eror")
} 

func Test_Worker(t *testing.T) {
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