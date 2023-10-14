package matchengine

import (
	"golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](vals ...T) (min T) {
	if len(vals) == 0 {
		return min
	}

	min = vals[0]

	for i := 1; i < len(vals); i++ {
		if vals[i] < min {
			min = vals[i]
		}
	}

	return min
}

func Max[T constraints.Ordered](vals ...T) (max T) {
	if len(vals) == 0 {
		return max
	}

	max = vals[0]

	for i := 1; i < len(vals); i++ {
		if vals[i] < max {
			max = vals[i]
		}
	}

	return max
}

func ReverseSlice[T any](vals []T) {
	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		vals[i], vals[j] = vals[j], vals[i]
	}
}
