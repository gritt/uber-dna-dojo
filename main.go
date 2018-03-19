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

	// the dna slice
	ds := getDNAFromFile()

	if !ds.isValid() {
		fmt.Println("error: DNA code has less than 6 nucleobases:")
		os.Exit(1)
	}

	// genes map with their permutations
	gm := map[string]genePermutations{
		"ACT": genePermutations{"CGT", "AGT"},
		"AGT": genePermutations{"CGT", "AGT"},
		"CGT": genePermutations{"ACT", "AGT"},
	}

	ds.findShortestOccurrence(gm)
}

func getDNAFromFile() dna {
	bs, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("coult not open file:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), "")

	return dna(s)
}

func (ds dna) isValid() bool {
	if len(ds) < 6 {
		return  false
	}
	return  true
}

func (ds dna) findShortestOccurrence(gpm map[string]genePermutations) {

	// g := ACT or AGT or CGT
	for g, gps := range gpm {

		//findGeneOccurrenceIndexes
		//when receives a result, start new routine for the subsequences search
		//with the index got from the channel, and with this first gene(g)

		fmt.Println("permutations of gene", g, "are", gps)

		// TODO: a channel to transport integers
		//ch := make(chan int)

		// TODO: run search in a new thread
		//go ds.findGeneOccurrenceIndexesConcurrent(g, ch)

		// TODO: listen for indexes, start a new search by the current gene from there
		//i := <-ch
		//fmt.Println("found gene", g, "at index", i)

		//idxs := ds.findGeneOccurrenceIndexes(g)
		//fmt.Println("indexes with permutation", g, ":", idxs)

	}
}

func (ds dna) findGeneOccurrenceIndexes(g string) []int {

	// dna length
	dsl := len(ds)

	var idxs []int

	for i := 0; i < dsl; i++ {

		if (i + 2) >= dsl {
			i++
			continue
		}

		cgs := string(ds[i] + ds[i+1] + ds[i+2])
		if cgs == g {

			// slice of all indexes with matches
			idxs = append(idxs, i)

			i += 2
		}
	}

	return idxs
}

func (ds dna) findGeneOccurrenceIndexesConcurrent(g string, ch chan int) {

	// dna length
	dsl := len(ds)

	for i := 0; i < dsl; i++ {

		if (i + 2) >= dsl {
			i++
			continue
		}

		cgs := string(ds[i] + ds[i+1] + ds[i+2])

		if cgs == g {

			// notify channel a match was found at this index
			ch <- i

			i += 2
		}
	}
}

func (ds dna) findRemainingGenesFromIndex(idx int, gpm map[string]genePermutations) string {
	return "hello from the other side"
}