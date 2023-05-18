package root_finders

import (
	"math"
	"testing"
)

func TestFirst(t *testing.T) {
	numEquations := 3
	eps := 1e-10

	x := make([]float64, numEquations)

	func(x []float64, _ int) {
		x[0] = 0
		x[1] = 0.5
		x[2] = 1.0
	}(x[:], numEquations)

	F := func(f, x []float64, _ int) {
		f[0] = math.Pow(x[0], 2) + math.Pow(x[1], 2) + math.Pow(x[2], 2) - 10
		f[1] = -math.Pow(x[0], 2) + math.Pow(x[1]-3, 2) - x[2]
		f[2] = x[0] + 2. - x[1]/3. + x[2]
	}

	dF := func(J []float64, x []float64, numEquations int) {
		J[0*numEquations+0] = 2.0 * x[0]
		J[0*numEquations+1] = 2.0 * x[1]
		J[0*numEquations+2] = 2.0 * x[2]
		J[1*numEquations+0] = -2.0 * x[0]
		J[1*numEquations+1] = 2.0 * (x[1] - 3.0)
		J[1*numEquations+2] = -1.0
		J[2*numEquations+0] = 1.0
		J[2*numEquations+1] = -1.0 / 3.0
		J[2*numEquations+2] = 1.0
	}

	groundTruth := []float64{-3.01007258012015, -0.16475754761303, 0.95515339758247}

	if err := NewtonSystem(x[:], F, dF, numEquations, eps); err != nil {
		t.Fatalf("%v\n", err)
	} else {
		for i := range x {
			if math.Abs(x[i]-groundTruth[i]) > eps {
				t.Errorf("too large error")
			}
		}
	}
}

func TestSecond(t *testing.T) {
	numEquations := 2
	eps := 1e-10

	x := make([]float64, numEquations)

	func(x []float64, _ int) {
		x[0] = 0.5
		x[1] = 0.5
	}(x[:], numEquations)

	F := func(f, x []float64, _ int) {
		f[0] = math.Pow(x[0], 2) - math.Pow(x[1], 2) - 1
		f[1] = math.Pow(x[0], 2) + math.Pow(x[1], 2) - 15
	}

	dF := func(J []float64, x []float64, numEquations int) {
		J[0*numEquations+0] = 2.0 * x[0]
		J[0*numEquations+1] = -2.0 * x[1]
		J[1*numEquations+0] = 2.0 * x[0]
		J[1*numEquations+1] = 2.0 * x[1]
	}

	groundTruth := []float64{2.82842712474619, 2.64575131106459}
	if err := NewtonSystem(x[:], F, dF, numEquations, eps); err != nil {
		t.Fatalf("%v\n", err)
	} else {
		for i := range x {
			if math.Abs(x[i]-groundTruth[i]) > eps {
				t.Errorf("too large error")
			}
		}
	}
}
