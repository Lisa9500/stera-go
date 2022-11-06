package pkg

import (
	"log"
	"reflect"
)

// OctaDiv は８角形を３つの４角形に分割する
func OctaDiv(XY [][]float64, order map[string]int, pttrn []string,
	indx []int) (cord [][]float64, rect1List [][]float64, rect2List [][]float64,
	rect3List [][]float64, hex1List [][]float64) {
	// 頂点データ数の確認
	nodOct := len(XY)
	if nodOct != 8 {
		return
	}
	var chkptn []string
	var octT string

	chkptn = []string{"L", "R", "L", "R", "R", "R", "R", "R"}
	octT = "歯型1"
	if reflect.DeepEqual(pttrn, chkptn) == true {
		log.Println(octT)
		cord, rect1List, rect2List, rect3List = MakeOct1(XY, order)
		// log.Println(cord)
		// log.Println(rect1List)
		// log.Println(rect2List)
		// log.Println(rect3List)
	}

	chkptn = []string{"L", "R", "R", "R", "R", "R", "L", "R"}
	octT = "歯型2"
	if reflect.DeepEqual(pttrn, chkptn) == true {
		log.Println(octT)
		cord, rect1List, rect2List, rect3List = MakeOct2(XY, order)
		// log.Println(cord)
		// log.Println(rect1List)
		// log.Println(rect2List)
		// log.Println(rect3List)
	}

	chkptn = []string{"L", "L", "R", "R", "R", "R", "R", "R"}
	octT = "凹型1"
	if reflect.DeepEqual(pttrn, chkptn) == true {
		log.Println(octT)
		cord, rect1List, rect2List, rect3List = MakeOct3(XY, order)
		// log.Println(cord)
		// log.Println(rect1List)
		// log.Println(rect2List)
		// log.Println(rect3List)
	}

	chkptn = []string{"L", "R", "R", "R", "R", "R", "R", "L"}
	octT = "凹型2"
	if reflect.DeepEqual(pttrn, chkptn) == true {
		log.Println(octT)
		cord, rect1List, rect2List, rect3List = MakeOct4(XY, order)
		// log.Println(cord)
		// log.Println(rect1List)
		// log.Println(rect2List)
		// log.Println(rect3List)
	}

	chkptn = []string{"L", "R", "R", "L", "R", "R", "R", "R"}
	octT = "凸型1"
	if reflect.DeepEqual(pttrn, chkptn) == true {
		log.Println(octT)
		cord, rect1List, rect2List, rect3List = MakeOct5(XY, order)
		// log.Println(cord)
		// log.Println(rect1List)
		// log.Println(rect2List)
		// log.Println(rect3List)
	}

	chkptn = []string{"L", "R", "R", "R", "R", "L", "R", "R"}
	octT = "凸型2"
	if reflect.DeepEqual(pttrn, chkptn) == true {
		log.Println(octT)
		cord, rect1List, rect2List, rect3List = MakeOct6(XY, order)
		// log.Println(cord)
		// log.Println(rect1List)
		// log.Println(rect2List)
		// log.Println(rect3List)
	}

	chkptn = []string{"L", "R", "R", "R", "L", "R", "R", "R"}
	octT = "Ｓ型"
	if reflect.DeepEqual(pttrn, chkptn) == true {
		log.Println(octT)
		cord, rect1List, rect2List, rect3List, hex1List = MakeOct7(XY, order)
		// log.Println(cord)
		// log.Println(rect1List)
		// log.Println(rect2List)
		// log.Println(rect3List)
		// log.Println(hex1List)
	}

	if hex1List != nil {
		// 追加したD点のために外積の計算をやり直す必要がある
		extL, _, _ := TriVert(6, hex1List)
		// L点，R点の辞書を作り直す
		_, _, orderN, _, _ := Lexicogra(6, hex1List, extL)
		// log.Println("orderN=", orderN)
		// 6角形の四角形分割プログラムに渡す
		_, rect2List, rect3List = HexaDiv(hex1List, orderN)
		// log.Println("rect1List", rect1List)
		// log.Println("rect2List", rect2List)
	}

	return cord, rect1List, rect2List, rect3List, hex1List
}
