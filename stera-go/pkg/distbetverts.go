package pkg

import (
	"math"
)

// DistVerts は頂点間の距離を求める
func DistVerts(x1, x2, y1, y2 float64) (dist float64) {
	dist = math.Sqrt(math.Pow((x1-x2), 2) + math.Pow((y1-y2), 2))
	// log.Println("dist=", dist)
	return dist
}
