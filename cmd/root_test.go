package cmd

import (
	"flag"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	testcases := []testcase{
		{
			Name:   "",
			Input:  "a b c",
			Number: []int{1},
			Expect: "a",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			flag.CommandLine.Set("h", "false")
			for _, i := range strings.Split(testcase.Input, " ") {
				flag.CommandLine.Set("n", i)
			}
			Run()
		})
	}
}
