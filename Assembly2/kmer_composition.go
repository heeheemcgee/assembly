package main

// import ("fmt")

// func main(){
// 	fmt.Println(KmerComposition(GenerateRandomGenome(20), 3))
// }

//KmerComposition returns the k-mer composition (all k-mer substrings) of a given genome.
func KmerComposition(genome string, k int) []string {
	ret:= make([]string, 0)
	for i:=0; i<=len(genome)-k; i++ {
		ret = append(ret, genome[i:i+k])
	}
	return ret
}
