// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	filesWithDup := []string{}
	if len(files) == 0 {
		filesWithDup = append(filesWithDup, countLines(os.Stdin, counts, files[0]))
	} else {
		for idx, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			filesWithDup = append(filesWithDup, countLines(f, counts, files[idx]))
			f.Close()
		}
	}

	i := 0
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, filesWithDup[i])
			i++
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileName string) string {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	return fileName
}
