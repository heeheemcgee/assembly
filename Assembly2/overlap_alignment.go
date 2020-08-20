package main

//ALL PENALTIES POSITIVE

//This implementation of overlap alignment intializes the first row of the
//scoring matrix to be -infinite float64(math.MinInt64)

//ScoreOverlapAlignment takes two strings along with match, mismatch, and gap penalties.
//It returns the maximum score of an overlap alignment in which str0 is overlapped with str1.
//Assume we are overlapping a suffix of str0 with a prefix of str1.
func ScoreOverlapAlignment(str0, str1 string, match, mismatch, gap float64) float64 {
    mtx := OverlapScoreTable(str0, str1, match, mismatch, gap)
    return mtx[len(mtx)-1][len(mtx[0])-1]
}

//OverlapScoreTable takes two strings and alignment penalties. It returns a 2-D array
//holding dynamic programming scores for overlap alignment of str0 with str1 with these penalties.
//Assume we are overlapping a suffix of str0 with a prefix of str1.
func OverlapScoreTable(str0, str1 string, match, mismatch, gap float64) [][]float64 {
    graph := InitializeScoreTable(len(str0)+1, len(str1)+1)
		for j := range graph[0] {
		graph[0][j] = float64(-1) * gap * float64(j)
		}
		for i := 1; i < len(graph); i++ {
			for j := 1; j < len(graph[i]); j++ {
				if str0[i-1] == str1[j-1] {
					graph[i][j] = MaxFloat(graph[i-1][j]-gap, graph[i][j-1]-gap, graph[i-1][j-1]+match)
					} else {
						graph[i][j] = MaxFloat(graph[i-1][j]-gap, graph[i][j-1]-gap, graph[i-1][j-1]-mismatch)
					}
				}
			}
				graph[len(graph)-1][len(graph[0])-1] = MaxArrayFloat(graph[len(graph)-1])
			return graph
}

//we provide some helper functions below.

//InitializeScoreTable takes a number of rows and columns as input.
//It returns a matrix of zeroes as floats with appropriate dimensions.
func InitializeScoreTable(numRows, numCols int) [][]float64 {
    scoreTable := make([][]float64, numRows)
    for i := range scoreTable {
        scoreTable[i] = make([]float64, numCols)
    }
    return scoreTable
}

//MaxArrayFloat takes a slice of integers as input and returns the maximum value in the slice.
func MaxArrayFloat(a []float64) float64 {
    if len(a) == 0 {
        panic("Error: array given to MaxArray has zero length.")
    }
    m := a[0]
    for i := 1; i < len(a); i++ {
        if a[i] > m {
            m = a[i]
        }
    }
    return m
}

//MaxFloat is a variadic function that takes an arbitrary collection of floats.
//It returns the maximum one.
func MaxFloat(nums ...float64) float64 {
    m := 0.0
    // nums gets converted to an array
    for i, val := range nums {
        if val > m || i == 0 {
            // update m
            m = val
        }
    }
    return m
}
