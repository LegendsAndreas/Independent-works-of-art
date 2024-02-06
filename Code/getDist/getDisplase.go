package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var acce, velo, displa, time float64

	fmt.Printf("Please enter acceleration> ")
	_, err := fmt.Scanln(&acce)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Please enter initial velocity> ")
	_, err1 := fmt.Scanln(&velo)
	if err != nil {
		log.Fatal(err1)
	}

	fmt.Printf("Please enter initial displacement> ")
	_, err2 := fmt.Scanln(&displa)
	if err != nil {
		log.Fatal(err2)
	}

	fmt.Printf("Please enter time> ")
	_, err3 := fmt.Scanln(&time)
	if err != nil {
		log.Fatal(err3)
	}

	// Sets displacefn to be equal to the function in genDisplaceFn. After which, the result gets printed.
	displaceFn := genDisplaceFn(acce, velo, displa)
	fmt.Printf("Displacement is: %f", displaceFn(time))

}

/*
Creates a function, that can be assigned to a variable.
The function "fn" returns the formula for displacement: s = ½a(t^2)+vot+so
*/
func genDisplaceFn(a float64, v float64, s float64) func(float64) float64 {
	fn := func(t float64) float64 {
		var displacement float64 = 0.5*a*math.Pow(t, 2) +
			v*t +
			s
		return displacement
	}
	return fn
}
