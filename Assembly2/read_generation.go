package main

import (
	"math/rand"
)

//SimulateReadsClean takes a genome along with a read length and a probability.
//It returns a collection of strings resulting from simulating clean reads,
//where a given position is sampled with the given probability.
func SimulateReadsClean(genome string, readLength int, probability float64) []string {
	reads:= make([]string,0)
	random:=0.0

	for i:= 0; i<len(genome)-readLength+1; i++ {
		random = rand.Float64()
		if random < probability {
			reads = append(reads,genome[i:i+readLength])
		}
	}
	return reads
}
//NOTE: we will simulate messy reads together.

//SimulateReadsMessy takes a genome along with a read length, a probability,
//and error rates for substitutions, insertions, and deletions.
//It returns a collection of reads resulting from simulating clean reads,
//where a given position is sampled with the given probability, and
//errors occur at the respective rates.
func SimulateReadsMessy(genome string, readLength int, probability float64, substitutionErrorRate, insertionErrorRate, deletionErrorRate float64) []string {
	n := len(genome)
	reads := make([]string, 0)

	for i := 0; i < n-readLength+1; i++ {
		x := rand.Float64()
		if x < probability {
			//string selected! start reading out symbols
			currString := ""
			currPosition := i
			for len(currString) < readLength && currPosition < len(genome) {
				x := rand.Float64()
				if x < substitutionErrorRate {
					//make a substitution
					sym := RandomDNASymbol()
					for sym == string(genome[currPosition]) {
						sym = RandomDNASymbol()
					}
					//symbol found! append to current string
					currString += string(sym)
					currPosition++
				} else if x < substitutionErrorRate+insertionErrorRate { //insertion!
					currString += string(RandomDNASymbol())
					//don't advance position
				} else if x < substitutionErrorRate+insertionErrorRate+deletionErrorRate {
					//take no action other than moving position
					currPosition++
				} else {
					//take current symbol
					currString += string(genome[currPosition])
					currPosition++
				}
			}
			//make sure that we don't fall off the edge of the string

			if len(currString) == readLength { // make sure that string has appropriate length (e.g., it didn't hit end of genome too quickly)
				reads = append(reads, currString)
			}
		}
	}
	return reads
}
