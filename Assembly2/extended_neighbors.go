package main

//GetExtendedNeighbors takes in a pattern (read), the overlap graph and maxK.
//It returns the extendedNeighbors list. For each neighbor *n* in this list,
//distance between n and pattern is between 2 to maxK.
func GetExtendedNeighbors(pattern string, adjList map[string][]string, maxK int) []string {
	extendedNeighbors := []string{}
	currentNeighbors := adjList[pattern]

	for j := 2; j <= maxK; j++ {
		currentNeighbors = AdjacentStrings(currentNeighbors, adjList)
		currentNeighbors = RemovePattern(currentNeighbors, pattern)
		extendedNeighbors = append(extendedNeighbors, currentNeighbors...)
	}

	return extendedNeighbors

}

func AdjacentStrings(currentNeighbors []string, adjList map[string]([]string)) []string {
	reachableNeighbors := make([]string, 0)

	for i := 0; i < len(currentNeighbors); i++ {
		reachableNeighbors = append(reachableNeighbors, adjList[currentNeighbors[i]]...)
	}
	return reachableNeighbors
}

func RemovePattern(patterns []string, text string) []string {
	for i := len(patterns) - 1; i >= 0; i-- {
		if patterns[i] == text {
			patterns = append(patterns[:i], patterns[i+1:]...)
		}
	}
	return patterns
}

func Contains(patterns []string, pattern string) bool {
	for _, val := range patterns {
		if val == pattern {
			return true
		}
	}
	return false
}
