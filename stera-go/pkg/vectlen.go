package pkg

import (
	"math"
)

// VectLen はベクトルの長さを求める．
func VectLen(oneX, oneY, intX, intY float64) (vectX, vectY, vectL float64) {
	// ベクトルのX座標の差分
	vectX = math.Abs(oneX - intX)
	// ベクトルのY座標の差分
	vectY = math.Abs(oneY - intY)
	// ベクトルの長さ
	vectL = math.Sqrt(math.Pow(vectX, 2) + math.Pow(vectY, 2))

	return vectX, vectY, vectL
}
