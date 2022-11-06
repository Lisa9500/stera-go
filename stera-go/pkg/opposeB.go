package pkg

// OpposeB は直交する辺ａを分割線として対向する辺の頂点ペアを求める
func OpposeB(n int, XY [][]float64, node int) (pairB [][]float64) {
	// pairB := make([][]float64, 2)
	pairB = append(pairB, XY[(n-2+node)%node])
	pairB = append(pairB, XY[(n-3+node)%node])

	return pairB
}
