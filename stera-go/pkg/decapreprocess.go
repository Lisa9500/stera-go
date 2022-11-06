package pkg

import (
	"log"
	"math"
	"sort"
)

// DecaPreprocess は10角形を１つの四角形と１つの8角形に分割する
func DecaPreprocess(num int, XY [][]float64, nod int,
	order map[string]int) (cord [][]float64, ordering map[string]int,
	keyList []string, areaA, areaB float64) {
	log.Println("num", num)
	log.Println("nod", nod)
	// 直交する辺aは．L点と1つ前の点で結ばれる線分
	chokuPairA, distA := VertsA(num, nod, XY)
	log.Println("distA", distA)
	// もう一方の直交する辺bは．L点と次の点で結ばれる線分
	chokuPairB, distB := VertsB(num, nod, XY)
	log.Println("distB", distB)
	// 直交する辺aに対抗する辺を求める
	// taikoPairA := make([][]float64, 2)
	taikoPairA := OpposeA(num, XY, nod)
	log.Println("taikoPairA", taikoPairA)
	// 直交する辺bに対抗する辺を求める
	// taikoPairB := make([][]float64, 2)
	taikoPairB := OpposeB(num, XY, nod)
	log.Println("taikoPairB", taikoPairB)

	// 直交する辺aに対抗する辺との交点を求める
	intAX, intAY, thetaA := OrthoAngle(chokuPairA, taikoPairA)
	// 交差角度が制限範囲内かどうか確認する
	if thetaA < 60 || thetaA > 120 {
		// TODO:四角形分割は行わない
		// return
	}
	// 交点が対向する辺上にあるかチェックする
	// 対向する辺の頂点が直行する直線に対して同じ側にある場合
	// 交点は対向する辺上にない．
	A1 := (chokuPairA[1][1]-chokuPairA[0][1])*(taikoPairA[0][0]-chokuPairA[0][0]) -
		(chokuPairA[1][0]-chokuPairA[0][0])*(taikoPairA[0][1]-chokuPairA[0][1])
	A2 := (chokuPairA[1][1]-chokuPairA[0][1])*(taikoPairA[1][0]-chokuPairA[0][0]) -
		(chokuPairA[1][0]-chokuPairA[0][0])*(taikoPairA[1][1]-chokuPairA[0][1])
	if A1*A2 < 0 {
		// 四角形aの面積areaAを求める
		var Sa1 float64
		num2 := (num + 2) % nod
		Sa1 = TriArea(XY[num][0], XY[num][1], intAX, intAY, XY[num2][0], XY[num2][1])
		log.Println("Sa1", Sa1)
		var Sa2 float64
		num1 := (num + 1) % nod
		Sa2 = TriArea(XY[num][0], XY[num][1], XY[num1][0], XY[num1][1], XY[num2][0], XY[num2][1])
		log.Println("Sa2", Sa2)
		areaA = Sa1 + Sa2
	} else {
		areaA = math.Inf(1)
	}
	log.Println("areaA", areaA)

	// 直交する辺bに対抗する辺との交点を求める
	intBX, intBY, thetaB := OrthoAngle(chokuPairB, taikoPairB)
	log.Println("intBX", intBX)
	log.Println("intBY", intBY)
	log.Println("thetaB", thetaB)
	// 交差角度が制限範囲内かどうか確認する
	if thetaB < 60 || thetaB > 120 {
		// TODO:四角形分割は行わない
		// return
	}
	// 交点が対向する辺上にあるかチェックする
	// 対向する辺の頂点が直行する直線に対して同じ側にある場合
	// 交点は対向する辺上にない．
	B1 := (chokuPairB[1][1]-chokuPairB[0][1])*(taikoPairB[0][0]-chokuPairB[0][0]) -
		(chokuPairB[1][0]-chokuPairB[0][0])*(taikoPairB[0][1]-chokuPairB[0][1])
	B2 := (chokuPairB[1][1]-chokuPairB[0][1])*(taikoPairB[1][0]-chokuPairB[0][0]) -
		(chokuPairB[1][0]-chokuPairB[0][0])*(taikoPairB[1][1]-chokuPairB[0][1])
	if B1*B2 < 0 {
		// 四角形aの面積areaBを求める
		var Sb1 float64
		numn2 := (num - 2 + nod) % nod
		log.Println("XY[num][1]", XY[num][1])
		log.Println("XY[num][0]", XY[num][0])
		log.Println("XY[num-2][1]", XY[numn2][1])
		log.Println("XY[num-2][0]", XY[numn2][0])
		Sb1 = TriArea(XY[num][0], XY[num][1], intBX, intBY, XY[numn2][0], XY[numn2][1])
		log.Println("Sb1", Sb1)
		var Sb2 float64
		numn1 := (num - 1 + nod) % nod
		Sb2 = TriArea(XY[num][0], XY[num][1], XY[numn1][0], XY[numn1][1], XY[numn2][0], XY[numn2][1])
		log.Println("Sb2", Sb2)
		areaB = Sb1 + Sb2
	} else {
		areaB = math.Inf(1)
	}
	log.Println("areaB", areaB)

	// 分割点はD0点
	var d0 []float64
	if areaA < areaB {
		d0 = []float64{intAY, intAX}
		log.Println("d0=", d0)
	} else if areaA > areaB {
		d0 = []float64{intBY, intBX}
		log.Println("d0=", d0)
	}
	// 座標値のリストにD0点の座標値を追加する
	cord = append(XY, d0)
	log.Println("cord", cord)
	// 頂点並びの辞書に分割点を追加する
	d0Num := nod
	order["D0"] = d0Num
	ordering = order
	log.Println("order", ordering)
	// 辞書からキーのリストを作る
	for key := range ordering {
		keyList = append(keyList, key)
	}
	// 頂点並びの順序で値（Ｌ･Ｒ点）のリストを作る
	sort.SliceStable(keyList, func(i int, j int) bool {
		return ordering[keyList[i]] < ordering[keyList[j]]
	})
	log.Println("keyList", keyList)

	return cord, ordering, keyList, areaA, areaB
}
