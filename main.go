package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
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

	result := wd(string(t), numbers)
	fmt.Println(result)
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

func helpMessage(){
	fmt.Print(`For a space-delimited string, outputs the string at the specified location.
usage: wd -n int
example: 
	$ echo "a b c" | wd -n 1
	a
`)
}

func wd(str string, number []int) string {
	inputArray := strings.Split(strings.Trim(str, "\n"), "\n")
	var out []string
	for _, input := range inputArray {
		array := strings.Fields(input)
		var outLine []string
		for _, n := range number {
			if n-1 < 0 && len(array) < n-1 {
				continue
			}
			outLine = append(outLine, array[n-1])
		}
		out = append(out, strings.Join(outLine, " "))
	}

	return strings.Join(out, "\n")
}