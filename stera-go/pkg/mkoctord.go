package pkg

// MkOctOrder は8角形の辞書を作り直す．
func MkOctOrder(oct []string, XY [][]float64) (order map[string]int) {
	// 辞書の中身に従ってリストの座標データで8角形を作る
	var octa1L [][]float64
	for _, name := range oct {
		n := order[name]
		octa1L = append(octa1L, XY[n])
	}
	// log.Println("octa1L", octa1L)
	return order
}
