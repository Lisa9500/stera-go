package pkg

import (
	"log"
	"strings"
)

func mkrectA(cordz [][]float64, order2 map[string]int, keyList []string, nodDec,
	num0, d0Num int) (rect1name []string, rect4L [][]float64) {
	// 四角形を作る
	// L3//L1点と次のR点とその次のR点とD点: num0, num0+1, num0+2, d0Num
	rect1name = append(rect1name, keyList[num0])
	num1 := (num0 + 1) % nodDec
	rect1name = append(rect1name, keyList[num1])
	num2 := (num0 + 2) % nodDec
	rect1name = append(rect1name, keyList[num2])
	rect1name = append(rect1name, keyList[d0Num])
	for _, name := range rect1name {
		n := order2[name]
		rect4L = append(rect4L, cordz[n])
	}
	return rect1name, rect4L
}

func mkrectC(cordz [][]float64, order2 map[string]int, keyList []string, nodDec,
	num0, d0Num int) (rect1name []string, rect4L [][]float64) {
	// 四角形を作る
	// L3点とD点と前の々のR点と前のR点: num0, d0Num, num0+8, num0+9
	rect1name = append(rect1name, keyList[num0])
	rect1name = append(rect1name, keyList[d0Num])
	num8 := (num0 + 8) % nodDec
	rect1name = append(rect1name, keyList[num8])
	num9 := (num0 + 9) % nodDec
	rect1name = append(rect1name, keyList[num9])
	for _, name := range rect1name {
		n := order2[name]
		rect4L = append(rect4L, cordz[n])
	}
	return rect1name, rect4L
}

func mkoctaA(cordz [][]float64, order2 map[string]int, keyList []string, nodDec,
	num0, d0Num int) (octa1name []string, octa1L [][]float64) {
	// 8角形を作る
	// d0Num, num0+3, num0+4, num0+5, num0+6, num0+7, num0+8, num0+9
	octa1name = append(octa1name, keyList[d0Num])
	for i := 3; i < 10; i++ {
		numi := (num0 + i) % nodDec
		octa1name = append(octa1name, keyList[numi])
	}
	for _, name := range octa1name {
		n := order2[name]
		octa1L = append(octa1L, cordz[n])
	}
	return octa1name, octa1L
}

func mkoctaC(cordz [][]float64, order2 map[string]int, keyList []string, nodDec,
	num0, d0Num int) (octa1name []string, octa1L [][]float64) {
	// 8角形を作る
	// keyListはインデックス0から順にＬ･Ｒ点が並んでいる．
	// L3点，R1点，R2点を除外する
	// num0とnum0-1，num0-2を削除する
	// d0Num, num0+1, num0+2, num0+3, num0+4, num0+5, num0+6, num0+7
	octa1name = append(octa1name, keyList[d0Num])
	for i := 1; i < 8; i++ {
		numi := (num0 + i) % nodDec
		octa1name = append(octa1name, keyList[numi])
	}
	for _, name := range octa1name {
		n := order2[name]
		octa1L = append(octa1L, cordz[n])
	}
	return octa1name, octa1L
}

