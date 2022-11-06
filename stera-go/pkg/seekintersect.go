package pkg

// SeekInsec は２直線の交点を求める．
func SeekInsec(m1, n1, m2, n2 float64) (ansX, ansY float64) {
	ansX = (n2 - n1) / (m1 - m2)
	ansY = m1*ansX + n1
	return ansX, ansY
}
