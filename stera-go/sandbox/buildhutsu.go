package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"stera/pkg"
	"strings"
	"unsafe"
)

// SlopeRoof は傾斜屋根プログラム用の構造体の定義
type SlopeRoof struct {
	ID   string
	Elv  float64
	List [][]float64
}

// Polygon は多角柱プログラム用の構造体の定義
type Polygon struct {
	ID   string
	Elv  float64
	List [][]float64
}

// 四角形データ（構造体）をリスト（スライス）に追加する
func addrect(data SlopeRoof, list []SlopeRoof) (rectList []SlopeRoof) {
	rectList = append(list, data)
	// log.Println("rect=", data)
	return rectList
}

// 多角形データ（構造体）をリスト（スライス）に追加する
func addpoly(data Polygon, list []Polygon) (polyList []Polygon) {
	polyList = append(list, data)
	// log.Println("polygon=", data)
	return polyList
}

func main() {
	// ログファイルを新規作成，追記，書き込み専用，パーミションは読むだけ
	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// ログの出力先を変更
	log.SetOutput(file)

	fp, er := os.Open("C:/data/hutsu_list.txt")
	if er != nil {
		log.Fatal(er)
	}
	defer fp.Close()
	log.Println("ファイルポインタ", fp)

	// 構造体のフィールド
	var id string
	var elv float64
	var cordnts [][][]float64

	// 四角形データ（構造体）のスライスを作成する
	var rectList = make([]SlopeRoof, 0)

	// 多角形データ（構造体）のスライスを作成する
	var polyList = make([]Polygon, 0)

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		// ここで一行ずつ処理
		jStr := scanner.Text()
		// 右端の「,」を削除，「,」がない行末でもエラーにならない
		jStr = strings.TrimRight(jStr, ",")
		// GeoJson構造体の変数stcDataを宣言
		// var stcData GeoJson
		// var id string
		// var elv float64
		// var cordnts [][][]float64
		id, elv, cordnts = pkg.ParseJSON(jStr)

		// 配列の長さを取得する
		l := (len(cordnts[0]))
		// 頂点データ数をチェックする，２以下の場合は処理を中止して次の行に進む
		if l <= 2 {
			// 頂点数が２以下なのでモデリングしない
			// 当該処理を飛ばして次の処理に移る
			continue
		}
		// 配列の要素を取得する
		// log.Println("配列の要素", cordnts[0][2][0])
		// 平面直角座標系のX軸は真北に向かう値が正，Y軸は真東に向かう値が正
		cords := cordnts[0][0:2][0:l]
		// log.Println("2次元配列", cords)

		// 傾斜屋根でモデリングする
		var slR bool
		slR = true
		// 閉じた図形かどうかを判断し頂点数を求める
		chkCls := pkg.ChkClose(l, cords)
		var numV int
		if chkCls == true {
			numV = l - 1
		} else {
			numV = l
		}
		// 閉じていない図形としての頂点数
		// log.Println("閉じていない図形の頂点数", numV)

		// 外積を計算して右向き・左向きを求める
		// 内積を計算して内角の角度を求める
		// 時計回りかどうか判断する
		extLst, degLst, rotate := pkg.TriVert(numV, cords)

		// 反時計周りは時計回りに修正する
		var nCords [][]float64
		var nExt, nDeg []float64
		if rotate == "ccw" {
			nCords, nExt, nDeg = pkg.CcwRev(numV, cords, extLst, degLst)
		} else if rotate == "cw" {
			nCords = cords[:numV]
			nExt = extLst[:numV]
			nDeg = degLst[:numV]
		}

		// 内角が約180度の頂点を削除する
		// 対象とする角度の変更はflattenvert.goで行う
		nodz, cordz, extLz, degLz := pkg.FlatVert(numV, nCords, nExt, nDeg)

		// 近接している頂点を削除する
		// 頂点間の距離の設定はnododel.goで行う
		nod2, cord2, extL2, _ := pkg.NodDel(nodz, cordz, extLz, degLz)

		// 内角条件を設定し，満たさない内角がある場合は，四角形分割を行わない
		for d := range degLz {
			if degLz[d] < 45.0 || degLz[d] > 135.0 {
				// TODO:三角メッシュの分割プログラムに渡す
				slR = false
				// poly := Polygon{ID: id, Elv: elv, List: cord2}
				// polyList = addpoly(poly, polyList)
			}
		}

		// log.Println("頂点数", nod2)
		log.Println("頂点座標", cord2)
		// log.Println("外積", extL2)
		// log.Println("内角", degL2)

		// 四角形分割のために多角形から凹頂点のL点を抽出する
		// Ｎ角形　内角数：N=2x,x=N/2，凹角数：L=x-2=N/2-2
		lcnt := nod2/2 - 2
		// L点の座標リストを作成する
		// var lL [][]float64
		// R点の座標リストを作成する
		// var rL [][]float64
		// 頂点並びのL点・R点の辞書を作成する
		// var order = map[string]int{}
		// L点とR点をリストおよび辞書に振り分ける
		lL, _, order, lrPtn, lrIdx := pkg.Lexicogra(nod2, cord2, extL2)
		// log.Println("lL", lL)
		// log.Println("rL", rL)
		log.Println("order", order)
		log.Println("lrPtn", lrPtn)
		log.Println("lrIdx", lrIdx)

		// L点と凹角数が一致しない場合は傾斜屋根でモデリングしない
		if lcnt != len(lL) {
			slR = false
		}

		// "頂点座標"cord2と"LR並び"orderを使って多角形の分割を行う
		// 四角形分割ができない場合，三角メッシュ分割を行う
		// 三角メッシュの分割プログラムでは，L点を基準としたLR並びでパターン分けする
		// まずLRRLを捜す，次いでLRLを探す，その後LRRを捜す
		// L点が無くなったら任意のR点で扇形分割を行う
		// 三角メッシュ分割プログラムの呼び出しは，普通建物においては例外処理となる

		// 四角形の場合の処理
		// 傾斜屋根モデリングのプログラムに処理を渡す

		if nod2 == 4 {
			rect0 := SlopeRoof{ID: id, Elv: elv, List: cord2}
			rectList = addrect(rect0, rectList)
			// log.Println("rectList=", rectList)

			// ６角形の四角形分割
		} else if nod2 == 6 {
			_, rect1L, rect2L := pkg.HexaDiv(cord2, order)
			if rect1L == nil && rect2L == nil {
				log.Println("6角形を四角形分割できない", id, elv, cord2)
				poly := Polygon{ID: id, Elv: elv, List: cord2}
				polyList = addpoly(poly, polyList)
			} else {
				rect1 := SlopeRoof{ID: id, Elv: elv, List: rect1L}
				rectList = addrect(rect1, rectList)
				rect2 := SlopeRoof{ID: id, Elv: elv, List: rect2L}
				rectList = addrect(rect2, rectList)
				// log.Println("rectList=", rectList)
			}

			// ８角形の四角形分割
		} else if nod2 == 8 {
			_, rect1L, rect2L, rect3L, _ := pkg.OctaDiv(cord2, order, lrPtn, lrIdx)
			if rect1L == nil && rect2L == nil && rect3L == nil {
				log.Println("8四角形分割できない", id, elv, cord2)
				poly := Polygon{ID: id, Elv: elv, List: cord2}
				polyList = addpoly(poly, polyList)
			} else {
				// hex1Lは６角形の分割プログラムに渡されて四角形分割される
				rect1 := SlopeRoof{ID: id, Elv: elv, List: rect1L}
				rectList = addrect(rect1, rectList)
				rect2 := SlopeRoof{ID: id, Elv: elv, List: rect2L}
				rectList = addrect(rect2, rectList)
				rect3 := SlopeRoof{ID: id, Elv: elv, List: rect3L}
				rectList = addrect(rect3, rectList)
				// log.Println("rectList=", rectList)
			}

			// 10角形の四角形分割
		} else if nod2 == 10 {
			rect1L, rect2L, rect3L, rect4L := pkg.DecaDiv(cord2, order, lrPtn, lrIdx)
			if rect1L == nil && rect2L == nil && rect3L == nil && rect4L == nil {
				log.Println("10角形を四角形分割できない", id, elv, cord2)
				poly := Polygon{ID: id, Elv: elv, List: cord2}
				polyList = addpoly(poly, polyList)
			} else {
				// oct1Lは８角形の四角形分割プログラムに渡されて四角形分割される
				rect1 := SlopeRoof{ID: id, Elv: elv, List: rect1L}
				rectList = addrect(rect1, rectList)
				rect2 := SlopeRoof{ID: id, Elv: elv, List: rect2L}
				rectList = addrect(rect2, rectList)
				rect3 := SlopeRoof{ID: id, Elv: elv, List: rect3L}
				rectList = addrect(rect3, rectList)
				rect4 := SlopeRoof{ID: id, Elv: elv, List: rect4L}
				rectList = addrect(rect4, rectList)
				// log.Println("rectList=", rectList)
			}

			// 傾斜屋根モデリングできない多角形は３角メッシュ分割する
		} else {
			// 四角形分割ができなかった場合，三角メッシュに分割される
			// ポリゴンリストに追加する
			slR = false
			// poly := Polygon{ID: id, Elv: elv, List: cord2}
			// polyList = addpoly(poly, polyList)
		}

		if slR == false {
			// TODO:三角メッシュの分割プログラムに渡す
			poly := Polygon{ID: id, Elv: elv, List: cord2}
			polyList = addpoly(poly, polyList)
		}

		if er = scanner.Err(); er != nil {
			// エラー処理
			break
		}
	}

	log.Println(rectList)
	p90 := &rectList
	log.Printf("pointer:%p\n", p90)
	log.Println("rectList", unsafe.Sizeof(rectList))

	log.Println(polyList)
	p99 := &polyList
	log.Printf("pointer:%p\n", p99)
	log.Println("polyList", unsafe.Sizeof(polyList))

	// 四角形データファイルの作成
	fr, err := os.Create("./rectfile.gob")
	if err != nil {
		log.Fatal(err)
	}
	defer fr.Close()
	// エンコーダーの作成
	rEncoder := gob.NewEncoder(fr)
	// エンコード
	if err := rEncoder.Encode(rectList); err != nil {
		log.Fatal(err)
	}

	// 多角形データファイルの作成
	// 三角メッシュ分割は多角柱モデリングで行う
	// ここでは多角形データを出力する
	fp, erp := os.Create("./polyfile.gob")
	if erp != nil {
		log.Fatal(erp)
	}
	defer fp.Close()
	// エンコーダーの作成
	pEncoder := gob.NewEncoder(fp)
	// エンコード
	if erp := pEncoder.Encode(polyList); erp != nil {
		log.Fatal(erp)
	}

	// セーブファイルから復元
	flr, err := os.Open("./rectfile.gob")
	if err != nil {
		log.Fatal(err)
	}
	defer flr.Close()
	var rectList2 = make([]SlopeRoof, 0)
	// デコーダーの作成
	rDecoder := gob.NewDecoder(flr)
	// デコード
	if err := rDecoder.Decode(&rectList2); err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Println("rectList", rectList2)

	// セーブファイルから復元
	flp, erp := os.Open("./polyfile.gob")
	if erp != nil {
		log.Fatal(erp)
	}
	defer flp.Close()
	var polyList2 = make([]Polygon, 0)
	// デコーダーの作成
	pDecoder := gob.NewDecoder(flp)
	// デコード
	if erp := pDecoder.Decode(&polyList2); erp != nil {
		log.Fatal("decode error:", erp)
	}
	fmt.Println("polyList", polyList2)
}
