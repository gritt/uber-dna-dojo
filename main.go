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
type permutation []string

func main() {
	ds := newDNAFromFile()
	fmt.Println(ds)

	// ['ACT', 'CGT', 'AGT']
	// ['AGT', 'CGT', 'ACT']
	// ['CGT', 'ACT', 'AGT']
	gp := map[string]permutation{
		"ACT": permutation{"CGT", "AGT"},
		"AGT": permutation{"CGT", "AGT"},
		"CGT": permutation{"ACT", "AGT"},
	}
	fmt.Println(gp)

}

func newDNAFromFile() dna {
	bs, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), "")

	if len(s) < 6 {
		os.Exit(1)
	}

	return dna(s)
}
