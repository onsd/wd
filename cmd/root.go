package cmd

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
	"syscall"
)

func helpMessage() {
	fmt.Print(`For a space-delimited string, outputs the string at the specified location.
usage: wd -n int
example: 
	$ echo "a b c" | wd -n 1
	a
`)
}

var (
	numbers arrayFlags
	help    bool
)

func checkFlag() {
	flag.Var(&numbers, "n", "position of word")
	flag.BoolVar(&help, "h", false, "help")
	flag.Parse()
}

func Run() error {
	checkFlag()

	if help {
		helpMessage()
		return nil
	}

	var filename string
	if args := flag.Args(); len(args) > 0 {
		filename = args[0]
	}

	var r io.Reader
	switch filename {
	case "", "-":
		if terminal.IsTerminal(syscall.Stdin) {
			return nil
		}
		r = os.Stdin
	default:
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		r = f
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(Wd(scanner.Text(), numbers))
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
