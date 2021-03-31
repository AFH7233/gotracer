package utils

type Transformation [4][4]float64

func (a Transformation) Combine(b Transformation) (result Transformation) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return
}
