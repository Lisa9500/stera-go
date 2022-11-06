package pkg

// CcwRev は反時計回りを時計回りに並び替える
func CcwRev(num int, xy [][]float64, ext []float64, deg []float64) (revXY [][]float64,
	revExt []float64, revDeg []float64) {
	// log.Println("num", num)

	for i := num - 1; i >= 0; i-- {
		revXY = append(revXY, xy[i])
		revExt = append(revExt, ext[i])
		revDeg = append(revDeg, deg[i])
	}
	// for i := 0; i < num; i++ {
	// log.Println(i, "Y", revXY[i][0])
	// log.Println(i, "X", revXY[i][1])
	// log.Println(i, "外積", revExt[i])
	// log.Println(i, "内角", revDeg[i])
	// }
	return revXY, revExt, revDeg
}
