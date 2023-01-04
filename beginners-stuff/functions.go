package main

import "fmt"

func main() {
	fmt.Print("\nFunctions with 1 argument\n\n")
	oneArgument("The")
	oneArgument("String")

	fmt.Print("\nFunctions with 2 arguments\n\n")
	twoArguments([]string{"cloud", "rain", "tree"}, oneArgument)

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
