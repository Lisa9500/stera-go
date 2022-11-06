package pkg

// LineEquat は直線の方程式を求める．
func LineEquat(x1, y1, x2, y2 float64) (line map[string]float64) {
	// mapオブジェクトを作成，傾きと切片を追加する
	line = map[string]float64{}
	if y1 == y2 {
		// x軸に平行な直線
		line["y"] = y1
	} else if x1 == x2 {
		// y軸に平行な直線
		line["x"] = x1
	} else {
		// y = mx + n
		line["m"] = (y2 - y1) / (x2 - x1)
		line["n"] = y1 - line["m"]*x1
	}
	return line
}
