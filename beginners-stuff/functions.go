package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("\nFunctions with 1 argument\n\n")
	oneArgument("The")
	oneArgument("String")

	fmt.Print("\nFunctions with 2 arguments\n\n")
	twoArguments([]string{"cloud", "rain", "tree"}, oneArgument)

	fmt.Print("\nReturn a Value\n\n")
	areaOne := returnValue(11)
	fmt.Printf("the radius of circle 1 is = %v\n", areaOne)

}

func oneArgument(x string) {
	//1 takes in 1 string
	fmt.Printf("The string that was passed in is: %v \n", x)
}

func twoArguments(a []string, f func(string)) {
	//1 [Going thru the slice of strings and calling the value v which is of type string. The function we are passing it is the oneArgument]
	for _, v := range a {
		f(v)
	}
}

func returnValue(radius float64) float64 {
	//1
	return math.Pi * radius * radius
}
