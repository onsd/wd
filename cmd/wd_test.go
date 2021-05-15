package cmd

import "testing"

type testcase struct {
	Name   string
	Input  string
	Number []int
	Expect string
}

func TestWd(t *testing.T) {
	testcases := []testcase{
		{
			Name:   "can select one word",
			Input:  "a b c",
			Number: []int{1},
			Expect: "a",
		},
		{
			Name:   "can select two words",
			Input:  "a b c",
			Number: []int{1, 3},
			Expect: "a c",
		},
		{
			Name:   "if negative number is selected, wd returns empty string",
			Input:  "a b c",
			Number: []int{-1},
			Expect: "",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			got := Wd(testcase.Input, testcase.Number)
			if got != testcase.Expect {
				t.Errorf("wd() = %v, want %v", got, testcase.Expect)
			}
		})
	}

}
