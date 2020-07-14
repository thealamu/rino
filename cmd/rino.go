package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	flag "github.com/spf13/pflag"
)

//flags
var (
	// stay silent
	silent bool
	// filepath
	file string
)

func main() {
	parseFlags()

	//check stdin is available for reading
	fInfo, err := (os.Stdin.Stat())
	if err != nil {
		errorf("Can't stat standard input, %v", err)
	}
	//check we have a pipe
	if fInfo.Mode()&os.ModeNamedPipe == 0 {
		errorf("Can only read from pipe")
	}

	//start reading
	var reader io.Reader
	if silent {
		reader = os.Stdin
	} else {
		reader = io.TeeReader(os.Stdin, os.Stdout)
	}

	var s = bufio.NewScanner(reader)
	for s.Scan() {
		line := s.Text()
		fmt.Println(line)
	}
}

func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args)
	os.Exit(2)
}

func parseFlags() {
	flag.BoolVarP(&silent, "silent", "s", false, "only write to file")
	flag.StringVarP(&file, "file", "f", "", "file to write to")
	flag.Parse()
}
