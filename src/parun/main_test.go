

package main

import (
	"bytes"
	"testing"
	"sync"
	"fmt"
)


func assert(t *testing.T, b bool, msg string) {
	if !b {
		t.Error(msg)
	}
}

func getArgs() Args {
	args := Args{
			File:    func(s string) *string { return &s }(""),
			Proc:    func(s int) *int { return &s }(1),
			Pholder:func(s bool) *bool { return &s }(false),
			Cmd:   []string{"echo"},
	}
		
	return args
}
func Test_cmdString(t *testing.T) {
	s, _ := cmdString("main.go")
	assert(t, len(s) > 0, "File open eror")
} 

func Test_Worker(t *testing.T) {
	ch := make(chan bool, 1)
	wg := sync.WaitGroup{}
	defer close(ch)
	defer wg.Wait()

	args := getArgs()

	buff := new(bytes.Buffer)
	go Worker("test", args, &wg, ch, buff)
		
	fmt.Println(buff)
	fmt.Println("testssss")
}
