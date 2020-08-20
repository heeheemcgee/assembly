package main

import (
	"math/rand"
	"time"
)

//GenerateRandomGenome takes a parameter genomeLength and returns
//a random DNA string of this length where each nucleotide has equal probability.
func GenerateRandomGenome(genomeLength int) string {
	rand.Seed(time.Now().UnixNano())
	var genome string
	var random int
	for i := 0; i < genomeLength; i++ {
		random = rand.Intn(4)
		if random == 0 {
			genome += "A"
		} else if random == 1 {
			genome += "T"
		} else if random == 2 {
			genome += "C"
		} else {
			genome += "G"
		}
	}
	return genome
}

func RandomDNASymbol() string {
	random := rand.Intn(4)
	if random == 0 {
		return  "A"
	} else if random == 1 {
		return  "T"
	} else if random == 2 {
		return  "C"
	} else {
		return  "G"
	}
}
