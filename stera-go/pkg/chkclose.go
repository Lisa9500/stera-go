package pkg

// ChkClose は図形が閉じているかどうかをチェックする
func ChkClose(l int, xy [][]float64) (cls bool) {
	cls = false
	if (xy[0][0] == xy[l-1][0]) && (xy[0][1] == xy[l-1][1]) {
		cls = true
	}
	// log.Println("cls:", cls)
	return cls
}
