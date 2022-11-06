package pkg

// MakeRectList は四角形の頂点座標のリストを作成する．
func MakeRectList(XY [][]float64, order map[string]int,
	name []string) (list [][]float64) {
	// 辞書の中身に従ってリストの座標データで四角形を作る
	for _, v := range name {
		// log.Println(i, v)
		n := order[v]
		list = append(list, XY[n])
	}
	return list
}
