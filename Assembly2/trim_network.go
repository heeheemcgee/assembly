package main

//TrimNetwork takes in an overlap network adjList and a max iteration maxK
//It removes all n-transitivity edges in adjList, n <= maxK.
//It returns a trimmed copy of the original network.
func TrimNetwork(adjList map[string][]string, maxK int) map[string][]string {
	adjListTrimmed := make(map[string][]string)
	for key := range adjList{
		adjListTrimmed[key] = GetTrimmedNeighbors(key,adjList, maxK)
	}
	return adjListTrimmed
}
