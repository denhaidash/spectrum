package helper

func FindMaxValue(m map[int]float64) float64 {
	max := 0.0

	for _, v := range m {
		if v > max {
			max = v
		}
	}

	return max
}

func FindNearestBiggerPowerOf2(x uint) uint {
	x = x - 1
	x = x | (x >> 1)
	x = x | (x >> 2)
	x = x | (x >> 4)
	x = x | (x >> 8)
	x = x | (x >> 16)

	return x + 1
}
