package main

import "fmt"

type OverlapMatrices map[string][][]int

//MakeOverlapNetworkBinary takes a collection of strings reads as well as a binary matrix.
//It returns the overlap network from this binary matrix.
func MakeOverlapGraphBinary(reads []string, binaryMatrix [][]int) map[string][]string {
	overlapGraph := make(map[string][]string)
	for i:= range binaryMatrix{
		currRead := reads[i]
		for j := range binaryMatrix[i]{
			if binaryMatrix[i][j] == 1{
				overlapGraph[currRead] = append(overlapGraph[currRead], reads[j])
			}
		}
	}
	return overlapGraph
}

//OverlapScoringMatrixMinimizers takes a collection of reads along with a value k,
//minimizer/maximizer dictionaries, and alignment penalties.
//It returns a matrix in which mtx[i][j] is the overlap alignment score of
//reads[i] with reads[j], if these two share a minimizer or maximizer k-mer.
func OverlapScoringMatrixMinimizers(reads []string, k int, numIndices int, minimizerDict StringIndex, match, mismatch, gap float64) [][]float64 {
	numReads := len(reads)
	mtx := InitializeSquareMatrix(numReads)
	for i := range mtx{
		if i % 100 == 0{
			fmt.Println("Filling in matrix for read", i)
		}
		for j:= range mtx[i]{
			mtx[i][j] = -1000000000.0
		}
		// grab all read indices sharing a minimizer hash with reads[i]
		currString := reads[i]
		n:= len(currString)/numIndices


		for p := 0; p < numIndices; p++ {
			var s string
			if p == 0 {
				s = currString[:(p+1)*n]
			} else if p == numIndices-1 {
				s = currString[p*n-k:]
			} else {
				s = currString[p*n-k : (p+1)*n]
			}
			// now find the minimizer of s
			m := Minimizer(s, k)

			indices := make([]int,0)
			for _,j := range minimizerDict[m]{
				if j != i && ArrayContains(j, indices)==false{
					indices= append(indices, j)
					// now we can update the matrix
					mtx[i][j] = ScoreOverlapAlignment(reads[i], reads[j], match, mismatch, gap)
				}
			}
		}
	}
	return mtx
}
