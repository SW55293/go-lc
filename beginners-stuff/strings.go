package strings

import "fmt"

// := cannot be used outside of a function

func main() {
	fmt.Print("Strings\n")
	strings()

}

func strings() {
	//strings
	//1
	var name string = "steph"
	fmt.Println(name)

	//2
	var nameTwo = "poop"
	fmt.Println(nameTwo)

	//3
	var nameThree string
	fmt.Println(nameThree)

	//4
	nameFour := "with no string declaration"
	fmt.Println(nameFour)

	//changing the values of the set strings
	name = "new Steph"
	nameTwo = "new poop"
	nameThree = "first name"
	fmt.Println(name, nameTwo, nameThree)
	fmt.Println()
}
