package cmd

import "strconv"

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
