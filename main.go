package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

type Flags struct {
	File    *string
	Proc    *int
	Pholder *bool
	VarArgs []string
}

func usage() Flags {
	args := Flags{}
	args.File = flag.String("f", "", "Input file path, default")
	args.Pholder = flag.Bool("i", false, "Placeholder (default : {})")
	args.Proc = flag.Int("p", 1, "Thread pool count")
	flag.Parse()

	args.VarArgs = flag.Args()
	return args
}

func readArgsFromStream(fp *os.File) []string {
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

func ReadArgs(path string) ([]string, bool) {
	var ss []string

	if len(path) > 0 {
		fp, err := os.Open(path)
		if err != nil {
			return ss, false
		}
		defer fp.Close()
		ss = readArgsFromStream(fp)
	} else {
		ss = readArgsFromStream(os.Stdin)
	}

	return ss, true
}

func ExecCommand(command string, args ...string) (string, error) {
	reargs := []string{}

	if runtime.GOOS == "windows" {
		reargs = append(reargs, "/c")
		reargs = append(reargs, command)
		command = "cmd"
	}
	reargs = append(reargs, args...)

	output, err := exec.Command(command, reargs...).Output()
	return strings.Trim(string(output), "\n"), err
}

// Worker function
func Worker(input string, flagArgs Flags, wg *sync.WaitGroup, ch chan bool, fout io.Writer) {
	ch <- true
	defer func() { <-ch }()
	defer wg.Done()

	varargs := flagArgs.VarArgs
	varcmd := []string{}

	if *flagArgs.Pholder {
		for i, args := range varargs {
			varargs[i] = strings.ReplaceAll(args, "{}", input)
		}
	}
	varcmd = append(varcmd, varargs[1:]...)

	if !*flagArgs.Pholder {
		varcmd = append(varcmd, input)
	}

	output, _ := ExecCommand(varargs[0], varcmd...)
	fmt.Fprint(fout, output)
}

func Run() {
	flagArgs := usage()
	inputstr, ok := ReadArgs(*flagArgs.File)
	if !ok {
		fmt.Println("No Arguments ")
		return
	}

	ch := make(chan bool, *flagArgs.Proc)
	wg := sync.WaitGroup{}
	defer close(ch)
	defer wg.Wait()

	for _, input := range inputstr {
		wg.Add(1)
		go Worker(input, flagArgs, &wg, ch, os.Stdout)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	Run()
}
