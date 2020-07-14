package main

import (
	"fmt"
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
}

func errorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args)
	os.Exit(2)
}

func parseFlags() {
	flag.BoolVarP(&silent, "silent", "s", false, "only write to file")
	flag.StringVarP(&file, "file", "f", "", "file to write to")
	flag.Parse()
}