// DecaDiv は10角形を１つの四角形と１つの8角形に分割する
func DecaDiv(cord2 [][]float64, order map[string]int, lrPtn []string,
	lrIdx []int) (rect1L [][]float64, rect2L [][]float64, rect3L [][]float64,
	rect4L [][]float64) {
	var hex1L [][]float64
	var octa1name []string
	var octa1L [][]float64

	// log.Println(lrPtn)
	// log.Println(lrIdx)

	// 頂点データ数の確認
	nodDec := len(cord2)
	if nodDec != 10 {
		// TODO:関数から戻る
		return
	}
	d0Num := nodDec
	var num0 int

	// LR並びの確認　L点から始まっていなければエラー
	if lrPtn[0] != "L" {
		// TODO:関数から戻る
		return
	}
	// 検索用にLR並びから半角スペースを除く
	lrjoin := strings.Join(lrPtn, "")
	// log.Println("lrjoin=", lrjoin)

	switch lrjoin {
	// １－①
	case "LLRLRRRRRR", "LRLRRRRRRL", "LRRRRRRLLR":
		if lrPtn[1] == "L" && lrPtn[3] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[2] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[2]
		} else if lrPtn[7] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaA, areaB := DecaPreprocess(num0, cord2, nodDec,
			order)
		// 小さい耳となる方の四角形を分割し，残りで8角形を作る
		if areaA < areaB {
			// 四角形を作る
			rect1name, _ := mkrectA(cordz, order2, keyList, nodDec, num0, d0Num)
			// // rect1name = ['L3', 'R2', 'R3', 'D1']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = mkoctaA(cordz, order2, keyList, nodDec, num0, d0Num)
			// L3点，R2点，R3点を除外する
			log.Println("octa1name", octa1name)
			// log.Println(octa1L)
		} else if areaA > areaB {
			// areaAはareaBより必ず小さい
			// TODO:
		}

	// １－②
	case "LLRRLRRRRR", "LRRLRRRRRL", "LRRRRRLLRR":
		// L3点のインデックスがnum0
		if lrPtn[1] == "L" && lrPtn[4] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[3] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[6] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaA, areaB := DecaPreprocess(num0, cord2, nodDec,
			order)
		// 小さい耳となる方の四角形を分割し，残りで8角形を作る
		if areaA < areaB {
			// 四角形を作る
			rect1name, _ := mkrectA(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'R3', 'R4', 'D1']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = mkoctaA(cordz, order2, keyList, nodDec, num0, d0Num)
			// L3点，R3点，R4点を除外する
			log.Println("octa1name", octa1name)
			log.Println(octa1L)
		} else if areaA > areaB {
			// 四角形を作る
			rect1name, _ := mkrectC(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'D0', 'R1', 'R2']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = mkoctaC(cordz, order2, keyList, nodDec, num0, d0Num)
			// L3点，R1点，R2点を除外する
			log.Println("octa1name", octa1name)
			log.Println(octa1L)
		}

	// １－③
	case "LLRRRLRRRR", "LRRRLRRRRL", "LRRRRLLRRR":
		if lrPtn[1] == "L" && lrPtn[5] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[4] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[5] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaA, areaB := DecaPreprocess(num0, cord2, nodDec,
			order)
		// 小さい耳となる方の四角形を分割し，残りで8角形を作る
		if areaA < areaB {
			// 四角形を作る
			rect1name, _ := mkrectA(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'R4', 'R5', 'D2']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = mkoctaA(cordz, order2, keyList, nodDec, num0, d0Num)
			// L3点，R4点，R5点を除外する
			log.Println("octa1name", octa1name)
			// log.Println(octa1L)
		} else if areaA > areaB {
			// 四角形を作る
			rect1name, _ := mkrectC(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'D2', 'R2', 'R3']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = mkoctaC(cordz, order2, keyList, nodDec, num0, d0Num)
			// L3点，R2点，R3点を除外する
			log.Println("octa1name", octa1name)
			// log.Println(octa1L)
		}

	// １－④
	case "LLRRRRLRRR", "LRRRRLRRRL", "LRRRLLRRRR":
		if lrPtn[1] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[5] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[4] == "L" && lrPtn[5] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("num0", num0)
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaA, areaB := DecaPreprocess(num0, cord2, nodDec,
			order)
		// 小さい耳となる方の四角形を分割し，残りで8角形を作る
		if areaA < areaB {
			// 四角形を作る
			rect1name, _ := mkrectA(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'R5', 'R6', 'D1']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = mkoctaA(cordz, order2, keyList, nodDec, num0, d0Num)
			// L3点，R5点，R6点を除外する
			log.Println("octa1name", octa1name)
			// log.Println(octa1L)
		} else if areaA > areaB {
			// 四角形を作る
			rect1name, _ := mkrectC(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'D0', 'R3', 'R4']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = mkoctaC(cordz, order2, keyList, nodDec, num0, d0Num)
			// L3点，R3点，R4点を除外する
			log.Println("octa1name", octa1name)
			// log.Println(octa1L)
		}

	// １－⑤
	case "LLRRRRRLRR", "LRRRRRLRRL", "LRRLLRRRRR":
		if lrPtn[1] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[6] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[3] == "L" && lrPtn[4] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaA, areaB := DecaPreprocess(num0, cord2, nodDec,
			order)
		// 小さい耳となる方の四角形を分割し，残りで8角形を作る
		if areaA < areaB {
			// 四角形を作る
			rect1name, _ := mkrectA(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'R6', 'R7', 'D1']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = mkoctaA(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			// log.Println(octa1L)
		} else if areaA > areaB {
			// 四角形を作る
			rect1name, _ := mkrectC(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'D0', 'R4', 'R5']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = mkoctaC(cordz, order2, keyList, nodDec, num0, d0Num)
			// L3点，R4点，R5点を除外する
			log.Println("octa1name", octa1name)
			// log.Println(octa1L)
		}

	// １－⑥
	case "LLRRRRRRLR", "LRRRRRRLRL", "LRLLRRRRRR":
		if lrPtn[1] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[8]
		} else if lrPtn[7] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[2] == "L" && lrPtn[3] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaA, areaB := DecaPreprocess(num0, cord2, nodDec,
			order)
		// 小さい耳となる方の四角形を分割し，残りで8角形を作る
		if areaA < areaB {
			// areaBはareaAより必ず小さい
			// TODO:
		} else if areaA > areaB {
			// 四角形を作る
			rect1name, _ := mkrectC(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'D2', 'R5', 'R6']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = mkoctaC(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			// log.Println(octa1L)
		}

	// ２－①
	case "LRLRLRRRRR", "LRLRRRRRLR", "LRRRRRLRLR":
		if lrPtn[2] == "L" && lrPtn[4] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[2] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[2]
		} else if lrPtn[6] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaA, areaB := DecaPreprocess(num0, cord2, nodDec,
			order)
		// 小さい耳となる方の四角形を分割し，残りで8角形を作る
		if areaA < areaB {
			// 四角形を作る
			// L3//L1点と次のR点とその次のR点とD点: num0, num0+1, num0+2, d0Num
			rect1name, _ := mkrectA(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'R3', 'R4', 'D1']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = mkoctaA(cordz, order2, keyList, nodDec, num0, d0Num)
			// L3点，R3点，R4点を除外する
			// num0とnum0+1，num0+2を削除する
			// d0Num, num0+3, num0+4, num0+5, num0+6, num0+7, num0+8, num0+9
			log.Println("octa1name", octa1name)
			// log.Println(octa1L)
		} else if areaA > areaB {
			// areaAはareaBより必ず小さい
			// TODO:
		}

	// ２－②
	case "LRLRRLRRRR", "LRRLRRRRLR", "LRRRRLRLRR":
		if lrPtn[2] == "L" && lrPtn[5] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[3] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[5] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaA, areaB := DecaPreprocess(num0, cord2, nodDec,
			order)
		// 小さい耳となる方の四角形を分割し，残りで8角形を作る
		if areaA < areaB {
			// 四角形を作る
			// L3//L1点と次のR点とその次のR点とD点: num0, num0+1, num0+2, d0Num
			rect1name, _ := mkrectA(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'R4', 'R5', 'D1']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			// d0Num, num0+3, num0+4, num0+5, num0+6, num0+7, num0+8, num0+9
			octa1name, octa1L = mkoctaA(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			// log.Println(octa1L)
		} else if areaA > areaB {
			// 四角形を作る
			// L3点とD点と前の々のR点と前のR点: num0, d0Num, num0+8, num0+9
			rect1name, _ := mkrectC(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'D0', 'R2', 'R3']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			// L3点，R2点，R3点を除外する
			// num0とnum0-1，num0-2を削除する
			// d0Num, num0+1, num0+2, num0+3, num0+4, num0+5, num0+6, num0+7
			octa1name, octa1L = mkoctaC(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			// log.Println(octa1L)
		}

	// ２－③
	case "LRLRRRLRRR", "LRRRLRRRLR", "LRRRLRLRRR":
		if lrPtn[2] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[4] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[4] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaA, areaB := DecaPreprocess(num0, cord2, nodDec,
			order)
		// 小さい耳となる方の四角形を分割し，残りで8角形を作る
		if areaA < areaB {
			// 四角形を作る
			// L3//L1点と次のR点とその次のR点とD点: num0, num0+1, num0+2, d0Num
			rect1name, _ := mkrectA(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'R5', 'R6', 'D1']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			// d0Num, num0+3, num0+4, num0+5, num0+6, num0+7, num0+8, num0+9
			octa1name, octa1L = mkoctaA(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			// log.Println(octa1L)
		} else if areaA > areaB {
			// 四角形を作る
			// L3点とD点と前の々のR点と前のR点: num0, d0Num, num0+8, num0+9
			rect1name, _ := mkrectC(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'D0', 'R3', 'R4']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			// L3点，R3点，R4点を除外する
			// num0とnum0-1，num0-2を削除する
			// d0Num, num0+1, num0+2, num0+3, num0+4, num0+5, num0+6, num0+7
			octa1name, octa1L = mkoctaC(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			// log.Println(octa1L)
		}

	// ２－④
	case "LRLRRRRLRR", "LRRRRLRRLR", "LRRLRLRRRR":
		if lrPtn[2] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[5] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[3] == "L" && lrPtn[5] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaA, areaB := DecaPreprocess(num0, cord2, nodDec,
			order)
		// 小さい耳となる方の四角形を分割し，残りで8角形を作る
		if areaA < areaB {
			// 四角形を作る
			// L3//L1点と次のR点とその次のR点とD点: num0, num0+1, num0+2, d0Num
			rect1name, _ := mkrectA(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'R6', 'R7', 'D1']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			// d0Num, num0+3, num0+4, num0+5, num0+6, num0+7, num0+8, num0+9
			octa1name, octa1L = mkoctaA(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			// log.Println(octa1L)
		} else if areaA > areaB {
			// 四角形を作る
			// L3点とD点と前の々のR点と前のR点: num0, d0Num, num0+8, num0+9
			rect1name, _ := mkrectC(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'D0', 'R2', 'R3']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			// L3点，R2点，R3点を除外する
			// num0とnum0-1，num0-2を削除する
			// d0Num, num0+1, num0+2, num0+3, num0+4, num0+5, num0+6, num0+7
			octa1name, octa1L = mkoctaC(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			// log.Println(octa1L)
		}

	// ３－②
	case "LRRLRRLRRR", "LRRLRRRLRR", "LRRRLRRLRR":
		if lrPtn[3] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[3] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[4] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaA, areaB := DecaPreprocess(num0, cord2, nodDec,
			order)
		// 小さい耳となる方の四角形を分割し，残りで8角形を作る
		if areaA < areaB {
			// 四角形を作る
			// L3//L1点と次のR点とその次のR点とD点: num0, num0+1, num0+2, d0Num
			rect1name, _ := mkrectA(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'R5', 'R6', 'D1']
			// rect1name = ['L2', 'R3', 'R4', 'D1']
			// rect1name = ['L1', 'R1', 'R2', 'D1']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			// d0Num, num0+3, num0+4, num0+5, num0+6, num0+7, num0+8, num0+9
			octa1name, octa1L = mkoctaA(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			// log.Println(octa1L)
		} else if areaA > areaB {
			// 四角形を作る
			// L3点とD点と前の々のR点と前のR点: num0, d0Num, num0+8, num0+9
			rect1name, _ := mkrectC(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'D0', 'R3', 'R4']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			// L3点，R2点，R3点を除外する
			// num0とnum0-1，num0-2を削除する
			// d0Num, num0+1, num0+2, num0+3, num0+4, num0+5, num0+6, num0+7
			octa1name, octa1L = mkoctaC(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			// log.Println(octa1L)
		}

		// ７－①
	case "LLLRRRRRRR", "LLRRRRRRRL", "LRRRRRRRLL":
		if lrPtn[1] == "L" && lrPtn[2] == "L" {
			num0 = lrIdx[2]
		} else if lrPtn[1] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[1]
		} else if lrPtn[8] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaA, areaB := DecaPreprocess(num0, cord2, nodDec,
			order)
		// 小さい耳となる方の四角形を分割し，残りで8角形を作る
		if areaA < areaB {
			// 四角形を作る
			// L3//L1点と次のR点とその次のR点とD点: num0, num0+1, num0+2, d0Num
			rect1name, _ := mkrectA(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'R1', 'R2', 'D1']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			// d0Num, num0+3, num0+4, num0+5, num0+6, num0+7, num0+8, num0+9
			octa1name, octa1L = mkoctaA(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			// log.Println("octa1L=", octa1L)
		} else if areaA > areaB {
			// areaAはareaBより必ず小さい
			// TODO:
		}
	default:
		// TODO:
	}

	if octa1L != nil {
		log.Println("octa1L=", octa1L)
		// 頂点データ数の確認
		nod := len(octa1L)
		if nod != 8 {
			// TODO:
		}
		// 追加したD点のために外積の計算をやり直す必要がある
		extL, _, _ := TriVert(nod, octa1L)
		// L点，R点の辞書を作り直す
		_, _, orderN, lrPtn, lrIdx := Lexicogra(nod, octa1L, extL)
		log.Println("orderN=", orderN)
		// 8角形の四角形分割プログラムに渡す
		_, rect1L, rect2L, rect3L, hex1L = OctaDiv(octa1L, orderN, lrPtn, lrIdx)
		log.Println("rectO1L=", rect1L)
		log.Println("rectO2L=", rect2L)
		log.Println("rectO3L=", rect3L)
		log.Println("hexO1L=", hex1L)
	}

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
		_, rect2L, rect3L = HexaDiv(hex1L, orderN)
		log.Println("rectH2L", rect2L)
		log.Println("rectH3L", rect3L)
	}

	return rect1L, rect2L, rect3L, rect4L
}
