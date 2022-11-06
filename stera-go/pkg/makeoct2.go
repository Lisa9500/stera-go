package pkg

import "log"

// MakeOct2 は歯型2の８角形を３つの４角形に分割する
func MakeOct2(XY [][]float64, order map[string]int) (cord [][]float64, rect1List [][]float64, rect2List [][]float64, rect3List [][]float64) {
	// octT := "歯型2"
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
	chokuCord1 := make([][]float64, 2)
	num1P1 := (num1 - 1 + nodOct) % nodOct
	chokuCord1[0] = XY[num1]
	chokuCord1[1] = XY[num1P1]
	// 対向する辺は，L1点から２つ目と３つ目の点で結ばれる線分
	// 対向する辺の座標ペア
	taikoCord1 := make([][]float64, 2)
	num1P2 := (num1 + 2) % nodOct
	taikoCord1[0] = XY[num1P2]
	num1P3 := (num1 + 3) % nodOct
	taikoCord1[1] = XY[num1P3]
	// 直交する直線aと対向する辺との直交条件を確認する
	intX, intY, theta := OrthoAngle(chokuCord1, taikoCord1)
	int1stX := intX
	// log.Println("int1stX=", int1stX)
	int1stY := intY
	// log.Println("int1stY=", int1stY)
	// 交差角度が制限範囲内でない場合は処理を中断する
	if theta < 60 || theta > 120 {
		// TODO:関数から戻る
		// return
	}

	num2 := order["L2"]
	// もう一本の直交する辺は．L2点と1つ次の点で結ばれる線分
	//  直交する辺の座標ペア
	chokuXY := make([][]float64, 2)
	num2N1 := (num2 + 1) % nodOct
	chokuXY[0] = XY[num2]
	chokuXY[1] = XY[num2N1]
	// 対向する辺は，L2点から５つ目と６つ目の点で結ばれる線分
	// 対向する辺の座標ペア
	taikoXY := make([][]float64, 2)
	num2N5 := (num2 + 5) % nodOct
	taikoXY[0] = XY[num2N5]
	num2N6 := (num2 + 6) % nodOct
	taikoXY[1] = XY[num2N6]
	// 直交する直線bと対向する辺との直交条件を確認する
	int2X, int2Y, theta2 := OrthoAngle(chokuXY, taikoXY)
	int2ndX := int2X
	// log.Println("int2ndX=", int2ndX)
	int2ndY := int2Y
	// log.Println("int2ndY=", int2ndY)
	// 交差角度が制限範囲内でない場合は処理を中断する
	if theta2 < 60 || theta2 > 120 {
		// TODO:関数から戻る
		// return
	}

	// 分割点はD1点（交点１）
	d1 := []float64{int1stY, int1stX}
	log.Println("d1=", d1)
	// 座標値のリストにD1点の座標値を追加する
	XY = append(XY, d1)
	// 分割点はD2点（交点２）
	d2 := []float64{int2ndY, int2ndX}
	log.Println("d2=", d2)
	// 座標値のリストにD2点の座標値を追加する
	XY = append(XY, d2)
	// log.Println("cord=", XY)
	// 頂点並びの辞書に分割点を追加する
	d1num := nodOct
	order["D1"] = d1num
	d2num := nodOct + 1
	order["D2"] = d2num
	// log.Println("order=", order)

	// 四角形L2-D2-R4-R5
	rect1name := []string{"L2", "D2", "R4", "R5"}
	// 四角形L1-R1-R2-D1
	rect2name := []string{"L1", "R1", "R2", "D1"}
	// 四角形R6-D1-R3-D2
	rect3name := []string{"R6", "D1", "R3", "D2"}

	// 辞書の中身に従ってリストの座標データで四角形を作る
	// var rect1List [][]float64
	for _, v := range rect1name {
		// log.Println(i, v)
		n := order[v]
		rect1List = append(rect1List, XY[n])
	}
	// log.Println(rect1List)

	// 辞書の中身に従ってリストの座標データで四角形を作る
	rect1List = MakeRectList(XY, order, rect1name)
	log.Println("rect1List=", rect1List)
	rect2List = MakeRectList(XY, order, rect2name)
	log.Println("rect2List=", rect2List)
	rect3List = MakeRectList(XY, order, rect3name)
	log.Println("rect3List=", rect3List)

	// スライスをコピーし，コピーされた要素の個数を返す
	// n := copy(cord, XY)
	// log.Println(n)
	cord = XY
	// log.Println("cord=", cord)
	return cord, rect1List, rect2List, rect3List
}
