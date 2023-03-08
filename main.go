package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// a shell is a basic user interface for
// an operating system,it takes in an arbitary
// amount of input and returns a corresponding
// output.

// bourne shell = sh
// bourne again shell = bash
// korn shell = ksh
// etc ...

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s", "> ")
		i, e := reader.ReadString('\n')
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s", e)
			return
		}
		if e = execInput(i); e != nil {
			fmt.Fprintf(os.Stderr, "%s", e)
		}
	}
}

func execInput(kinput string) error {
	rmline := strings.TrimSuffix(kinput, "\n")
	args := strings.Split(rmline, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
