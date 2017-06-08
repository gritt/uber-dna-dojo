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
    /** @var  DNA */
    private $dna;

    /** @var array */
    private $genePermnutations = [
        ['ACT', 'CGT', 'AGT'],
        ['ACT', 'AGT', 'CGT'],
        ['AGT', 'ACT', 'CGT'],
        ['AGT', 'CGT', 'ACT'],
        ['CGT', 'AGT', 'ACT'],
        ['CGT', 'ACT', 'AGT'],
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
        var_dump($this->dna->getCode());

        die;
    }
}

// How it works

// Will loop through the dna array

// For each permutation, each match of the first gene (ACT | CGT | AGT)
// will create a new search instance, which will loop from the instance index till the end

// The asymptotic notation for this algorithm is presented below, and
// the performance can be improved using more modern algorithms


// Result:
// ..................CGT...AGT............ACT

// TTTCGTTTTCGTTTTAGTCGTTTTAGTTTTTTTAGTTTTACTTTTACTTTTACTTTTCGTACT (dna string)
// ...4,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,39.....................
// ........10..............................33.....................
// .................19.....................24.....................
// ........................................................58.....


try {

    // Read from standard IO
    $handle = fopen("php://stdin", "r");
    $code = trim(fgets($handle));
    fclose($handle);


    /** @var DNA $dna */
    $dna = new DNA($code);

    /** @var DNAFinder $dnaFinder */
    $dnaFinder = new DNAFinder($dna);

    print $shortestPiece = $dnaFinder->findShortestPiece();

} catch (Exception $ex) {

    var_dump('ERROR');
    var_dump($ex->getLine());
    var_dump($ex->getMessage());
    die;
}