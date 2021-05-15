package cmd

import (
	"strings"
)

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
