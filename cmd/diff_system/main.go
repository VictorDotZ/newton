package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/VictorDotZ/newton/pkg/root_finders"
)

func usage() {
	fmt.Println("usage:")
	flag.PrintDefaults()
	os.Exit(2)
}

var eps float64
var rightBound float64
var a float64
var b float64
var numEquations int

func init() {
	flag.Float64Var(&eps, "eps", 1e-10, "precision")
	flag.Float64Var(&rightBound, "rightBound", 1.0, "right bound")
	flag.Float64Var(&a, "a", 1.0, "y(0)")
	flag.Float64Var(&b, "b", 1.0, "y(rightBound)")
	flag.IntVar(&numEquations, "numEquations", 10, "initial approximation of a root")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	x := make([]float64, numEquations)

	func(x []float64, numEquations int) {
		for i := 0; i < numEquations; i++ {
			x[i] = 0.0
		}
	}(x[:], numEquations)

	f := func(x, y float64) float64 {
		return math.Cos(x) * math.Sin(y)
	}

	df_dy := func(x, y float64) float64 {
		return math.Cos(x) * math.Cos(y)
	}

	F := func(y, x []float64, _ int) {
		h := rightBound / (float64)(numEquations+1)
		divHh := 1.0 / h / h

		y[0] = (x[1]-2.*x[0]+a)*divHh - f(h, x[0])

		y[numEquations-1] = (b-2.*x[numEquations-1]+x[numEquations-2])*divHh - f(rightBound-h, x[numEquations-1])

		for k := 1; k < numEquations-1; k++ {
			y[k] = (x[k+1]-2.*x[k]+x[k-1])*divHh - f(h*(float64)(k+1), x[k])
		}
	}

	dF := func(J []float64, x []float64, numEquations int) {
		h := rightBound / (float64)(numEquations+1)
		mainDiag := -2. / h / h
		subDiag := 1. / h / h
		supDiag := subDiag

		for i := 0; i < numEquations*numEquations; i++ {
			J[i] = 0.
		}

		J[0*numEquations+0] = mainDiag - df_dy(h, x[0])
		J[0*numEquations+1] = supDiag
		J[(numEquations-1)*numEquations+numEquations-1] = mainDiag - df_dy(rightBound-h, x[numEquations-1])
		J[(numEquations-1)*numEquations+numEquations-2] = subDiag

		for k := 1; k < numEquations-1; k++ {
			J[k*numEquations+k-1] = subDiag
			J[k*numEquations+k+0] = mainDiag - df_dy((float64)(k+1)*h, x[k])
			J[k*numEquations+k+1] = supDiag
		}
	}

	if err := root_finders.NewtonSystem(x[:], F, dF, numEquations, eps); err != nil {
		log.Fatalf("%v\n", err)
	} else {
		h := rightBound / (float64)(numEquations+1)
		fmt.Printf("%g %g\n", 0., a)
		for k := 0; k < numEquations; k++ {
			fmt.Printf("%g %g\n", (float64)(k+1)*h, x[k])
		}
		fmt.Printf("%g %g\n", rightBound, b)
	}
}
