package pkg

import "math"

// TriArea は三角形の面積を計算する
func TriArea(x1, y1, x2, y2, x3, y3 float64) (area float64) {
	Bx := math.Abs(x2 - x1)
	By := math.Abs(y2 - y1)
	Cx := math.Abs(x3 - x1)
	Cy := math.Abs(y3 - y1)
	// 三角形の面積公式
	area = math.Abs(Bx*Cy-By*Cx) / 2

	return area
}
