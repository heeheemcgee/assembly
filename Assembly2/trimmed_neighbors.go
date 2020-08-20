package main

//GetTrimmedNeighbors takes in a string pattern (read), an adjacency list of an overlap network and an integer maxK.
//It returns the set of neighbors of pattern in the updated adjacency list after trimming all edges x -> y
//such that we can reach y from x via a path of length between 2 and maxK.
func GetTrimmedNeighbors(pattern string, adjList map[string][]string, maxK int) []string {
	neighbors := adjList[pattern]
	extendedNeighbors := GetExtendedNeighbors(pattern, adjList, maxK)
	numNeighbors := len(neighbors)
	for i := numNeighbors-1; i >= 0 ; i-- {
		if Contains(extendedNeighbors, neighbors[i]){
				neighbors = append(neighbors[:i], neighbors[i+1:]...)
		}
	}
	return neighbors
}
