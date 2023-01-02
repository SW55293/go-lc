package integers

import "fmt"

// := cannot be used outside of a function

func main() {

	fmt.Print("\nIntegers\n")
	ints()

}

func ints() {
	//integers
	//1
	var ageOne int = 10
	fmt.Println(ageOne)

	//2
	var ageTwo = 20
	fmt.Println(ageTwo)

	//3
	ageThree := 30
	fmt.Println(ageThree)
	fmt.Println()
}
