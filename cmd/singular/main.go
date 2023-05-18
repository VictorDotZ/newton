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

var x0 float64
var eps float64
var z *float64

func init() {
	flag.Float64Var(&x0, "x0", 0.0, "initial approximation of a root")
	flag.Float64Var(&eps, "eps", 1e-2, "precision")
	z = flag.Float64("z", math.Inf(-1), "known root (optional)")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	// I do not know how to handle optional flag with less code
	if *z == math.Inf(-1) {
		z = nil
	}

	f := func(x float64) float64 {
		// return math.Pow(x-1.0, 2)
		return math.Pow(x, 3.) * math.Pow(math.Cos(x), 2)
	}

	// df := func(x float64) float64 {
	// 	return 2.0 * (x - 1.0)
	// }

	h := 1e-5
	df := func(x float64) float64 {
		return (f(x+h) - f(x-h)) / (2.0 * h)
	}

	if err := root_finders.NewtonRaphson(&x0, z, f, df, eps); err != nil {
		log.Fatalf("%v", err)
	} else {
		fmt.Printf("%.10f\n", x0)
	}
}
