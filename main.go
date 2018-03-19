package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sort"
)

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
		"AGT": genePermutations{"CGT", "ACT"},
		"CGT": genePermutations{"ACT", "AGT"},
	}

	// find shortest piece of matching genes
	rs := ds.findShortestOccurrence(gm)

	// result string
	fmt.Println(rs)
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

func (ds dna) findShortestOccurrence(gpm map[string]genePermutations) string {

	// to store results
	var rsl []string

	// gs : ACT|AGT|CGT
	// gpsl : {"CGT", "AGT"},{"CGT", "AGT"},{"CGT", "AGT"}
	for gs, gpsl := range gpm {

		// TODO: a channel to transport integers
		//ch := make(chan int)

		// TODO: run search in a new routine
		//go ds.findGeneOccurrenceIndexesConcurrent(gs, ch)

		// TODO: listen for indexes, start a new search by the current gene from there
		//i := <-ch

		isl := ds.findGeneOccurrenceIndexes(gs)

		if len(isl) == 0 {
			continue
		}

		for _, i := range isl {

			ps := ds.findRemainingGenesFromIndex(i, gpsl)

			if ps != "false" {
				rsl = append(rsl, ps)
			}
		}
	}

	if len(rsl) > 0 {
		// sort slice of strings by length of each string
		sort.Slice(rsl, func(j, k int) bool { return len(rsl[j]) < len(rsl[k]) })

		// return the first one, the shortest!
		return rsl[0]
	}

	return ""
}

func (ds dna) findGeneOccurrenceIndexes(g string) []int {

	// dna length
	l := len(ds)

	// indexes slice
	var isl []int

	for i := 0; i <= l; i++ {

		if (i + 2) >= l {
			i++
			continue
		}

		// compare found nucleobase string with given nucleobase
		cgs := string(ds[i] + ds[i+1] + ds[i+2])
		if cgs == g {

			// append match index to slice of results
			isl = append(isl, i)

			i += 2
		}
	}

	return isl
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

func (ds dna) findRemainingGenesFromIndex(i int, gps []string) string {

	// dna length
	l := len(ds)

	// we already know the first nucleobase was found at [i][i+1][i+2] of the slice
	fnbs := ds[i] + ds[i+1] + ds[i+2]

	// so our path starts here
	i += 3

	// we have to find the following nucleobases:
	snbs := gps[0]
	fsn := false
	tnbs := gps[1]
	ftn := false

	// till there we'll keep appending to the gs
	for j:= i; j < l; j++ {

		// add the [0] char
		fnbs = fnbs + ds[j]

		// needs at least more 2 chars to form a nucleobase
		if j+2 >= l {
			j++
			continue
		}

		// comparable nocleobase
		cnbs := ds[j] + ds[j+1] + ds[j+2]

		// has found the second nucleobase
		if cnbs == snbs {

			// add the second nucleobase to the first nucleobase
			fnbs = fnbs + ds[j+1] + ds[j+2]
			fsn = true

			// has found the second, third have been found already
			if ftn {
				return fnbs
			}

			j += 2
			continue
		}

		// has found the third nucleobase
		if cnbs == tnbs {

			// add the third nucleobase to the first nucleobase
			fnbs = fnbs + ds[j+1] + ds[j+2]
			ftn = true

			// has found the third, second have been found already
			if fsn {
				return fnbs
			}

			j += 2
			continue
		}
	}

	return "false"
}