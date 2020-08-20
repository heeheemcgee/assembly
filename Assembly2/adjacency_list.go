package main

//MakeOverlapNetwork() takes a slice of reads with match, mismatch, gap and a threshold.
//It returns adjacency lists of reads; edges are only included
//in the overlap graph is the alignment score of the two reads is at least the threshold.
func MakeOverlapNetwork(reads []string, match, mismatch, gap, threshold float64) map[string][]string {
	//Initialize an adjacency list to represent the overlap graph.
	adjacencyList := make(map[string][]string)
	mtx:= BinarizeMatrix(OverlapScoringMatrix(reads,match,mismatch,gap),threshold)
  for i := range mtx {
		for j := range mtx[i] {
			if mtx[i][j]==1 {
				adjacencyList[reads[i]] = append(adjacencyList[reads[i]], reads[j])
			}
		}
	}

	// fill in details of the function here.

	return adjacencyList
}
