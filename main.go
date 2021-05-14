package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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
	help bool
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

	t, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Println(numbers)
	fmt.Print(string(t))
	return nil
}

func main() {
	if help {
		fmt.Print(`For a space-delimited string, outputs the string at the specified location.
usage: wd -n int
example: 
	$ echo "a b c" | wd -n 1
	a
`)
		return
	}
	err := run()
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
}