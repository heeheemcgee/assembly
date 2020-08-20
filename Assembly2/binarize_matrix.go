package main

//BinarizeMatrix takes a matrix of values and a threshold.
//It binarizes the matrix according to the threshold.
//If entries across main diagonal are both above threshold, only retain the bigger one.
func BinarizeMatrix(mtx [][]float64, threshold float64) [][]int {
	b := make([][]int, len(mtx))
	for i := range b {
		b[i] = make([]int, len(mtx[i]))
	}
	for i := range b {
		for j := range b[i] {
			if mtx[i][j] >= threshold {
				if mtx[i][j] > mtx[j][i] {
					b[i][j] = 1
				} else if mtx[i][j] == mtx[j][i] && i < j {
					b[i][j] = 1
				}
			}
		}
	}
	return b
}
