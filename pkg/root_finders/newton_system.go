package root_finders

import (
	"fmt"
	"math"

	"github.com/VictorDotZ/newton/pkg/matrix"
)

type vector = []float64
type nDimFunc = func(vector, vector, int)

func UniformMetric(vecOne, vecTwo vector, N int) float64 {
	max := 0.0
	current := 0.0
	for i := 0; i < N; i++ {
		current = math.Abs(vecTwo[i] - vecOne[i])
		if current > max {
			max = current
		}
	}
	return max
}

func MatrixOnVectorProduct(m, v, res vector, N int) {
	for i := 0; i < N; i++ {
		res[i] = 0.0
		for j := 0; j < N; j++ {
			res[i] += m[i*N+j] * v[j]
		}
	}
}

func NewtonSystem(x0 []float64, F, dF nDimFunc, numEquations int, eps float64) error {
	const stepLimit = 200
	x := make([]float64, numEquations)
	f := make([]float64, numEquations)
	jacobian := make([]float64, numEquations*numEquations)
	jacobianInv := make([]float64, numEquations*numEquations)

	for numSteps := 0; numSteps < stepLimit; numSteps++ {
		F(f[:], x0[:], numEquations)

		dF(jacobian[:], x0[:], numEquations)

		// (F'(x_n))^-1
		if err := matrix.GaussMaxCol(jacobian[:], jacobianInv[:], numEquations); err != nil {
			return fmt.Errorf("jacobian degenerated")
		}

		// (F'(x_n))^-1 * F(x_n)
		MatrixOnVectorProduct(jacobianInv, f, x[:], numEquations)

		// x_{n+1} = x_n - (F'(x_n))^-1 * F(x_n)
		for i := 0; i < numEquations; i++ {
			x[i] = x0[i] - x[i]
		}

		if UniformMetric(x0, x, numEquations) < eps {
			return nil
		}

		for i := 0; i < numEquations; i++ {
			x0[i] = x[i]
		}
	}

	return fmt.Errorf("numSteps limit")
}
