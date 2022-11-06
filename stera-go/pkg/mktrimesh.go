package pkg

import (
	"log"
	"strings"
)

// MkTriMesh は三角メッシュのリストを作成する
func MkTriMesh(cord2 [][]float64, lrPtn []string, lrIdx []int) {
	// 頂点データ数の確認
	nod := len(cord2)
	if nod <= 3 {
		// TODO:
		return
	}
	// LR並びの確認　L点から始まっていなければエラー
	if lrPtn[0] != "L" {
		// TODO:
	}
	// 検索用にLR並びから半角スペースを除く
	lrjoin := strings.Join(lrPtn, "")
	log.Println("lrjoin=", lrjoin)

	// LRRL四角形を三角メッシュに分割する
	if strings.Contains(lrjoin, "LRRL") {
		// LRRL四角形の開始位置を見つける
		idx := strings.Index(lrjoin, "LRRL")
		s := lrIdx[idx]
		// LRRL四角形の頂点座標を取り出す
		rect := cord2[s : s+4]
		// LRR三角形とRRL三角形に分割する
		tri1 := rect[0:3]
		tri2 := rect[1:4]
		log.Println("tri1=", tri1)
		log.Println("tri2=", tri2)
		// 三角メッシュのリストに追加する

	}
}
