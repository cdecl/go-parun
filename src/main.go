package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

type Args struct {
	File    *string
	Proc    *int
	Pholder *bool
	Cmd     []string
}

func (p Args) Debug() {
	fmt.Println(fmt.Sprintf("File[%s], Proc[%d], Pholder[%s] : Cmd[%s]", *p.File, *p.Proc, *p.Pholder, p.Cmd))
}

func usage() Args {
	args := Args{}
	args.File = flag.String("f", "", "Input file path, default")
	args.Pholder = flag.Bool("i", false, "Placeholder (default : {})")
	args.Proc = flag.Int("p", 1, "Thread pool count")
	flag.Parse()

	args.Cmd = flag.Args()
	// args.Debug()
	return args
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

func Worker(redirect string, args Args, wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	defer func() { <-ch }()
	defer wg.Done()

	cmdstr := args.Cmd
	varcmd := []string{}

	if *args.Pholder {
		for i, cmd := range cmdstr {
			cmdstr[i] = strings.ReplaceAll(cmd, "{}", redirect)
		}
	}

	varcmd = append(varcmd, cmdstr[1:]...)

	if !*args.Pholder {
		varcmd = append(varcmd, redirect)
	}

	cmd := exec.Command(cmdstr[0], varcmd...)
	out, _ := cmd.Output()
	fmt.Print(string(out))
}

func Run() {
	args := usage()
	cmds, ok := cmdString(*args.File)
	if !ok {
		fmt.Println("No Command")
		return
	}

	ch := make(chan bool, *args.Proc)
	wg := sync.WaitGroup{}
	defer close(ch)
	defer wg.Wait()

	for _, cmd := range cmds {
		wg.Add(1)
		go Worker(cmd, args, &wg, ch)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	Run()
}
