package main

//AverageOutDegree takes the adjacency list of an overlap network as input. It returns
//the average number of elements to which a given string is "connected"
//in the network.
func AverageOutDegree(adjList map[string][]string) float64 {
	avg := 0.0
	for _,val := range adjList{
		avg+=float64(len(val))
	}
	return avg/float64(len(adjList))
}
