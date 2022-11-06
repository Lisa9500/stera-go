package pkg

// VertsA は直交する辺ａの座標ペアを求める
func VertsA(n int, node int, XY [][]float64) (orthoA [][]float64, distA float64) {
	// 直交する辺は．L点と1つ前の点で結ばれる線分
	orthoA = make([][]float64, 2)
	na := (n - 1 + node) % node
	orthoA[0] = XY[n]
	orthoA[1] = XY[na]

	distA = PairCords(orthoA)
	return orthoA, distA
}
