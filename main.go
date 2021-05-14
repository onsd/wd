package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type arrayFlags []int

func (i *arrayFlags) String() string {
	return ""
}
func (i *arrayFlags) Set(value string) error {
	n, err := strconv.Atoi(value)
	if err != nil {
		return err
	}

	*i = append(*i, n)
	return nil
}

var (
	numbers arrayFlags
	help    bool
)

func init() {
	flag.Var(&numbers, "n", "position of word")
	flag.BoolVar(&help, "h", false, "help")
	flag.Parse()
}

func run() error {
	var filename string
	if args := flag.Args(); len(args) > 0 {
		filename = args[0]
	}

	var r io.Reader
	switch filename {
	case "", "-":
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
		fmt.Println(wd(scanner.Text(), numbers))
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func main() {
	if help {
		helpMessage()
		return
	}
	err := run()
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
}

func helpMessage() {
	fmt.Print(`For a space-delimited string, outputs the string at the specified location.
usage: wd -n int
example: 
	$ echo "a b c" | wd -n 1
	a
`)
}

func wd(line string, number []int) string {
	var out []string
	array := strings.Fields(line)
	for _, n := range number {
		if n < 0 || len(array) < n {
			continue
		}
		out = append(out, array[n-1])
	}
	return strings.Join(out, " ")
}
