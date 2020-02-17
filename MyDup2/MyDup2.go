/* MyDup2 prints the count and text of lines that appear more than once
in the input.  It reads either from stdin or from a list of named files
It prints the name of the file in which each duplicated line occurs*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "Stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, arg string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		myArg := arg
		myArg += " " + input.Text()
		counts[myArg]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
