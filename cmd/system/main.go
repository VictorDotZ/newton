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

var numEquations int
var eps float64

func init() {
	flag.IntVar(&numEquations, "numEquations", 1, "initial approximation of a root")
	flag.Float64Var(&eps, "eps", 1e-2, "precision")
}

func main() {
	flag.Usage = usage
	flag.Parse()

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

	if err := root_finders.NewtonSystem(x[:], F, dF, numEquations, eps); err != nil {
		log.Fatalf("%v\n", err)
	} else {
		for _, elem := range x {
			fmt.Printf("%g\n", elem)
		}
	}
}
