package pkg

import "log"

// HexaDiv は６角形を２つの４角形に分割する
func HexaDiv(XY [][]float64, order map[string]int) (cord [][]float64,
	rect1List [][]float64, rect2List [][]float64) {
	// 頂点データ数の確認
	nodHex := len(XY)
	// log.Println("len(XY)=", nodHex)
	if nodHex != 6 {
		return
	}
	var num int

	// L点の直交条件．対向する辺との交点の角度制限を確認する．
	// var LRkey []string
	var int1stX float64
	var int1stY float64
	var int2ndX float64
	var int2ndY float64
	for LRkey := range order {
		if LRkey == "L1" || LRkey == "L2" {
			// log.Println("LRkey", LRkey)
			num = order[LRkey]
			// log.Println("num", num)

			// 直交する辺は．L点と1つ前の点で結ばれる線分
			// 直交する辺の座標ペア
			chokuCord1 := make([][]float64, 2)
			numP1 := (num - 1 + nodHex) % nodHex
			chokuCord1[0] = XY[num]
			chokuCord1[1] = XY[numP1]
			// 対向する辺は，L点から２つ目と３つ目の点で結ばれる線分
			// 対向する辺の座標ペア
			taikoCord1 := make([][]float64, 2)
			numP2 := (num + 2) % nodHex
			taikoCord1[0] = XY[numP2]
			numP3 := (num + 3) % nodHex
			taikoCord1[1] = XY[numP3]
			// 直交する直線aと対向する辺との直交条件を確認する
			intY, intX, theta := OrthoAngle(chokuCord1, taikoCord1)
			int1stX = intX
			// log.Println("int1stX=", int1stX)
			int1stY = intY
			// log.Println("int1stY=", int1stY)
			// 交差角度が制限範囲内でない場合は処理を中断する
			if theta < 60 || theta > 120 {
				// TODO:折れ曲がりの切妻屋根
				// return
			}

			// もう一方の直交する辺は．L点と1つ次の点で結ばれる線分
			// 直交する辺の座標ペア
			chokuXY := make([][]float64, 2)
			numN1 := (num + 1) % nodHex
			chokuXY[0] = XY[num]
			chokuXY[1] = XY[numN1]
			// もう一方の対向する辺は，L点から３つ目と４つ目の点で結ばれる線分
			// 対向する辺の座標ペア
			taikoXY := make([][]float64, 2)
			numN3 := (num + 3) % nodHex
			taikoXY[0] = XY[numN3]
			numN4 := (num + 4) % nodHex
			taikoXY[1] = XY[numN4]
			// 直交する直線bと対向する辺との直交条件を確認する
			int2X, int2Y, theta2 := OrthoAngle(chokuXY, taikoXY)
			int2ndX = int2X
			// log.Println("int2ndX=", int2ndX)
			int2ndY = int2Y
			// log.Println("int2ndY=", int2ndY)
			// 交差角度が制限範囲内でない場合は処理を中断する
			if theta2 < 60 || theta2 > 120 {
				// TODO:折れ曲がりの切妻屋根
				// return
			}
		} else {
			// TODO:L点がない６角形は六角堂，三角メッシュ分割
			// 関数から戻る
			return
		}
		// log.Println("normal termination")
		continue
	}
	// log.Println("All finished")

	// L点から対向する二辺までの距離を比較する
	// L点の座標
	// log.Println(XY[num][1])
	// log.Println(XY[num][0])
	// 交点１までの距離
	divLa := DistVerts(XY[num][0], int1stX, XY[num][1], int1stY)
	// divLa := DistVerts(XY[num][1], int1stX, XY[num][0], int1stY)
	log.Println("divLa=", divLa)
	// 交点２までの距離
	divLb := DistVerts(XY[num][0], int2ndX, XY[num][1], int2ndY)
	// divLb := DistVerts(XY[num][1], int2ndX, XY[num][0], int2ndY)
	log.Println("divLb=", divLb)

	// 距離の短い方の線分を分割線とする
	var rect1name []string
	var rect2name []string
	if divLa < divLb {
		// log.Println("分割線はdivLa")
		// 分割点はD1点（交点１）
		d1 := []float64{int1stX, int1stY}
		log.Println("d1=", d1)
		// 座標値のリストにD1点の座標値を追加する
		XY = append(XY, d1)
		// log.Println(XY)
		// 頂点並びの辞書に分割点を追加する
		d1Num := nodHex
		order["D1"] = d1Num
		// log.Println("line_a", order)

		// 四角形L1-R1-R2-D1
		rect1name = []string{"L1", "R1", "R2", "D1"}
		// 四角形D1-R3-R4-R5
		rect2name = []string{"D1", "R3", "R4", "R5"}
	} else if divLa > divLb {
		// log.Println("分割線はdivLb")
		// 分割点はD2点（交点２）
		d2 := []float64{int2ndX, int2ndY}
		log.Println("d2=", d2)
		// 座標値のリストにD2点の座標値を追加する
		XY = append(XY, d2)
		// log.Println(XY)
		// 頂点並びの辞書に分割点を追加する
		d2Num := nodHex
		order["D2"] = d2Num
		// log.Println("line_a", order)

		// 四角形D2-R1-R2-R3
		rect1name = []string{"D2", "R1", "R2", "R3"}
		// 四角形L1-D2-R4-R5
		rect2name = []string{"L1", "D2", "R4", "R5"}
	}

	// 辞書の中身に従ってリストの座標データで四角形を作る
	// var rect1List [][]float64
	for _, v := range rect1name {
		// log.Println(i, v)
		n := order[v]
		rect1List = append(rect1List, XY[n])
	}
	// log.Println(rect1List)

	// 辞書の中身に従ってリストの座標データで四角形を作る
	// var rect2List [][]float64
	for _, v := range rect2name {
		// log.Println(i, v)
		n := order[v]
		rect2List = append(rect2List, XY[n])
	}
	// log.Println(rect2List)

	cord = XY
	// log.Println("cord=", cord)
	return cord, rect1List, rect2List
}
