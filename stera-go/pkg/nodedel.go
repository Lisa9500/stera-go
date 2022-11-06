package pkg

import (
	"math"
)

// NodDel は近接する頂点を削除する
func NodDel(nod int, XY [][]float64, ext []float64, deg []float64) (nod2 int,
	cord2 [][]float64, ext2 []float64, deg2 []float64) {

	// 削除する頂点のリストを作成する
	var delLst []int
	for i := 0; i < nod; i++ {
		// 次の頂点までの距離を求める
		next := (i + 1) % nod
		chkX := XY[next][1] - XY[i][1]
		chkY := XY[next][0] - XY[i][0]
		// log.Println(i, "X距離", chkX)
		// log.Println(i, "Y距離", chkY)

		// 頂点間の距離が0.1ｍの場合は，前方の頂点を削除する
		if math.Abs(chkX) < 0.1 && math.Abs(chkY) < 0.1 {
			if ext[i] < 0 {
				// 後方が右回りなので削除する
				delLst = append(delLst, i)
				// log.Println("削除ノードリスト1", delLst)
			} else if ext[next] < 0 {
				// 後方が左回り，前方が右回りなので削除する
				delLst = append(delLst, next)
				// log.Println("削除ノードリスト2", delLst)
			} else {
				// 前方，後方が共に左回りなので，後方の座標値を前方との平均にする
				XY[i][1] = XY[i][1] + chkX/2
				XY[i][0] = XY[i][0] + chkY/2
				delLst = append(delLst, next)
				// log.Println("削除ノードリスト3", delLst)
			}
		}
	}

	// 近接する頂点を削除する
	delCnt := len(delLst)
	if delCnt != 0 {
		inCnt := 0
		for i := 0; i < delCnt; i++ {
			// log.Println("削除するノード", delLst[i])
			XY = append(XY[:delLst[i]-inCnt], XY[delLst[i]+1-inCnt:]...)
			ext = append(ext[:delLst[i]-inCnt], ext[delLst[i]+1-inCnt:]...)
			deg = append(deg[:delLst[i]-inCnt], deg[delLst[i]+1-inCnt:]...)
			inCnt++
		}
	}
	nod2 = nod - delCnt
	cord2 = XY
	ext2 = ext
	deg2 = deg
	return nod2, cord2, ext2, deg2
}
