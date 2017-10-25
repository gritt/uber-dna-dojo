package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
 * [v] load file into string slice
 * [ ] loop through DNA slice
 * [ ] for each permutation (gene) found point
 * [ ] find subsequences from the gene point till end/complete the DNA
 * [ ] compare matching subsequences sizes
 * [ ] choose the smaller :)
 */

type dna []string

func main() {
	ds := newDNAFromFile()
	fmt.Println(ds)
}

func newDNAFromFile() dna {
	bs, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), "")
	return dna(s)
}
