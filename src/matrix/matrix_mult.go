package matrix

func MatrixMult[T float64 | uint32 | int](a [][]T, b [][]T) [][]T {
	numRowsA := len(a)
	numColsA := len(a[0])
	numRowsB := len(b)
	numColsB := len(b[0])

	if numColsA != numRowsB {
		panic("Número de colunas da matriz A tem que ser igual ao número de linhas da matriz B")
	}

	result := make([][]T, numRowsA)
	for i := 0; i < numRowsA; i++ {
		result[i] = make([]T, numColsB)
		for j := 0; j < numColsB; j++ {
			sum := T(0)
			for k := 0; k < numColsA; k++ {
				sum += a[i][k] * (b[k][j])
			}
			result[i][j] = sum
		}
	}

	return result
}
