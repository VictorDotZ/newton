package root_finders

import (
	"fmt"
	"math"
)

func NewtonRaphson(x0, z *float64, f, df func(float64) float64, eps float64) error {
	const stepLimit = 200
	const machineEps float64 = 1e-14
	var x1 float64
	var f0 float64
	var df0 float64

	if z != nil {
		fmt.Printf("step\tx0\tf(x0)\tdf(x0)\tAbs(x1-z)\n")
	}

	for numSteps := 0; numSteps < stepLimit; numSteps++ {
		f0 = f(*x0)
		df0 = df(*x0)

		if math.Abs(df0) < machineEps {
			return fmt.Errorf("fail: derivative near to 0")
		}

		x1 = *x0 - (f0 / df0)

		if math.Abs(x1-*x0) < eps {
			return nil
		}

		if z != nil {
			fmt.Printf("%d\t%.10f\t%.10f\t%.10f\t%.10f\n", numSteps, *x0, f0, df0, math.Abs(x1-*z))
		}

		*x0 = x1
	}

	return fmt.Errorf("fail: step limit")
}
