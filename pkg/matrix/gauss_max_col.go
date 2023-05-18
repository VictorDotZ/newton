package matrix

import (
	"fmt"
	"math"
)

func GaussMaxCol(A, AInv []float64, N int) error {
	var i, j, k int
	var maxElemByCols float64 = 0.0
	var maxElemColIdx int = 0
	const err float64 = 1e-15
	var c float64 = 0.

	// make AInv singular
	for i = 0; i < N; i++ {
		for j = 0; j < N; j++ {
			if i == j {
				AInv[i*N+j] = 1.0
			} else {
				AInv[i*N+j] = 0.0
			}
		}

	}

	for j = 0; j < N; j++ {
		maxElemByCols = math.Abs(A[j*N+j])
		maxElemColIdx = j

		// find max elem between columns
		for i = j; i < N; i++ {
			if math.Abs(A[i*N+j]) > maxElemByCols {
				maxElemColIdx = i
				maxElemByCols = math.Abs(A[i*N+j])
			}
		}

		if maxElemByCols < err {
			return fmt.Errorf("degenerate matrix\n")
		}

		// swap current row with row with max elem
		if maxElemColIdx != j {
			for i = 0; i < j; i++ {
				AInv[j*N+i], AInv[maxElemColIdx*N+i] = AInv[maxElemColIdx*N+i], AInv[j*N+i]
			}
			for i = j; i < N; i++ {
				A[j*N+i], A[maxElemColIdx*N+i] = A[maxElemColIdx*N+i], A[j*N+i]
				AInv[j*N+i], AInv[maxElemColIdx*N+i] = AInv[maxElemColIdx*N+i], AInv[j*N+i]
			}
		}

		// divide all row by first elem
		c = 1. / A[j*N+j]
		for i = 0; i < j; i++ {
			AInv[j*N+i] *= c
		}
		for i = j; i < N; i++ {
			A[j*N+i] *= c
			AInv[j*N+i] *= c
		}

		// substract all elements under diagonal
		for i = j + 1; i < N; i++ {
			c = A[i*N+j]
			for k = 0; k < j; k++ {
				AInv[i*N+k] -= c * AInv[j*N+k]
			}
			for k = j; k < N; k++ {
				A[i*N+k] -= c * A[j*N+k]
				AInv[i*N+k] -= c * AInv[j*N+k]
			}
		}
	}

	// reverse step
	for j = N; j >= 1; j-- {
		for i = 0; i < j-1; i++ {
			for k = 0; k < N; k++ {
				AInv[i*N+k] -= A[i*N+j-1] * AInv[(j-1)*N+k]
			}
			A[i*N+j-1] = 0.
		}
	}

	return nil
}
