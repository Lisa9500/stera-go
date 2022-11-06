package pkg

// PairCords は頂点ペアのXY座標を求める
func PairCords(pair [][]float64) (dist float64) {
	x1 := pair[0][1]
	y1 := pair[0][0]
	x2 := pair[1][1]
	y2 := pair[1][0]
	dist = DistVerts(x1, x2, y1, y2)
	return dist
}
