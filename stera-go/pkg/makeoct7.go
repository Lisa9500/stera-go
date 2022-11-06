package pkg

import (
	"log"
	"math"
	"sort"
)

// Pair は分割線の長さの構造体
type Pair struct {
	Key   string
	Value float64
}

// DivlineList は構造体のスライス
type DivlineList []Pair

// Sort関数のインターフェイス
func (p DivlineList) Len() int           { return len(p) }
func (p DivlineList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p DivlineList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// 直交条件を確認するために２点による線分に対してある点の位置を求める
func posline(x1, x0, y1, y0, Y, X float64) (t float64) {
	t = (x1-x0)*(Y-y0) - (y1-y0)*(X-x0)
	return
}

// MakeOct7 はＳ型の８角形を３つの４角形に分割する
func MakeOct7(XY [][]float64, order map[string]int) (cord [][]float64,
	rect1List [][]float64, rect2List [][]float64, rect3List [][]float64,
	hex1L [][]float64) {
	// octT := "Ｓ型"
	nodOct := len(XY)
	if nodOct != 8 {
		// TODO:8頂点でない多角形は，三角メッシュ分割
		// 関数から戻る
		return
	}
	// L1点から頂点2-3，頂点3-4，頂点4-5，頂点5-6の辺との直交対向関係により型分類する
	var flagEdge2to3 bool // Line 1
	var flagEdge3to4 bool // Line 2
	var flagEdge4to5 bool // Line 3
	var flagEdge5to6 bool // Line 4

	// L1点からの分割線による対向する辺との交点までの距離を求める
	num1 := order["L1"]
	// log.Println("num1=", num1)
	// 直交する辺は．L1点と1つ前の点で結ばれる線分
	// 直交する辺の座標ペア（分割線1a）
	chokuCord1ap := make([][]float64, 2)
	num1P1ap := (num1 - 1 + nodOct) % nodOct
	chokuCord1ap[0] = XY[num1]
	chokuCord1ap[1] = XY[num1P1ap]
	// 頂点2-3の辺の座標ペア
	// 対向する辺の座標ペア
	taikoCord1a1 := make([][]float64, 2)
	num1P2a1 := (num1 + 2) % nodOct
	taikoCord1a1[0] = XY[num1P2a1]
	num1P3a1 := (num1 + 3) % nodOct
	taikoCord1a1[1] = XY[num1P3a1]
	// 直交する直線1aと対向する辺との交点を求める
	int1a1X, int1a1Y, theta1a1 := OrthoAngle(chokuCord1ap, taikoCord1a1)
	// log.Println("int1a1X=", int1a1X)
	// log.Println("int1a1Y=", int1a1Y)
	// log.Println("theta1a1=", theta1a1)
	// 交差角度が制限範囲内かどうか確認する
	if theta1a1 > 60 || theta1a1 < 120 {
		// 対向する辺の頂点が直行する直線に対して同じ側にある場合
		// 交点は対向する辺上にない．
		t1 := posline(chokuCord1ap[1][1], chokuCord1ap[0][1], chokuCord1ap[1][0],
			chokuCord1ap[0][0], taikoCord1a1[0][0], taikoCord1a1[0][1])
		t2 := posline(chokuCord1ap[1][1], chokuCord1ap[0][1], chokuCord1ap[1][0],
			chokuCord1ap[0][0], taikoCord1a1[1][0], taikoCord1a1[1][1])
		if t1*t2 < 0 {
			// L1点から２つ目-３つ目の辺の直交対向条件を満たす
			flagEdge2to3 = true
			// log.Println("flagEdge2to3", flagEdge2to3)
		} else {
			// TODO:関数から戻る
			// return
		}
	}
	// 交点1a1までの距離
	divLine1a1 := DistVerts(XY[num1][0], int1a1X, XY[num1][1], int1a1Y)
	// log.Println("XY[num1][1]=", XY[num1][1])
	// log.Println("int1a1X=", int1a1X)
	// log.Println("XY[num1][0]=", XY[num1][0])
	// log.Println("int1a1Y=", int1a1Y)
	// log.Println("divLine1a1=", divLine1a1)

	// 頂点4-5の辺の座標ペア
	// 対向する辺の座標ペア
	taikoCord1a3 := make([][]float64, 2)
	num1P4a3 := (num1 + 4) % nodOct
	taikoCord1a3[0] = XY[num1P4a3]
	num1P5a3 := (num1 + 5) % nodOct
	taikoCord1a3[1] = XY[num1P5a3]
	// 直交する直線1aと対向する辺との交点を求める
	int1a3X, int1a3Y, theta1a3 := OrthoAngle(chokuCord1ap, taikoCord1a3)
	// log.Println("int1a3X=", int1a3X)
	// log.Println("int1a3Y=", int1a3Y)
	// log.Println("theta1a3=", theta1a3)
	// 交差角度が制限範囲内かどうか確認する
	if theta1a3 > 60 || theta1a3 < 120 {
		// 対向する辺の頂点が直行する直線に対して同じ側にある場合
		// 交点は対向する辺上にない．
		t1 := posline(chokuCord1ap[1][1], chokuCord1ap[0][1], chokuCord1ap[1][0],
			chokuCord1ap[0][0], taikoCord1a3[0][0], taikoCord1a3[0][1])
		t2 := posline(chokuCord1ap[1][1], chokuCord1ap[0][1], chokuCord1ap[1][0],
			chokuCord1ap[0][0], taikoCord1a3[1][0], taikoCord1a3[1][1])
		if t1*t2 < 0 {
			// L1点から４つ目-５つ目の辺の直交対向条件を満たす
			flagEdge4to5 = true
			// log.Println("flagEdge4to5", flagEdge4to5)
		} else {
			// TODO:関数から戻る
			// return
		}
	}
	// 交点1a3までの距離
	divLine1a3 := DistVerts(XY[num1][0], int1a3X, XY[num1][1], int1a3Y)
	// log.Println("XY[num1][0]=", XY[num1][0])
	// log.Println("int1a3X=", int1a3X)
	// log.Println("XY[num1][1]=", XY[num1][1])
	// log.Println("int1a3Y=", int1a3Y)
	// log.Println("divLine1a3=", divLine1a3)

	// もう一方の直交する辺は．L点と次の点で結ばれる線分
	// 直交する辺の座標ペア（分割線1b）
	chokuCord1bn := make([][]float64, 2)
	chokuCord1bn[0] = XY[num1]
	chokuCord1bn[1] = XY[(num1+1)%nodOct]
	// 頂点3-4の辺の座標ペア
	// 対向する辺の座標ペア
	taikoCord1b2 := make([][]float64, 2)
	taikoCord1b2[0] = XY[(num1+3)%nodOct]
	taikoCord1b2[1] = XY[(num1+4)%nodOct]
	// 直交する直線1bと対向する辺との交点を求める
	int1b2X, int1b2Y, theta1b2 := OrthoAngle(chokuCord1bn, taikoCord1b2)
	// log.Println("int1b2X=", int1b2X)
	// log.Println("int1b2Y=", int1b2Y)
	// log.Println("theta1b2=", theta1b2)
	// 交差角度が制限範囲内かどうか確認する
	if theta1b2 > 60 || theta1b2 < 120 {
		// 対向する辺の頂点が直行する直線に対して同じ側にある場合
		// 交点は対向する辺上にない．
		t1 := posline(chokuCord1bn[1][1], chokuCord1bn[0][1], chokuCord1bn[1][0],
			chokuCord1bn[0][0], taikoCord1b2[0][0], taikoCord1b2[0][1])
		t2 := posline(chokuCord1bn[1][1], chokuCord1bn[0][1], chokuCord1bn[1][0],
			chokuCord1bn[0][0], taikoCord1b2[1][0], taikoCord1b2[1][1])
		if t1*t2 < 0 {
			// L1点から３つ目-４つ目の辺の直交対向条件を満たす
			flagEdge3to4 = true
			// log.Println("flagEdge3to4", flagEdge3to4)
		} else {
			// TODO:関数から戻る
			// return
		}
	}
	// 交点1b2までの距離
	divLine1b2 := DistVerts(XY[num1][0], int1b2X, XY[num1][1], int1b2Y)
	// log.Println("XY[num1][0]=", XY[num1][0])
	// log.Println("int1b2X=", int1b2X)
	// log.Println("XY[num1][1]=", XY[num1][1])
	// log.Println("int1b2Y=", int1b2Y)
	// log.Println("divLine1b2=", divLine1b2)

	// 頂点5-6の辺の座標ペア
	// 対向する辺の座標ペア
	taikoCord1b4 := make([][]float64, 2)
	taikoCord1b4[0] = XY[(num1+5)%nodOct]
	taikoCord1b4[1] = XY[(num1+6)%nodOct]
	// 直交する直線1bと対向する辺との交点を求める
	int1b4X, int1b4Y, theta1b4 := OrthoAngle(chokuCord1bn, taikoCord1b4)
	// log.Println("int1b4X=", int1b4X)
	// log.Println("int1b4Y=", int1b4Y)
	// log.Println("theta1b4=", theta1b4)
	// 交差角度が制限範囲内かどうか確認する
	if theta1b4 > 60 || theta1b4 < 120 {
		// 対向する辺の頂点が直行する直線に対して同じ側にある場合
		// 交点は対向する辺上にない．
		t1 := posline(chokuCord1bn[1][1], chokuCord1bn[0][1], chokuCord1bn[1][0],
			chokuCord1bn[0][0], taikoCord1b4[0][0], taikoCord1b4[0][1])
		t2 := posline(chokuCord1bn[1][1], chokuCord1bn[0][1], chokuCord1bn[1][0],
			chokuCord1bn[0][0], taikoCord1b4[1][0], taikoCord1b4[1][1])
		if t1*t2 < 0 {
			// L1点から５つ目-６つ目の辺の直交対向条件を満たす
			flagEdge5to6 = true
			// log.Println("flagEdge5to6", flagEdge5to6)
		} else {
			// TODO:関数から戻る
			// return
		}
	}
	// 交点1b4までの距離
	divLine1b4 := DistVerts(XY[num1][0], int1b4X, XY[num1][1], int1b4Y)
	// log.Println("XY[num1][0]=", XY[num1][0])
	// log.Println("int1b4X=", int1b4X)
	// log.Println("XY[num1][1]=", XY[num1][1])
	// log.Println("int1b4Y=", int1b4Y)
	// log.Println("divLine1b4=", divLine1b4)

	// 四角形の頂点のリストを３つ用意する．
	var rect1name []string
	var rect2name []string
	var rect3name []string
	var hex1name []string

	// L2点からの分割線による対向する辺との交点までの距離を求める
	num2 := order["L2"]
	// 対抗する辺との直交条件から型分類を行う
	if (flagEdge2to3 == true) && (flagEdge3to4 == true) {
		octStype := "typeA_S"
		log.Println(octStype)
		// L2点と対向する辺との交点を求める
		// １つ目の直交する辺は．L2点と1つ前の点で結ばれる線分
		// 直交する辺の座標ペア（分割線2a）
		chokuXYap := make([][]float64, 2)
		chokuXYap[0] = XY[num2]
		chokuXYap[1] = XY[(num2-1+nodOct)%nodOct]
		// 頂点2-3の辺の座標ペア
		// 対向する辺の座標ペア
		taikoXYa1 := make([][]float64, 2)
		taikoXYa1[0] = XY[(num2+2)%nodOct]
		taikoXYa1[1] = XY[(num2+3)%nodOct]
		// 直交する直線2aと対向する辺との交点を求める
		int2a1X, int2a1Y, theta2a1 := OrthoAngle(chokuXYap, taikoXYa1)
		// log.Println("int2a1X=", int2a1X)
		// log.Println("int2a1Y=", int2a1Y)
		// log.Println("theta2a1=", theta2a1)
		// 交差角度が制限範囲内かどうか確認する
		divLine2a1 := math.Inf(1)
		if theta2a1 > 60 || theta2a1 < 120 {
			// 対向する辺の頂点が直行する直線に対して同じ側にある場合
			// 交点は対向する辺上にない．
			t1 := posline(chokuXYap[1][1], chokuXYap[0][1], chokuXYap[1][0],
				chokuXYap[0][0], taikoXYa1[0][0], taikoXYa1[0][1])
			t2 := posline(chokuXYap[1][1], chokuXYap[0][1], chokuXYap[1][0],
				chokuXYap[0][0], taikoXYa1[1][0], taikoXYa1[1][1])
			if t1*t2 < 0 {
				// 交点2a_2_3までの距離
				divLine2a1 = DistVerts(XY[num2][0], int2a1X, XY[num2][1], int2a1Y)
				// log.Println("XY[num2][0]=", XY[num2][0])
				// log.Println("int2a1X=", int2a1X)
				// log.Println("XY[num2][1]=", XY[num2][1])
				// log.Println("int2a1Y=", int2a1Y)
				// log.Println("divLine2a1", divLine2a1)
			} else {
				// TODO:関数から戻る
				// return
			}
			// log.Println("divLine2a1", divLine2a1)
		}
		// もう一方の直交する辺は．L2点と次の点で結ばれる線分
		// 直交する辺の座標ペア（分割線2b）
		chokuXYbn := make([][]float64, 2)
		chokuXYbn[0] = XY[num2]
		chokuXYbn[1] = XY[(num2+1)%nodOct]
		// 頂点3-4の辺の座標ペア
		// 対向する辺の座標ペア
		taikoXYb2 := make([][]float64, 2)
		taikoXYb2[0] = XY[(num2+3)%nodOct]
		taikoXYb2[1] = XY[(num2+4)%nodOct]
		// 直交する直線2bと対向する辺との交点を求める
		int2b2X, int2b2Y, theta2b2 := OrthoAngle(chokuXYbn, taikoXYb2)
		// log.Println("int2b2X=", int2b2X)
		// log.Println("int2b2Y=", int2b2Y)
		// log.Println("theta2b2=", theta2b2)
		// 交差角度が制限範囲内かどうか確認する
		divLine2b2 := math.Inf(1)
		if theta2b2 > 60 || theta2b2 < 120 {
			// 対向する辺の頂点が直行する直線に対して同じ側にある場合
			// 交点は対向する辺上にない．
			t1 := posline(chokuXYap[1][1], chokuXYap[0][1], chokuXYap[1][0],
				chokuXYap[0][0], taikoXYb2[0][0], taikoXYb2[0][1])
			t2 := posline(chokuXYap[1][1], chokuXYap[0][1], chokuXYap[1][0],
				chokuXYap[0][0], taikoXYb2[1][0], taikoXYb2[1][1])
			if t1*t2 < 0 {
				// 交点2b_5_6までの距離
				divLine2b2 = DistVerts(XY[num2][0], int2b2X, XY[num2][1], int2b2Y)
				// log.Println("XY[num2][0]=", XY[num2][0])
				// log.Println("int2b2X=", int2b2X)
				// log.Println("XY[num2][1]=", XY[num2][1])
				// log.Println("int2b2Y=", int2b2Y)
				// log.Println("divLine2b2", divLine2b2)
			} else {
				// TODO:関数から戻る
				// return
			}
			// log.Println("divLine2b2", divLine2b2)
		}
		// divLine2a1 := DistVerts(XY[num2][1], int2a1X, XY[num2][0], int2a1Y)
		// divLine2b2 := DistVerts(XY[num2][1], int2b2X, XY[num2][0], int2b2Y)

		// 分割線1aと1bを比較する，分割線2aと2bを比較する
		// 距離の短い方の線分を分割線とする
		if divLine1a1 < divLine1b2 {
			// log.Println("分割線はdivLine1a")
			// 分割点はD1a点（交点１）
			d1a := []float64{int1a1Y, int1a1X}
			// 座標値のリストにD1a点の座標値を追加する
			XY = append(XY, d1a)
			// log.Println(XY)
			// 頂点並びの辞書に分割点を追加する
			d1anum := nodOct
			order["D1a"] = d1anum
			// log.Println("line1a", order)

			// 四角形L1-R1-R2-D1a
			rect1name = []string{"L1", "R1", "R2", "D1a"}
			// log.Println("rect1name", rect1name)

			// 距離の短い方の線分を分割線とする
			if divLine2a1 < divLine2b2 {
				// log.Println("分割線はdivLine2a")
				// 分割点はD2a点（交点2）
				d2a := []float64{int2a1Y, int2a1X}
				// 座標値のリストにD2点の座標値を追加する
				XY = append(XY, d2a)
				// log.Println(XY)
				// 頂点並びの辞書に分割点を追加する
				d2anum := nodOct + 1
				order["D2a"] = d2anum
				// log.Println("line2a", order)

				// 四角形D1a-R3-D2a-R6
				rect2name = []string{"D1a", "R3", "D2a", "R6"}
				// log.Println("rect2name=", rect2name)
				// 四角形L2-R4-R5-D2a
				rect3name = []string{"L2", "R4", "R5", "D2a"}
				// log.Println("rect3name=", rect3name)

			} else if divLine2a1 > divLine2b2 {
				// log.Println("分割線はdivLine2b")
				// 分割点はD2b点（交点2）
				d2b := []float64{int2b2Y, int2b2X}
				// 座標値のリストにD2b点の座標値を追加する
				XY = append(XY, d2b)
				// log.Println(XY)
				// 頂点並びの辞書に分割点を追加する
				d2bnum := nodOct + 1
				order["D2b"] = d2bnum
				// log.Println("line2b", order)

				// 四角形D1a-R3-L2-D2b
				rect2name = []string{"D1a", "R3", "L2", "D2b"}
				// log.Println("rect2name=", rect2name)
				// 四角形D2b-R4-R5-R6
				rect3name = []string{"D2b", "R4", "R5", "R6"}
				// log.Println("rect3name=", rect3name)
			}
		} else if divLine1a1 > divLine1b2 {
			// log.Println("分割線はdivLine1b")
			// 分割点はD1b点（交点２）
			d1b := []float64{int1b2Y, int1b2X}
			// 座標値のリストにD1b点の座標値を追加する
			XY = append(XY, d1b)
			// log.Println(XY)
			// 頂点並びの辞書に分割点を追加する
			d1bnum := nodOct
			order["D1b"] = d1bnum
			// log.Println("line1b", order)

			// 四角形R1-R2-R3-D1b
			rect1name = []string{"R1", "R2", "R3", "D1b"}
			// log.Println("rect1name=", rect1name)

			// 距離の短い方の線分を分割線とする
			if divLine2a1 < divLine2b2 {
				// log.Println("分割線はdivLine2a")
				// 分割点はD2a点（交点２）
				d2a := []float64{int2a1Y, int2a1X}
				// 座標値のリストにD2a点の座標値を追加する
				XY = append(XY, d2a)
				// log.Println(XY)
				// 頂点並びの辞書に分割点を追加する
				d2anum := nodOct + 1
				order["D2a"] = d2anum
				// log.Println("line2a", order)

				// 四角形L1-D1b-D2a-R6
				rect2name = []string{"L1", "D1b", "D2a", "R6"}
				// log.Println("rect2name=", rect2name)
				// 四角形L2-R4-R5-D2a
				rect3name = []string{"L2", "R4", "R5", "D2a"}
				// log.Println("rect3name=", rect3name)
			} else if divLine2a1 > divLine2b2 {
				// log.Println("分割線はdivLine2b")
				// 分割点はD2b点（交点２）
				d2b := []float64{int2b2Y, int2b2X}
				// 座標値のリストにD2b点の座標値を追加する
				XY = append(XY, d2b)
				// log.Println(XY)
				// 頂点並びの辞書に分割点を追加する
				d2bnum := nodOct + 1
				order["D2b"] = d2bnum
				// log.Println("line2b", order)

				// 四角形L1-D1b-L2-D2b
				rect2name = []string{"L1", "D1b", "L2", "D2b"}
				// log.Println("rect2name=", rect2name)
				// 四角形D2b-R4-R5-R6
				rect3name = []string{"D2b", "R4", "R5", "R6"}
				// log.Println("rect3name=", rect3name)
			}
		}
		// L点と分割点が共に近接する場合，真ん中の四角形が非常に小さく細長くなる．
		// この場合は，2つのL点を結ぶ線を分割線とする．
		// L1点とD2b点の間の距離
		distL1D2b := DistVerts(XY[num1][0], int2b2X, XY[num1][1], int2b2Y)
		// log.Println("XY[num1][0]=", XY[num1][0])
		// log.Println("int2b2X=", int2b2X)
		// log.Println("XY[num1][1]=", XY[num1][1])
		// log.Println("int2b2Y=", int2b2Y)
		// L2点とD1b点の間の距離
		distL2D1b := DistVerts(XY[num2][0], int1b2X, XY[num2][1], int1b2Y)
		// log.Println("XY[num2][0]=", XY[num2][0])
		// log.Println("int1b2X=", int1b2X)
		// log.Println("XY[num2][1]=", XY[num2][1])
		// log.Println("int1b2Y=", int1b2Y)

		// L点と分割点の間の距離が共に短い場合は，L点を結ぶ線を分割線とする．
		if distL1D2b < 1.0 && distL2D1b < 1.0 {
			// ２つの四角形に分割する．なお，中間点となるL点は使用しない．
			rect1name = []string{"R1", "R2", "R3", "L2"}
			// log.Println("rect1name", rect1name)
			rect2name = []string{"L1", "R4", "R5", "R6"}
			// log.Println("rect2name", rect2name)
			// 辞書の中身に従ってリストの座標データで四角形を作る
			rect1List = MakeRectList(XY, order, rect1name)
			log.Println("rect1List=", rect1List)
			rect2List = MakeRectList(XY, order, rect2name)
			log.Println("rect2List=", rect2List)
		} else if distL1D2b > 1.0 || distL2D1b > 1.0 {
			// 辞書の中身に従ってリストの座標データで四角形を作る
			rect1List = MakeRectList(XY, order, rect1name)
			log.Println("rect1List=", rect1List)
			rect2List = MakeRectList(XY, order, rect2name)
			log.Println("rect2List=", rect2List)
			rect3List = MakeRectList(XY, order, rect3name)
			log.Println("rect3List=", rect3List)
		}
	} else if (flagEdge4to5 == true) && (flagEdge5to6 == true) {
		octStype := "typeB_S"
		log.Println(octStype)
		// L2点と対向する辺との交点を求める
		// １つ目の直交する辺は．L2点と1つ前の点で結ばれる線分
		// 直交する辺の座標ペア（分割線2a）
		chokuXYap := make([][]float64, 2)
		chokuXYap[0] = XY[num2]
		chokuXYap[1] = XY[(num2-1+nodOct)%nodOct]
		// 頂点4-5の辺の座標ペア
		// 対向する辺の座標ペア
		taikoXYa3 := make([][]float64, 2)
		taikoXYa3[0] = XY[(num2+4)%nodOct]
		taikoXYa3[1] = XY[(num2+5)%nodOct]
		// 直交する直線2aと対向する辺との交点を求める
		int2a3X, int2a3Y, theta2a3 := OrthoAngle(chokuXYap, taikoXYa3)
		// log.Println("int2a3X=", int2a3X)
		// log.Println("int2a3Y=", int2a3Y)
		// log.Println("theta2a3=", theta2a3)
		// 交差角度が制限範囲内かどうか確認する
		divLine2a3 := math.Inf(1)
		if theta2a3 > 60 || theta2a3 < 120 {
			// 対向する辺の頂点が直行する直線に対して同じ側にある場合
			// 交点は対向する辺上にない．
			t1 := posline(chokuXYap[1][1], chokuXYap[0][1], chokuXYap[1][0],
				chokuXYap[0][0], taikoXYa3[0][0], taikoXYa3[0][1])
			t2 := posline(chokuXYap[1][1], chokuXYap[0][1], chokuXYap[1][0],
				chokuXYap[0][0], taikoXYa3[1][0], taikoXYa3[1][1])
			if t1*t2 < 0 {
				// 交点2a_4_5までの距離
				divLine2a3 = DistVerts(XY[num2][0], int2a3X, XY[num2][1], int2a3Y)
				// log.Println("XY[num2][0]=", XY[num2][0])
				// log.Println("int2a3X=", int2a3X)
				// log.Println("XY[num2][1]=", XY[num2][1])
				// log.Println("int2a3Y=", int2a3Y)
				// log.Println("divLine2a3", divLine2a3)
			} else {
				// TODO:関数から戻る
				// return
			}
		}
		// もう一方の直交する辺は．L2点と次の点で結ばれる線分
		// 直交する辺の座標ペア（分割線2b）
		chokuXYbn := make([][]float64, 2)
		chokuXYbn[0] = XY[num2]
		chokuXYbn[1] = XY[(num2+1)%nodOct]
		// 頂点5-6の辺の座標ペア
		// 対向する辺の座標ペア
		taikoXYb4 := make([][]float64, 2)
		taikoXYb4[0] = XY[(num2+5)%nodOct]
		taikoXYb4[1] = XY[(num2+6)%nodOct]
		// 直交する直線2bと対向する辺との交点を求める
		int2b4X, int2b4Y, theta2b4 := OrthoAngle(chokuXYbn, taikoXYb4)
		// log.Println("int2b4X=", int2b4X)
		// log.Println("int2b4Y=", int2b4Y)
		// log.Println("theta2b4=", theta2b4)
		// 交差角度が制限範囲内かどうか確認する
		divLine2b4 := math.Inf(1)
		if theta2b4 > 60 || theta2b4 < 120 {
			// 対向する辺の頂点が直行する直線に対して同じ側にある場合
			// 交点は対向する辺上にない．
			t1 := posline(chokuXYap[1][1], chokuXYap[0][1], chokuXYap[1][0],
				chokuXYap[0][0], taikoXYb4[0][0], taikoXYb4[0][1])
			t2 := posline(chokuXYap[1][1], chokuXYap[0][1], chokuXYap[1][0],
				chokuXYap[0][0], taikoXYb4[1][0], taikoXYb4[1][1])
			if t1*t2 < 0 {
				// 交点2b_5_6までの距離
				divLine2b4 = DistVerts(XY[num2][0], int2b4X, XY[num2][1], int2b4Y)
				// log.Println("XY[num2][0]=", XY[num2][0])
				// log.Println("int2b4X=", int2b4X)
				// log.Println("XY[num2][1]=", XY[num2][1])
				// log.Println("int2b4Y=", int2b4Y)
				// log.Println("divLine2b4", divLine2b4)
			} else {
				// TODO:関数から戻る
				// return
			}
		}
		// divLine2a3 := DistVerts(XY[num2][1], int2a3X, XY[num2][0], int2a3Y)
		// divLine2b4 := DistVerts(XY[num2][1], int2b4X, XY[num2][0], int2b4Y)

		// 分割線1aと1bを比較する，分割線2aと2bを比較する
		// 距離の短い方の線分を分割線とする
		if divLine1a3 < divLine1b4 {
			// log.Println("分割線はdivLine1a")
			// 分割点はD1a点（交点１）
			d1a := []float64{int1a3Y, int1a3X}
			// 座標値のリストにD1a点の座標値を追加する
			XY = append(XY, d1a)
			// log.Println(XY)
			// 頂点並びの辞書に分割点を追加する
			d1anum := nodOct
			order["D1a"] = d1anum
			// log.Println("line1a", order)

			// 四角形D1a-R4-R5-R6
			rect1name = []string{"D1a", "R4", "R5", "R6"}
			// log.Println("rect1name", rect1name)

			// 距離の短い方の線分を分割線とする
			if divLine2a3 < divLine2b4 {
				// log.Println("分割線はdivLine2a")
				// 分割点はD2a点（交点2）
				d2a := []float64{int2a3Y, int2a3X}
				// 座標値のリストにD2a点の座標値を追加する
				XY = append(XY, d2a)
				// log.Println(XY)
				// 頂点並びの辞書に分割点を追加する
				d2anum := nodOct + 1
				order["D2a"] = d2anum
				// log.Println("line2a", order)

				// 四角形D1a-L1-D2a-L2
				rect2name = []string{"D1a", "L1", "D2a", "L2"}
				// log.Println("rect2name", rect2name)
				// 四角形R1-R2-R3-D2a
				rect3name = []string{"R1", "R2", "R3", "D2a"}
				// log.Println("rect3name", rect3name)

			} else if divLine2a3 > divLine2b4 {
				// log.Println("分割線はdivLine2b")
				// 分割点はD2b点（交点２）
				d2b := []float64{int2b4Y, int2b4X}
				// 座標値のリストにD2b点の座標値を追加する
				XY = append(XY, d2b)
				// log.Println(XY)
				// 頂点並びの辞書に分割点を追加する
				d2bnum := nodOct + 1
				order["D2b"] = d2bnum
				// log.Println("line2b", order)

				// 四角形R2-R3-L2-D2b
				rect2name = []string{"R2", "R3", "L2", "D2b"}
				// log.Println("rect2name", rect2name)
				// 四角形R1-D2b-D1a-L1
				rect3name = []string{"R1", "D2b", "D1a", "L1"}
				// log.Println("rect3name", rect3name)
			}
		} else if divLine1a3 > divLine1b4 {
			// log.Println("分割線はdivLine1b")
			// 分割点はD1b点（交点２）
			d1b := []float64{int1b4Y, int1b4X}
			// 座標値のリストにD1b点の座標値を追加する
			XY = append(XY, d1b)
			// log.Println(XY)
			// 頂点並びの辞書に分割点を追加する
			d1bnum := nodOct
			order["D1b"] = d1bnum
			// log.Println("line1b", order)

			// 四角形L1-D1b-R5-R6
			rect1name = []string{"L1", "D1b", "R5", "R6"}
			// log.Println("rect1name", rect1name)

			// 距離の短い方の線分を分割線とする
			if divLine2a3 < divLine2b4 {
				// log.Println("分割線はdivLine2a")
				// 分割点はD2a点（交点２）
				d2a := []float64{int2a3Y, int2a3X}
				// 座標値のリストにD2a点の座標値を追加する
				XY = append(XY, d2a)
				// log.Println(XY)
				// 頂点並びの辞書に分割点を追加する
				d2anum := nodOct + 1
				order["D2a"] = d2anum
				// log.Println("line2a", order)

				// 四角形D1b-D2a-L2-R4
				rect2name = []string{"D1b", "D2a", "L2", "R4"}
				// log.Println("rect2name", rect2name)
				// 四角形R1-R2-R3-D2a
				rect3name = []string{"R1", "R2", "R3", "D2a"}
				// log.Println("rect3name", rect3name)
			} else if divLine2a3 > divLine2b4 {
				// log.Println("分割線はdivLine2b")
				// 分割点はD2b点（交点２）
				d2b := []float64{int2b4Y, int2b4X}
				// 座標値のリストにD2b点の座標値を追加する
				XY = append(XY, d2b)
				// log.Println(XY)
				// 頂点並びの辞書に分割点を追加する
				d2bnum := nodOct + 1
				order["D2b"] = d2bnum
				// log.Println("line2b", order)

				// 四角形R1-D2b-R4-D1b
				rect2name = []string{"R1", "D2b", "R4", "D1b"}
				// log.Println("rect2name", rect2name)
				// 四角形R2-R3-L2-D2b
				rect3name = []string{"R2", "R3", "L2", "D2b"}
				// log.Println("rect3name", rect3name)
			}
		}
		// L点と分割点が共に近接する場合，真ん中の四角形が非常に小さく細長くなる．
		// この場合は，2つのL点を結ぶ線を分割線とする．
		// L1点とD2a点の間の距離
		distL1D2a := DistVerts(XY[num1][0], int2a3X, XY[num1][1], int2a3Y)
		// log.Println("XY[num1][0]=", XY[num1][0])
		// log.Println("int2a3X=", int2a3X)
		// log.Println("XY[num1][1]=", XY[num1][1])
		// log.Println("int2a3Y=", int2a3Y)
		// L2点とD1b点の間の距離
		distL2D1a := DistVerts(XY[num2][0], int1a3X, XY[num2][1], int1a3Y)
		// log.Println("XY[num2][0]=", XY[num2][0])
		// log.Println("int1a3X=", int1a3X)
		// log.Println("XY[num2][1]=", XY[num2][1])
		// log.Println("int1a3Y=", int1a3Y)

		// L点と分割点の間の距離が共に短い場合は，L点を結ぶ線を分割線とする．
		if distL1D2a < 1.0 && distL2D1a < 1.0 {
			// ２つの四角形に分割する．なお，中間点となるL点は使用しない．
			rect1name = []string{"R1", "R2", "R3", "L1"}
			// log.Println("rect1name", rect1name)
			rect2name = []string{"L2", "R4", "R5", "R6"}
			// log.Println("rect2name", rect2name)
			// 辞書の中身に従ってリストの座標データで四角形を作る
			rect1List = MakeRectList(XY, order, rect1name)
			log.Println("rect1List=", rect1List)
			rect2List = MakeRectList(XY, order, rect2name)
			log.Println("rect2List=", rect2List)
		} else if distL1D2a > 1.0 || distL2D1a > 1.0 {
			// 辞書の中身に従ってリストの座標データで四角形を作る
			rect1List = MakeRectList(XY, order, rect1name)
			log.Println("rect1List=", rect1List)
			rect2List = MakeRectList(XY, order, rect2name)
			log.Println("rect2List=", rect2List)
			rect3List = MakeRectList(XY, order, rect3name)
			log.Println("rect3List=", rect3List)
		}
		// 対抗する辺との直交条件から型分類を行う
	} else if (flagEdge2to3 == true) && (flagEdge5to6 == true) {
		octStype := "Rotation_S"
		log.Println(octStype)
		// L2点と対向する辺との交点を求める
		// １つ目の直交する辺は．L2点と1つ前の点で結ばれる線分
		// 直交する辺の座標ペア（分割線2a）
		chokuXYap := make([][]float64, 2)
		chokuXYap[0] = XY[num2]
		chokuXYap[1] = XY[(num2-1+nodOct)%nodOct]
		// 頂点2-3の辺の座標ペア
		// 対向する辺の座標ペア
		taikoXYa1 := make([][]float64, 2)
		taikoXYa1[0] = XY[(num2+2)%nodOct]
		taikoXYa1[1] = XY[(num2+3)%nodOct]
		// 直交する直線2aと対向する辺との交点を求める
		int2a1X, int2a1Y, theta2a1 := OrthoAngle(chokuXYap, taikoXYa1)
		// log.Println("int2a1X=", int2a1X)
		// log.Println("int2a1Y=", int2a1Y)
		// log.Println("theta2a1=", theta2a1)
		// 交差角度が制限範囲内かどうか確認する
		divLine2a1 := math.Inf(1)
		if theta2a1 > 60 || theta2a1 < 120 {
			// 対向する辺の頂点が直行する直線に対して同じ側にある場合
			// 交点は対向する辺上にない．
			t1 := posline(chokuXYap[1][1], chokuXYap[0][1], chokuXYap[1][0],
				chokuXYap[0][0], taikoXYa1[0][0], taikoXYa1[0][1])
			t2 := posline(chokuXYap[1][1], chokuXYap[0][1], chokuXYap[1][0],
				chokuXYap[0][0], taikoXYa1[1][0], taikoXYa1[1][1])
			if t1*t2 < 0 {
				// 交点2a_2_3までの距離
				divLine2a1 = DistVerts(XY[num2][0], int2a1X, XY[num2][1], int2a1Y)
				// log.Println("XY[num2][0]=", XY[num2][0])
				// log.Println("int2a1X=", int2a1X)
				// log.Println("XY[num2][1]=", XY[num2][1])
				// log.Println("int2a1Y=", int2a1Y)
				// log.Println("divLine2a1", divLine2a1)
			} else {
				// TODO:関数から戻る
				// return
			}
		}
		// もう一方の直交する辺は．L2点と次の点で結ばれる線分
		// 直交する辺の座標ペア（分割線2b）
		chokuXYbn := make([][]float64, 2)
		chokuXYbn[0] = XY[num2]
		chokuXYbn[1] = XY[(num2+1)%nodOct]
		// 頂点5-6の辺の座標ペア
		// 対向する辺の座標ペア
		taikoXYb4 := make([][]float64, 2)
		taikoXYb4[0] = XY[(num2+5)%nodOct]
		taikoXYb4[1] = XY[(num2+6)%nodOct]
		// 直交する直線2bと対向する辺との交点を求める
		int2b4X, int2b4Y, theta2b4 := OrthoAngle(chokuXYbn, taikoXYb4)
		// log.Println("int2b4X=", int2b4X)
		// log.Println("int2b4Y=", int2b4Y)
		// log.Println("theta2b4=", theta2b4)
		// 交差角度が制限範囲内かどうか確認する
		divLine2b4 := math.Inf(1)
		if theta2b4 > 60 || theta2b4 < 120 {
			// 対向する辺の頂点が直行する直線に対して同じ側にある場合
			// 交点は対向する辺上にない．
			t1 := posline(chokuXYap[1][1], chokuXYap[0][1], chokuXYap[1][0],
				chokuXYap[0][0], taikoXYb4[0][0], taikoXYb4[0][1])
			t2 := posline(chokuXYap[1][1], chokuXYap[0][1], chokuXYap[1][0],
				chokuXYap[0][0], taikoXYb4[1][0], taikoXYb4[1][1])
			if t1*t2 < 0 {
				// 交点2b_5_6までの距離
				divLine2b4 = DistVerts(XY[num2][0], int2b4X, XY[num2][1], int2b4Y)
				// log.Println("XY[num2][0]=", XY[num2][0])
				// log.Println("int2b4X=", int2b4X)
				// log.Println("XY[num2][1]=", XY[num2][1])
				// log.Println("int2b4Y=", int2b4Y)
				// log.Println("divLine2b4", divLine2b4)
			} else {
				// TODO:関数から戻る
				// return
			}
		}
		// divLine2a1 := DistVerts(XY[num2][1], int2a1X, XY[num2][0], int2a1Y)
		// divLine2b4 := DistVerts(XY[num2][1], int2b4X, XY[num2][0], int2b4Y)

		// 分割点が共にR点に近接する場合，真ん中の四角形が非常に小さく細長くなる．
		// この場合は，2つのL点を結ぶ線を分割線とする．
		// R1点とD2b点の間の距離
		noR1 := (num1 + 1) % nodOct
		distR1D2b := DistVerts(XY[noR1][0], int2b4X, XY[noR1][1], int2b4Y)
		// log.Println("XY[noR1][0]=", XY[noR1][0])
		// log.Println("int2b4X=", int2b4X)
		// log.Println("XY[noR1][1]=", XY[noR1][1])
		// log.Println("int2b4Y=", int2b4Y)
		// R3点とD1a点の間の距離
		noR3 := (num1 + 3) % nodOct
		distR3D1a := DistVerts(XY[noR3][0], int1a1X, XY[noR3][1], int1a1Y)
		// log.Println("XY[noR3][0]=", XY[noR3][0])
		// log.Println("int1a1X=", int1a1X)
		// log.Println("XY[noR3][1]=", XY[noR3][1])
		// log.Println("int1a1Y=", int1a1Y)
		// R4点とD1b点の間の距離
		noR4 := (num2 + 1) % nodOct
		distR4D1b := DistVerts(XY[noR4][0], int1b4X, XY[noR4][1], int1b4Y)
		// log.Println("XY[noR4][0]=", XY[noR4][0])
		// log.Println("int1b4X=", int1b4X)
		// log.Println("XY[noR4][1]=", XY[noR4][1])
		// log.Println("int1b4Y=", int1b4Y)
		// R6点とD2a点の間の距離
		noR6 := (num2 + 3) % nodOct
		distR6D2a := DistVerts(XY[noR6][0], int2a1X, XY[noR6][1], int2a1Y)
		// log.Println("XY[noR6][0]=", XY[noR6][0])
		// log.Println("int2a1X=", int2a1X)
		// log.Println("XY[noR6][1]=", XY[noR6][1])
		// log.Println("int2a1Y=", int2a1Y)

		// 分割線1aと分割線2aが近接している場合
		// L点を結ぶ線を分割線とする．
		if distR3D1a < 1.0 && distR6D2a < 1.0 {
			rect1name = []string{"L1", "R1", "R2", "R3"}
			// log.Println("rect1name", rect1name)
			rect2name = []string{"L2", "R4", "R5", "R6"}
			// log.Println("rect2name", rect2name)
			// 辞書の中身に従ってリストの座標データで四角形を作る
			rect1List = MakeRectList(XY, order, rect1name)
			log.Println("rect1List=", rect1List)
			rect2List = MakeRectList(XY, order, rect2name)
			log.Println("rect2List=", rect2List)

			// 分割線1bと分割線2bが近接している場合
			// L点を結ぶ線を分割線とする．
		} else if distR1D2b < 1.0 && distR4D1b < 1.0 {
			rect1name = []string{"L2", "R1", "R2", "R3"}
			// log.Println("rect1name", rect1name)
			rect2name = []string{"L1", "R4", "R5", "R6"}
			// log.Println("rect2name", rect2name)
			// 辞書の中身に従ってリストの座標データで四角形を作る
			rect1List = MakeRectList(XY, order, rect1name)
			log.Println("rect1List=", rect1List)
			rect2List = MakeRectList(XY, order, rect2name)
			log.Println("rect2List=", rect2List)
		} else {
			// 分割線の長さの比較（並び替え）
			divideline := map[string]float64{
				"D1a1": divLine1a1,
				"D1b4": divLine1b4,
				"D2a1": divLine2a1,
				"D2b4": divLine2b4,
			}
			// 分割線の長さの構造体のリストを作る
			p := make(DivlineList, len(divideline))
			i := 0
			for k, v := range divideline {
				p[i] = Pair{k, v}
				i++
			}
			// 分割線の長さで並び替える
			sort.Sort(p)
			//p is sorted
			// 並び替えた結果を表示する
			for _, k := range p {
				log.Printf("%v\t%v\n", k.Key, k.Value)
			}
			log.Println(p[0].Key)

			// 分割点D1aから辺2-3への分割線が最短
			if p[0].Key == "D1a1" {
				// 分割線は1a
				// log.Println(p[0].Key)
				// 分割点はD1a点
				d1a := []float64{int1a1Y, int1a1X}
				// 座標値のリストにD1a点の座標値を追加する
				XY = append(XY, d1a)
				// log.Println(XY)
				// 頂点並びの辞書に分割点を追加する
				d1aNum := nodOct
				order["D1a"] = d1aNum
				// log.Println("line1a", order)

				// 四角形 L1-R1-R2-D1a
				rect1name = []string{"L1", "R1", "R2", "D1a"}
				// log.Println("rect1name", rect1name)
				// ６角形 L2-R4-R5-R6-D1a-R3
				hex1name = []string{"L2", "R4", "R5", "R6", "D1a", "R3"}
				// log.Println("hex1name", hex1name)
			}
			// 分割点D1bから辺5-6への分割線が最短
			if p[0].Key == "D1b4" {
				// 分割線は1b
				// log.Println(p[0].Key)
				// 分割点はD1b点
				d1b := []float64{int1b4Y, int1b4X}
				// 座標値のリストにD1b点の座標値を追加する
				XY = append(XY, d1b)
				// log.Println(XY)
				// 頂点並びの辞書に分割点を追加する
				d1bNum := nodOct
				order["D1b"] = d1bNum
				// log.Println("line1b", order)

				// 四角形 L1-D1b-R5-R6
				rect1name = []string{"L1", "D1b", "R5", "R6"}
				log.Println("rect1name", rect1name)
				// ６角形 L2-R4-D1b-R1-R2-R3
				hex1name = []string{"L2", "R4", "D1b", "R1", "R2", "R3"}
				log.Println("hex1name", hex1name)
			}
			// 分割点D2aから辺2-3への分割線が最短
			if p[0].Key == "D2a1" {
				// 分割線は2a
				// log.Println(p[0].Key)
				// 分割点はD2a点
				d2a := []float64{int2a1Y, int2a1X}
				// 座標値のリストにD1a点の座標値を追加する
				XY = append(XY, d2a)
				// log.Println(XY)
				// 頂点並びの辞書に分割点を追加する
				d2aNum := nodOct
				order["D2a"] = d2aNum
				// log.Println("line2a", order)

				// 四角形 L2-R4-R5-D2a
				rect1name = []string{"L2", "R4", "R5", "D2a"}
				// log.Println("rect1name", rect1name)
				// ６角形 L1-R1-R2-R3-D2a-R6
				hex1name = []string{"L1", "R1", "R2", "R3", "D2a", "R6"}
				// log.Println("hex1name", hex1name)
			}
			// 分割点D2bから辺5-6への分割線が最短
			if p[0].Key == "D2b4" {
				// 分割線は2b
				// log.Println(p[0].Key)
				// 分割点はD2b点
				d2b := []float64{int2b4Y, int2b4X}
				// 座標値のリストにD1b点の座標値を追加する
				XY = append(XY, d2b)
				// log.Println(XY)
				// 頂点並びの辞書に分割点を追加する
				d2bNum := nodOct
				order["D2b"] = d2bNum
				// log.Println("line2b", order)

				// 四角形 L2-D2b-R2-R3
				rect1name = []string{"L2", "D2b", "R2", "R3"}
				// log.Println("rect1name", rect1name)
				// ６角形 L1-R1-D2b-R4-R5-R6
				hex1name = []string{"L1", "R1", "D2b", "R4", "R5", "R6"}
				// log.Println("hex1name", hex1name)
			}
			// 辞書の中身に従ってリストの座標データで四角形を作る
			rect1List = MakeRectList(XY, order, rect1name)
			log.Println("rect1List=", rect1List)
			// 辞書の中身に従ってリストの座標データで６角形を作る
			// var rect1H [][]float64
			for _, v := range hex1name {
				// log.Println(i, v)
				n := order[v]
				hex1L = append(hex1L, XY[n])
			}
			// log.Println("hex1L=", hex1L)
			if hex1L != nil {
				// 頂点データ数の確認
				nod := len(hex1L)
				if nod != 6 {
					// TODO:
				}
				// 追加したD点のために外積の計算をやり直す必要がある
				extL, _, _ := TriVert(nod, hex1L)
				// L点，R点の辞書を作り直す
				_, _, orderN, _, _ := Lexicogra(nod, hex1L, extL)
				// log.Println("orderN=", orderN)
				// 6角形の四角形分割プログラムに渡す
				_, rect2List, rect3List = HexaDiv(hex1L, orderN)
				log.Println("rectO2L", rect2List)
				log.Println("rectO3L", rect3List)
			}
		}
	} else {
		octStype := "Others：陸屋根"
		log.Println(octStype)
		// 関数から戻る
		return
	}

	cord = XY
	// log.Println("cord=", cord)
	return cord, rect1List, rect2List, rect3List, hex1L
}
