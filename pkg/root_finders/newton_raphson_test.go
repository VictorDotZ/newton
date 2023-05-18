package root_finders

import (
	"fmt"
	"math"
	"testing"
)

func Test1DAnalytic(t *testing.T) {
	x0 := 1.0001
	eps := 1e-7

	f := func(x float64) float64 {
		return math.Pow(x-1.0, 2)
	}

	df := func(x float64) float64 {
		return 2.0 * (x - 1.0)
	}

	if err := NewtonRaphson(&x0, nil, f, df, eps); err != nil {
		t.Fatalf("%v", err)
	} else {
		fmt.Println(x0)
		if math.Abs(x0-1.0000001953125) > 1e-16 {
			t.Fatalf("so much error")
		}
	}
}

func Test1DNumerical(t *testing.T) {
	x0 := 1.6
	eps := 1e-7

	f := func(x float64) float64 {
		return math.Pow(x, 3.) * math.Pow(math.Cos(x), 2)
	}

	h := 1e-5
	df := func(x float64) float64 {
		return (f(x+h) - f(x-h)) / (2.0 * h)
	}

	if err := NewtonRaphson(&x0, nil, f, df, eps); err != nil {
		t.Fatalf("%v", err)
	} else {
		if math.Abs(x0-1.5707964445042175) > 1e-16 {
			t.Fatalf("so much error")
		}
	}
}
