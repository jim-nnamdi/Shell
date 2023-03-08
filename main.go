package main

import (
	"bufio"
	"errors"
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

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "%s", errors.New("path required"))
			return errors.New("path required")
		}
		return os.Chdir(args[1])
	case "touch":
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "%s", errors.New("no file specified"))
			isfile := strings.Contains(args[1], ".")
			if !isfile {
				fmt.Fprintf(os.Stderr, "%s", errors.New("specify a valid extension"))
				return errors.New("specify a valid file extension")
			}
			newfile, e := os.Create(args[1])
			if e != nil {
				fmt.Fprintf(os.Stderr, "%s", errors.New("cannot create file"))
				return e
			}
			return newfile.Chmod(0777)
		}
	case "exit":
		os.Exit(0)
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
