package pkg

// VertsB は直交する辺ｂの座標ペアを求める
func VertsB(n int, node int, XY [][]float64) (orthoB [][]float64, distB float64) {
	// もう一方の直交する辺は．L点と次の点で結ばれる線分
	orthoB = make([][]float64, 2)
	nb := (n + 1) % node
	orthoB[0] = XY[n]
	orthoB[1] = XY[nb]

	distB = PairCords(orthoB)
	return orthoB, distB
}
