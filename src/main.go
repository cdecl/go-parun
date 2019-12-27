package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

type Param struct {
	File    *string
	Proc    *int
	Pholder *string
	Cmd     []string
}

func (p Param) Debug() {
	fmt.Println(fmt.Sprintf("File[%s], Proc[%d], Pholder[%s] : Cmd[%s]", *p.File, *p.Proc, *p.Pholder, p.Cmd))
}

func usage() Param {
	param := Param{}
	param.File = flag.String("f", "", "Input file path, default")
	param.Pholder = flag.String("i", "{}", "Placeholder string (default : {})")
	param.Proc = flag.Int("p", 1, "Thread pool count")
	flag.Parse()

	param.Cmd = flag.Args()
	// param.Debug()
	return param
}

func readFile(fp *os.File) []string {
	var ss []string

	io := bufio.NewReader(fp)
	for {
		line, isPrefix, err := io.ReadLine()
		if isPrefix || err != nil {
			break
		}
		ss = append(ss, string(line))
	}
	return ss
}

func cmdString(path string) ([]string, bool) {
	var ss []string

	if len(path) > 0 {
		fp, err := os.Open(path)
		if err != nil {
			return ss, false
		}
		defer fp.Close()
		ss = readFile(fp)
	} else {
		ss = readFile(os.Stdin)
	}

	return ss, true
}

func Worker(s string, param []string, wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	defer func() { <-ch }()
	defer wg.Done()

	varcmd := []string{"/c", s}
	varcmd = append(varcmd, param...)

	// cmd := exec.Command("cmd", "/c", s)
	cmd := exec.Command("cmd", varcmd...)
	out, _ := cmd.Output()
	fmt.Print(string(out))
}

func Run() {
	param := usage()
	cmds, ok := cmdString(*param.File)
	if !ok {
		fmt.Println("No Command")
		return
	}

	ch := make(chan bool, *param.Proc)
	wg := sync.WaitGroup{}
	defer close(ch)
	defer wg.Wait()

	for _, cmd := range cmds {
		wg.Add(1)
		go Worker(cmd, param.Cmd, &wg, ch)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	Run()
}
