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

	// DNA slice
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

	for fg, gps := range gm {

		//todo
		//channels
		//findGeneOccurenceIndexes
		//when receives a result, start new routine for the subsequences search
		//with the index got from the channel, and with this first gene(fg)

		fmt.Println("permutations of gene", fg, "are", gps)

		ch := make(chan int)

		go ds.findGeneOccurenceIndexes(fg, ch)

		i := <-ch

		fmt.Println("found gene", fg, "at index", i)
	}

	// fmt.Println(gm)
	// fmt.Println(ds)
}

func (ds dna) findGeneOccurenceIndexes(g string, ch chan int) {

	// DNA length
	dsl := len(ds)

	for i := 0; i < dsl; i++ {

		if (i + 2) >= dsl {
			i++
			continue
		}

		cgs := string(ds[i] + ds[i+1] + ds[i+2])
		if cgs == g {

			ch <- i

			// todo
			// found?
			// communicate through channel the index which was found
			// start subsequences search from there witht the permutations

			i += 2
		}
		// fmt.Println(i)
		// fmt.Println(gs)
	}
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
