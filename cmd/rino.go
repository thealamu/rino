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
		errorf("Could not stat standard input, %v\n", err)
	}
	//check we have a pipe
	if fInfo.Mode()&os.ModeNamedPipe == 0 {
		errorf("Can only read from pipe\n")
	}

	//prep for reading
	var reader io.Reader
	if silent {
		reader = os.Stdin
	} else {
		reader = io.TeeReader(os.Stdin, os.Stdout)
	}

	//prep file for writing
	if file == "" {
		errorf("Have no file to write to, file flag required\n")
	}
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		errorf("Could not open file for writing, %v\n", err)
	}
	defer f.Close()
	outFile := bufio.NewWriter(f)

	var s = bufio.NewScanner(reader)
	for s.Scan() {
		line := s.Text()
		fmt.Fprintln(outFile, line)
	}
}

func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(2)
}

func parseFlags() {
	flag.BoolVarP(&silent, "silent", "s", false, "only write to file")
	flag.StringVarP(&file, "file", "f", "", "file to write to")
	flag.Parse()
}
