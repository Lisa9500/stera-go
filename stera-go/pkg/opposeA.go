package pkg

// OpposeA は直交する辺ａを分割線として対向する辺の頂点ペアを求める
func OpposeA(n int, XY [][]float64, node int) (pairA [][]float64) {
	// pairA := make([][]float64, 2)
	pairA = append(pairA, XY[(n+2)%node])
	pairA = append(pairA, XY[(n+3)%node])

	return pairA
}
