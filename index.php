<?php

/**
 * Class DNA
 * @author Gilvan Ritter <gilvanritter@gmail.com>
 */
class DNA
{
    /** @var string */
    private $code;


    /**
     * DNA constructor.
     * @param string $code
     */
    public function __construct(string $code)
    {
        if ($this->isValidDNA($code)) {
            $this->code = $code;
        }
    }

    /**
     * @return string
     */
    public function getCode()
    {
        return $this->code;
    }

    /**
     * @return int
     */
    public function getSize()
    {
        return strlen($this->code);
    }

    /**
     * @return array
     */
    public function toArray()
    {
        return str_split($this->code, 1);
    }

    /**
     * @param $code
     * @return bool
     * @throws Exception
     */
    private function isValidDNA($code)
    {
        if (strlen($code) < 6) {
            throw new Exception('Error: DNA code has less than 6 nucleobases');
        }
        return true;
    }
}


/**
 * Class DNAFinder
 * @author Gilvan Ritter <gilvanritter@gmail.com>
 */
class DNAFinder
{
    /** @var DNA */
    private $dna;

    /** @var array */
    private $dnaArray;

    /** @var array */
    private $genePermutations = [
        ['ACT', 'CGT', 'AGT'], # or ['ACT', 'AGT', 'CGT']
        ['AGT', 'ACT', 'CGT'], # or ['AGT', 'CGT', 'ACT']
        ['CGT', 'AGT', 'ACT'], # or ['CGT', 'ACT', 'AGT']
    ];


    /**
     * DNAFinder constructor.
     * @param DNA $dna
     */
    public function __construct(DNA $dna)
    {
        $this->dna = $dna;
    }

    public function findShortestPiece()
    {
        /** @var array dnaArray */
        $this->dnaArray = $this->dna->toArray();

        /** @var array $searchResults */
        $searchResults = [];

        foreach ($this->genePermutations as $genePermutation) {

            /** @var string $gene : ACT | CGT | AGT */
            $gene = $genePermutation[0];

            /** @var array $indexesToStartSearches */
            $indexesToStartSearches = $this->findGeneIndexesOccurrences($gene);

            if (empty($indexesToStartSearches)) {
                continue;
            }

            foreach ($indexesToStartSearches as $searchIndex) {

                $remainingGenesToFind = [
                    $genePermutation[1],
                    $genePermutation[2],
                ];

                $indexForSearchResults = implode('', $genePermutation);

                /** @var string $path */
                $path = $this->searchFromIndexTillFindPath($searchIndex, $remainingGenesToFind);

                if ($path) {
                    $searchResults[$indexForSearchResults]['path_from_index'][$searchIndex] = $path;
                }
            }
        }

        var_dump($searchResults);
        die;
    }

    /**
     * @param $gene
     * @return array
     */
    private function findGeneIndexesOccurrences($gene)
    {
        /** @var $dnaArray */
        $dnaArray = $this->dnaArray;

        /** @var  $indexesOccurrences */
        $indexesOccurrences = [];

        for ($i = 0; $i <= $this->dna->getSize(); $i++) {

            if (!isset($dnaArray[$i + 2])) {
                $i++;
                continue;
            }

            $comparableGene = $dnaArray[$i] . $dnaArray[$i + 1] . $dnaArray[$i + 2];

            if ($comparableGene == $gene) {
                $indexesOccurrences[] = $i;
                $i += 2;
            }
        }

        return $indexesOccurrences;
    }

    /**
     * @param $index
     * @param $genesArray
     * @return bool|string
     */
    private function searchFromIndexTillFindPath($index, $genesArray)
    {
        $dnaArray = $this->dnaArray;

        // We already know the first gene was found on [$i][$i+1][$i+2] positions of the array
        $path = $dnaArray[$index] . $dnaArray[$index + 1] . $dnaArray[$index + 2];

        // So our path will start here
        $index += 3;

        // We have to find these genes, till there we'll keep incrementing the $path
        $secondGene = $genesArray[0];
        $foundSecond = false;
        $thirdGene = $genesArray[1];
        $foundThird = false;

        for ($i = $index; $i <= $this->dna->getSize(); $i++) {

            if (!isset($dnaArray[$i])) {
                continue;
            }

            $path .= $dnaArray[$i];

            if (!isset($dnaArray[$i + 2])) {
                $i++;
                continue;
            }

            $comparableGene = $dnaArray[$i] . $dnaArray[$i + 1] . $dnaArray[$i + 2];

            if ($comparableGene == $secondGene) {
                $foundSecond = true;

                $path .= $dnaArray[$i + 1] . $dnaArray[$i + 2];
                if ($foundThird) {
                    return $path;
                }

                $i += 2;
                continue;
            }

            if ($comparableGene == $thirdGene) {
                $foundThird = true;

                $path .= $dnaArray[$i + 1] . $dnaArray[$i + 2];

                if ($foundSecond) {
                    return $path;
                }

                $i += 2;
                continue;
            }
        }

        // Means it has not found both genes in the sequence of indexes
        return false;
    }
}

// How it works

// Will loop through the dna array

// Each permutation, and each match of the first gene (ACT | CGT | AGT)
// will result in a new search instance,
// which will loop through the dnsArray from the index that the first match happened till the end
// searching for the remaining two genes

// The asymptotic notation for this algorithm is presented below,
// the performance could probably be improved using more modern algorithms


// Result:
// ..................CGT...AGT............ACT

// TTTCGTTTTCGTTTTAGTCGTTTTAGTTTTTTTAGTTTTACTTTTACTTTTACTTTTCGTACT (dna string)
// ...4,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,39.....................
// ........10..............................33.....................
// .................19.....................24.....................
// ........................................................58.....


try {

    // From standard input
    $handle = fopen("php://stdin", "r");
    $code = trim(fgets($handle));
    fclose($handle);


    /** @var DNA $dna */
    $dna = new DNA($code);

    /** @var DNAFinder $dnaFinder */
    $dnaFinder = new DNAFinder($dna);

    /** @var string $shortestPiece */
    $shortestPiece = $dnaFinder->findShortestPiece();

    // To standard output
    print $shortestPiece . PHP_EOL;

} catch (Exception $ex) {

    var_dump('ERROR');
    var_dump($ex->getLine());
    var_dump($ex->getMessage());
    die;
}