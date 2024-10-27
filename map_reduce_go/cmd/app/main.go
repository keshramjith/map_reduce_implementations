package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		panic("input files required.")
	}
	dirEntries, err := os.ReadDir(args[0])
	if err != nil {
		panic(err)
	}

	for _, dirEntry := range dirEntries {
		fullPath := fmt.Sprintf("../input_files/%s", dirEntry.Name())
		data, err := os.ReadFile(fullPath)
		if err != nil {
			panic(err)
		}
		fmt.Print(string(data))
	}
}
