package main

import (
	"github.com/onsd/wd/cmd"
	"log"
	"os"
)

func main() {
	err := cmd.Run()
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
}
