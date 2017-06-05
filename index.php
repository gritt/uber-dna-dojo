<?php

/**
 * Date: 04/06/17
 * Time: 15:19
 */

/*
 *
 * each match of the first gene, will create a new search with the rest of array
 *
CGT[n]AGT[n]ACT
TTTCGTTTTCGTTTTAGTCGTTTTAGTTTTTTTAGTTTTACTTTTACTTTTACTTTTCGTACT
...4,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,39.....................
........10..............................33.....................
.................19.....................24.....................
........................................................58.....
 *
 */

function get_gene_index_in_array($gene, $dnaArray)
{
    $dnaArraySize = count($dnaArray);

    $searches = [];

    for ($i = 0; $i <= $dnaArraySize; $i++) {

        if (!isset($dnaArray[$i + 2])) {
            $i++;
            continue;
        }

        $dnaString = $dnaArray[$i] . $dnaArray[$i + 1] . $dnaArray[$i + 2];

        if ($dnaString == $gene) {

            $searches['index'] = $i;

            $i = $i + 3;
            continue;
        }
        $i++;
        continue;
    }

    return $searches;
}

function find_genes($dna)
{
    $genePermutations = [
        ['ACT', 'CGT', 'AGT'],
        ['ACT', 'AGT', 'CGT'],
        ['AGT', 'ACT', 'CGT'],
        ['AGT', 'CGT', 'ACT'],
        ['CGT', 'AGT', 'ACT'],
        ['CGT', 'ACT', 'AGT'],
    ];

    $dnaArray = str_split($dna, 1);
    $dnaArraySize = count($dnaArray);
    $searches = [];


    // $permutation = ['CGT','AGT','ACT'],
    foreach ($genePermutations as $genePermutation) {

        // $gene = 'CGT'
        $gene = $genePermutation[0];
        $sequence = $genePermutation[1];
        $match = $genePermutation[2];

        $geneSearches = get_gene_index_in_array($gene, $dnaArray);

        if (empty($geneSearches)) {
            print "" . PHP_EOL;
            return;
        }

        // TTTCGTTTTCGTTTTAGTCGTTTTAGTTTTTTTAGTTTTACTTTTACTTTTACTTTTCGTACT
        // $searches[0][index] = 4  + 3
        // $searches[1][index] = 10 + 3
        // $searches[2][index] = 19 + 3
        // $searches[3][index] = 58 + 3

        $searchIndex = 0;
        foreach ($geneSearches as $search) {

            // search started in index
            // search ended on index
            // diff between end and start = implode piece from array
            $searchStartAt = $search['index'] + 3;



        }
    }

//    $genePermutationsSize = count($genePermutations);
    // $genePermutation = ['CGT','AGT','ACT'],
//    for ($genePermutationIndex = 0; $genePermutationIndex <= $genePermutationsSize; $genePermutationIndex++) {
//
//        $genePermutation = $genePermutations[$genePermutationsSize];
////    foreach ($genePermutations as $genePermutation) {
//

//        foreach ($genePermutation as $gene) {
//
//            for ($i = 0; $i <= $dnaArraySize; $i++) {
//
//                if (!isset($dnaArray[$i + 2])) {
//                    $i++;
//                    continue;
//                }
//
//                $dnaString = $dnaArray[$i] . $dnaArray[$i + 1] . $dnaArray[$i + 2];
//
//                if ($dnaString == $gene) {
//                    $searches[$gene][]['index'] = $i;
//
//                    $i = $i + 3;
//                    continue;
//
//                    // found CGT three times will lead to three searches for cgt.next starting at

//                }
//
//                $i++;
//                continue;
//            }
//
//            //$nextGene =;
//            foreach ($searches[$gene] as $search => $searchKeyIndex) {
//
//
//            }
////        }
//        }
//    }
}


$handle = fopen("php://stdin", "r");

$dna = fgets($handle);

find_genes($dna);

fclose($handle);