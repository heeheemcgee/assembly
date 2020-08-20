package main

func GreedyAssembler(reads []string) string {
	genome := reads[0]
	k := len(reads[0])
	reads2 := make([]string, len(reads))
	reads2 = reads

	reads2 = append(reads2[1:])
	for i := 0; i < len(reads2); i++ {
		if reads2[i][:k-1] == genome[len(genome)-k+1:] {
			genome = genome + reads2[i][k-1:]
			reads2 = append(reads2[:i], reads2[i+1:]...)
			i--
		}
	}

	for i := 0; i < len(reads2); i++ {
		if reads2[i][1:] == genome[:k-1] {
			genome = string(reads2[i][0]) + genome
			reads2 = append(reads2[:i], reads2[i+1:]...)
			i--
		}
	}

	return genome
}
