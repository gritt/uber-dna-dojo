# UBER Coding Challenge (DNA)

A DNA is a string containing the following 4 characters: A, C, G or T (nucleobases). For a given DNA, the researchers would like to know what is the shortest piece of the DNA (consecutive nucleobases) that contain 3 predefined genes (namely ACT, AGT and CGT), in any order. If such piece does not exist, return an empty string;

Example 01:
    
    Sample input: ACTACGTTTAGTAACTCGTCT
    Sample output: AGTAACTCGT

Example 02:
    
    Sample input: ACTACGTACTTTAG
    Sample output: [empty string]


## How my solution works

    1. Will loop through the DNA array

    2. For each different permutation

    2.1. It will get the first gene, and will search in the array for occurrences of that gene, 
    it's going to return the index that the gene was found.
     
    2.2. For each result of the last search, a new search against the DNA array will happen, 
    starting already from the index that the first gene was found, 
    this search will look for the remaining two genes, in any order. 
    Each iteration increments the shortest piece path.
        
    2.3. When both genes are found, for all searches, the pathes will be added to an array 
    of results, so that we can compare how the different permudations performed on their searches, 
    with a different permutation we might have a shorter path.
        
    3. The array of results is sorter by length, and we print the first index, 
    which is the shortest one found;

    Extra. It does it concurrently, using Go routines and channels implementation,
    while it is not perceived for such small inputs (samples), it will do a big difference
    when dealing with bigger texts;


## Graphical representation
    ..................CGT...AGT............ACT
    
    TTTCGTTTTCGTTTTAGTCGTTTTAGTTTTTTTAGTTTTACTTTTACTTTTACTTTTCGTACT
    ...4,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,39.....................
    ........10..............................33.....................
    .................19.....................24.....................
    ........................................................58.....