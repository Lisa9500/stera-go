package pkg

import (
	"math"
)

// CalcTheta は余弦定理により角度を求める．
func CalcTheta(vecAX, vecAY, vecBX, vecBY, vecA, vecB float64) (theta float64) {
	// 対辺の長さを求める
	taihen := DistVerts(vecAX, vecBX, vecAY, vecBY)
	// log.Println("taihen=", taihen)
	// 余弦定理　第二余弦定理を変形した公式を使えば，辺の長さから余弦を求めることができる．
	cosTheta := (math.Pow(vecA, 2) + math.Pow(vecB, 2) - math.Pow(taihen, 2)) /
		(2 * vecA * vecB)
	// log.Println("cosθ=", cosTheta)
	// 角度を求める
	theta = (math.Acos(cosTheta)) * 180 / math.Pi
	// log.Println("角度=", theta)
	if theta < 0 {
		// TODO:
	}

	return theta
}
