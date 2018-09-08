package fft

import (
	"errors"
	"math"
)

func Transform(samples *[]complex128) error {
	n := len(*samples)

	if n <= 1 {
		return errors.New("Invalid input")
	}

	a0 := make([]complex128, n/2)
	a1 := make([]complex128, n/2)

	for i, j := 0, 0; i < n; i, j = i+2, j+1 {
		a0[j] = (*samples)[i]
		a1[j] = (*samples)[i+1]
	}

	Transform(&a0)
	Transform(&a1)

	angle := 2 * math.Pi / float64(n)

	w := complex128(1)
	wn := complex(math.Cos(angle), math.Sin(angle))

	for i := 0; i < n/2; i++ {
		(*samples)[i] = a0[i] + w*a1[i]
		(*samples)[i+n/2] = a0[i] - w*a1[i]
		w *= wn
	}

	return nil
}
