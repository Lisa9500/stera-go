package pkg

import (
	"math"
)

// CrossAngl はベクトルの交差角度を求める
func CrossAngl(ax, ay, bx, by, al, bl float64) (deg float64) {
	// ３つ目の辺の長さを求める
	// tHen = math.Sqrt(math.Pow((ax - bx), 2) + math.Pow((ay - by), 2))
	tHen := DistVerts(ax, bx, ay, by)
	// 余弦定理　第二余弦定理を変形した公式を使えば、辺の長さから余弦を求める
	// cosθを求める
	cosT := (math.Pow(al, 2) + math.Pow(bl, 2) - math.Pow(tHen, 2)) / (2 * al * bl)
	// log.Println("cosθ", cosT)
	// cosθからアークコサインで角度（ラジアン→度）を求める
	deg = math.Acos(cosT) * (180 / math.Pi)
	// log.Println("角度", deg)
	if deg < 0 {
		// TODO:
	}
	return deg
}
