package pkg

import "log"

// MakeOct4 は凹型2の８角形を３つの４角形に分割する
func MakeOct4(XY [][]float64, order map[string]int) (cord [][]float64, rect1List [][]float64, rect2List [][]float64, rect3List [][]float64) {
	// octT := "凹型2"
	nodOct := len(XY)
	if nodOct != 8 {
		// TODO:8頂点でない多角形は，三角メッシュ分割
		// 関数から戻る
		return
	}
	num1 := order["L1"]
	// log.Println("num1=", num1)
	// 直交する辺は．L1点と1つ前の点で結ばれる線分
	// 直交する辺の座標ペア
	chokuCord1a := make([][]float64, 2)
	num1P1a := (num1 - 1 + nodOct) % nodOct
	chokuCord1a[0] = XY[num1]
	chokuCord1a[1] = XY[num1P1a]
	// 対向する辺は，L1点から２つ目と３つ目の点で結ばれる線分
	// 対向する辺の座標ペア
	taikoCord1a := make([][]float64, 2)
	num1P2a := (num1 + 2) % nodOct
	taikoCord1a[0] = XY[num1P2a]
	num1P3a := (num1 + 3) % nodOct
	taikoCord1a[1] = XY[num1P3a]
	// 直交する直線1aと対向する辺との直交条件を確認する
	int1aX, int1aY, theta := OrthoAngle(chokuCord1a, taikoCord1a)
	// log.Println("int1aX=", int1aX)
	// log.Println("int1aY=", int1aY)
	// 交差角度が制限範囲内でない場合は処理を中断する
	if theta < 60 || theta > 120 {
		// TODO:関数から戻る
		// return
	}
	// もう一方の直交する辺は．L点と次の点で結ばれる線分
	// 直交する辺の座標ペア
	chokuCord1b := make([][]float64, 2)
	num1N1b := (num1 + 1) % nodOct
	chokuCord1b[0] = XY[num1]
	chokuCord1b[1] = XY[num1N1b]
	// 対向する辺は，L1点から３つ目と４つ目の点で結ばれる線分
	// 対向する辺の座標ペア
	taikoCord1b := make([][]float64, 2)
	num1N3b := (num1 + 3) % nodOct
	taikoCord1b[0] = XY[num1N3b]
	num1N4b := (num1 + 4) % nodOct
	taikoCord1b[1] = XY[num1N4b]
	// 直交する直線1bと対向する辺との直交条件を確認する
	int1bX, int1bY, theta := OrthoAngle(chokuCord1b, taikoCord1b)
	// log.Println("int1bX=", int1bX)
	// log.Println("int1bY=", int1bY)
	// 交差角度が制限範囲内でない場合は処理を中断する
	if theta < 60 || theta > 120 {
		// TODO:関数から戻る
		// return
	}

	num2 := order["L2"]
	// ２つ目の直交する辺は．L2点と1つ前の点で結ばれる線分
	//  直交する辺の座標ペア
	chokuXYa := make([][]float64, 2)
	num2P1a := (num2 - 1 + nodOct) % nodOct
	chokuXYa[0] = XY[num2]
	chokuXYa[1] = XY[num2P1a]
	// 対向する辺は，L2点から４つ目と５つ目の点で結ばれる線分
	// 対向する辺の座標ペア
	taikoXYa := make([][]float64, 2)
	num2P4 := (num2 + 4) % nodOct
	taikoXYa[0] = XY[num2P4]
	num2P5 := (num2 + 5) % nodOct
	taikoXYa[1] = XY[num2P5]
	// 直交する直線2aと対向する辺との直交条件を確認する
	int2aX, int2aY, theta2 := OrthoAngle(chokuXYa, taikoXYa)
	// log.Println("int2aX=", int2aX)
	// log.Println("int2aY=", int2aY)
	// 交差角度が制限範囲内でない場合は処理を中断する
	if theta2 < 60 || theta2 > 120 {
		// TODO:関数から戻る
		// return
	}
	// もう一方の直交する辺は．L2点と次の点で結ばれる線分
	//  直交する辺の座標ペア
	chokuXYb := make([][]float64, 2)
	num2N1b := (num2 + 1) % nodOct
	chokuXYb[0] = XY[num2]
	chokuXYb[1] = XY[num2N1b]
	// 対向する辺は，L2点から５つ目と６つ目の点で結ばれる線分
	// 対向する辺の座標ペア
	taikoXYb := make([][]float64, 2)
	num2N5b := (num2 + 5) % nodOct
	taikoXYb[0] = XY[num2N5b]
	num2N6b := (num2 + 6) % nodOct
	taikoXYb[1] = XY[num2N6b]
	// 直交する直線2bと対向する辺との直交条件を確認する
	int2bX, int2bY, theta2 := OrthoAngle(chokuXYb, taikoXYb)
	// log.Println("int2bX=", int2bX)
	// log.Println("int2bY=", int2bY)
	// 交差角度が制限範囲内でない場合は処理を中断する
	if theta2 < 60 || theta2 > 120 {
		// TODO:関数から戻る
		// return
	}

	// 四角形の頂点のリストを３つ用意する．
	var rect1name []string
	var rect2name []string
	var rect3name []string

	// L点から対向する二辺までの距離を比較する
	// L1点の座標
	// log.Println(XY[num1][1])
	// log.Println(XY[num1][0])
	// 交点1aまでの距離
	divLine1a := DistVerts(XY[num1][1], int1aX, XY[num1][0], int1aY)
	// log.Println("divLine1a=", divLine1a)
	// 交点1bまでの距離
	divLine1b := DistVerts(XY[num1][1], int1bX, XY[num1][0], int1bY)
	// log.Println("divLine1b=", divLine1b)
	// L2点の座標
	// log.Println(XY[num2][1])
	// log.Println(XY[num2][0])
	// 交点2aまでの距離
	divLine2a := DistVerts(XY[num2][1], int2aX, XY[num2][0], int2aY)
	// log.Println("divLine2a=", divLine2a)
	// 交点2bまでの距離
	divLine2b := DistVerts(XY[num2][1], int2bX, XY[num2][0], int2bY)
	// log.Println("divLine2b=", divLine2b)

	// 距離の短い方の線分を分割線とする
	if divLine1a < divLine1b {
		// log.Println("分割線はdivLine1a")
		// 分割点はD1a点（交点１）
		d1 := []float64{int1aY, int1aX}
		// 座標値のリストにD1点の座標値を追加する
		XY = append(XY, d1)
		// log.Println(XY)
		// 頂点並びの辞書に分割点を追加する
		d1num := nodOct
		order["D1"] = d1num
		// log.Println("line1a", order)

		// 四角形L1-R1-R2-D1
		rect1name = []string{"L1", "R1", "R2", "D1"}

		// 距離の短い方の線分を分割線とする
		if divLine2a < divLine2b {
			// log.Println("分割線はdivLine2a")
			// 分割点はD1a点（交点2）
			d2 := []float64{int2aY, int2aX}
			// 座標値のリストにD2点の座標値を追加する
			XY = append(XY, d2)
			// log.Println(XY)
			// 頂点並びの辞書に分割点を追加する
			d2num := nodOct + 1
			order["D2"] = d2num
			// log.Println("line2a", order)

			// 四角形D2-R4-R5-R6
			rect2name = []string{"D2", "R4", "R5", "R6"}
			// 四角形L2-D1-R3-D2
			rect3name = []string{"L2", "D1", "R3", "D2"}
		} else if divLine2a > divLine2b {
			// log.Println("分割線はdivLine2b")
			// 分割点はD1b点（交点2）
			d2 := []float64{int2bY, int2bX}
			// 座標値のリストにD2点の座標値を追加する
			XY = append(XY, d2)
			// log.Println(XY)
			// 頂点並びの辞書に分割点を追加する
			d2num := nodOct + 1
			order["D2"] = d2num
			// log.Println("line2b", order)

			// 四角形L2-D2-R5-R6
			rect2name = []string{"L2", "D2", "R5", "R6"}
			// 四角形D1-R3-R4-D2
			rect3name = []string{"D1", "R3", "R4", "D2"}
		}
	} else if divLine1a > divLine1b {
		// log.Println("分割線はdivLine1b")
		// 分割点はD1b点（交点１）
		d1 := []float64{int1bY, int1bX}
		// 座標値のリストにD1点の座標値を追加する
		XY = append(XY, d1)
		print(XY)
		// 頂点並びの辞書に分割点を追加する
		d1num := nodOct
		order["D1"] = d1num
		// log.Println("line1b", order)

		// 四角形R1-R2-R3-D1
		rect1name = []string{"R1", "R2", "R3", "D1"}

		// 距離の短い方の線分を分割線とする
		if divLine2a < divLine2b {
			// log.Println("分割線はdivLine2a")
			// 分割点はD1a点（交点2）
			d2 := []float64{int2aY, int2aX}
			// 座標値のリストにD2点の座標値を追加する
			XY = append(XY, d2)
			// log.Println(XY)
			// 頂点並びの辞書に分割点を追加する
			d2num := nodOct + 1
			order["D2"] = d2num
			// log.Println("line2a", order)

			// 四角形D2-R4-R5-R6
			rect2name = []string{"D2", "R4", "R5", "R6"}
			// 四角形L2-L1-D1-D2
			rect3name = []string{"L2", "L1", "D1", "D2"}
		} else if divLine2a > divLine2b {
			// log.Println("分割線はdivLine2b")
			// 分割点はD1b点（交点2）
			d2 := []float64{int2bY, int2bX}
			// 座標値のリストにD2点の座標値を追加する
			XY = append(XY, d2)
			// log.Println(XY)
			// 頂点並びの辞書に分割点を追加する
			d2num := nodOct + 1
			order["D2"] = d2num
			// log.Println("line2b", order)

			// 四角形L2-D2-R5-R6
			rect2name = []string{"L2", "D2", "R5", "R6"}
			// 四角形L1-D1-R4-D2
			rect3name = []string{"L1", "D1", "R4", "D2"}
		}
	}

	// 辞書の中身に従ってリストの座標データで四角形を作る
	rect1List = MakeRectList(XY, order, rect1name)
	log.Println("rect1List=", rect1List)
	rect2List = MakeRectList(XY, order, rect2name)
	log.Println("rect2List=", rect2List)
	rect3List = MakeRectList(XY, order, rect3name)
	log.Println("rect3List=", rect3List)

	cord = XY
	// log.Println("cord=", cord)
	return cord, rect1List, rect2List, rect3List
}
