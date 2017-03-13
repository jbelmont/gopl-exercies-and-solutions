// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Echo1 prints its command-line arguments.
package main

import (
	"os"
	"strings"
)

func main() {
	echo(os.Args)
}

func echo(args []string) string {
	return strings.Join(args[0:], " ")
}

// Exercise 1.1: Modify the echo program to also print os.Args[0], the name of the command that invoked it.
