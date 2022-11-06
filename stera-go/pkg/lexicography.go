package pkg

import (
	"strconv"
)

// Lexicogra はL点とR点をリストおよび辞書に振り分ける
func Lexicogra(nod int, XY [][]float64, ext []float64) (lList [][]float64,
	rList [][]float64, order map[string]int, lrPttrn []string, lrIndx []int) {
	// log.Println("nod=", nod)
	// L点の番号
	var lNum = 1
	// R点の番号
	var rNum = 1
	// R点の予備の座標リストを用意する
	var rsusL [][]float64
	order = map[string]int{}

	for i := 0; i < nod; i++ {
		// L点の場合の処理
		if ext[i] > 0 {
			// L点が見つかった順に処理している
			// log.Println("i=", i)
			// log.Println("ext=", ext[i])
			lList = append(lList, XY[i][0:2])
			// log.Println("lList=", lList)
			order["L"+strconv.Itoa(lNum)] = i
			// log.Println("order=", order)
			// lrPttrnは必ずLから始まる
			lrPttrn = append(lrPttrn, "L")
			lrIndx = append(lrIndx, i)
			lNum++
		} else {
			// R点の場合の処理
			// L点以前のR点は仮置きする
			if len(lList) == 0 {
				rsusL = append(rsusL, XY[i][0:2])
				// log.Println("rsusL=", rsusL)
			} else {
				rList = append(rList, XY[i][0:2])
				// log.Println("rList=", rList)
				order["R"+strconv.Itoa(rNum)] = i
				// log.Println("order=", order)
				lrPttrn = append(lrPttrn, "R")
				lrIndx = append(lrIndx, i)
				rNum++
			}
		}
	}
	// 仮置きしたR点を後置する
	rList = append(rList, rsusL...)
	// log.Println("rList=", rList)
	sLen := len(rsusL)
	// log.Println("sLen=", sLen)
	for j := 0; j < sLen; j++ {
		order["R"+strconv.Itoa(rNum)] = j
		lrPttrn = append(lrPttrn, "R")
		lrIndx = append(lrIndx, j)
		rNum++
	}
	// log.Println("order=", order)
	// log.Println("lrPttrn=", lrPttrn)
	// log.Println("lrIndx=", lrIndx)

	return lList, rList, order, lrPttrn, lrIndx
}
