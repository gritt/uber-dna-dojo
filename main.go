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


	// channels to transport int and string of the first gene occurrences
	ci := make(chan []int)
	cs := make(chan []string)


	// gs : ACT|AGT|CGT
	// gpsl : {"CGT", "AGT"},{"CGT", "AGT"},{"CGT", "AGT"}
	for gs, gpsl := range gpm {

		// search matching genes indexes in new routine
		go ds.findGeneOccurrenceIndexes(gs, gpsl, ci, cs)
	}

	// wait for channel response till len(gpm) iterations
	for a := 0; a < len(gpm); a++ {

		isl := <- ci
		gpsl := <- cs

		if len(isl) == 0 {
			continue
		}

		// a channel to transport strings
		c := make(chan string)

		for _, i := range isl {
			// search remaining genes in new routine
			go ds.findRemainingGenesFromIndex(i, gpsl, c)
		}


		// wait for channel response till len(isl) iterations
		for j := 0; j < len(isl); j++ {

			// capture response
			// append correct ones to further be compared
			ps := <- c

			if ps != "false" {
				rsl = append(rsl, ps)
			}
		}
	}

	if len(rsl) > 0 {

		// sort slice of strings by length of each string
		sort.Slice(rsl, func(j, k int) bool {
			return len(rsl[j]) < len(rsl[k])
		})

		// return the first one, the shortest!
		return rsl[0]
	}

	return ""
}

func (ds dna) findGeneOccurrenceIndexes(g string, gpsl []string, ci chan []int, cs chan []string) {

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

	ci <- isl
	cs <- gpsl
	return
}

func (ds dna) findRemainingGenesFromIndex(i int, gps []string, c chan string) {

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
				c <- fnbs
				return
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
				c <- fnbs
				return
			}

			j += 2
			continue
		}
	}

	c <- "false"
}