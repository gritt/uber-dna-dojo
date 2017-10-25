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
type genePermutations []string
type results []string

func main() {
	//DNA slice
	ds := dnaFromFile()
	// Genes map with their permutations
	gm := map[string]genePermutations{
		"ACT": genePermutations{"CGT", "AGT"},
		"AGT": genePermutations{"CGT", "AGT"},
		"CGT": genePermutations{"ACT", "AGT"},
	}
	ds.findShortestOccurence(gm)
}

func (ds dna) findShortestOccurence(gm map[string]genePermutations) {
	fmt.Println(gm)
	fmt.Println(ds)
}

func dnaFromFile() dna {
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
