package main

import "math/rand"

//ShuffleStrings takes a collection of strings patterns as input.
//It returns a random shuffle of the strings.
func ShuffleStrings(patterns []string) []string {
	ret := make([]string, len(patterns))
	order := rand.Perm(len(patterns))
	for i:=0; i<len(patterns); i++ {
		ret[i] = patterns[order[i]]
	}
	return ret
}
