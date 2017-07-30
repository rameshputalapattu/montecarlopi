// Montecarlo project main.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//This function calculates Pi using
//Aryabhatta's rule
//Add 4 to 100, Multiply by 8 and add 62000
//By this rule you approach the circumference of
//a circle with a diameter 20000
func calcPiAB() float64 {
	return ((100.0+4.0)*8 + 62000.0) / 20000.0
}

//This function calculates Pi using
//Montecarlo method.
//This method involves shooting random
//darts uniformly on an unit square with a quarter
//circle inscribed in it. The ratio of the
//quarter circle's area to the square is Pi/4.
//So the value of Pi can be estimated as below.
//Pi = 4 *  (darts inside/total number of darts)
func calcPiMC(numDarts int) float64 {
	//seeding the random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var numInside int
	for i := 0; i < numDarts; i++ {
		//shoot the darts at random
		x := r.Float64()
		y := r.Float64()

		//count the points inside the unit circle
		if x*x+y*y < 1.0 {
			numInside++
		}
	}

	return 4.0 * float64(numInside) / float64(numDarts)

}

func main() {
	fmt.Printf("The value of Pi through Aryabhatta's rule is %.4f\n", calcPiAB())
	fmt.Printf("The value of Pi through Montecarlo method is %.4f\n", calcPiMC(100000))
}
