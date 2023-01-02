package main

import "fmt"

// := cannot be used outside of a function

func main() {

	fmt.Print("\nFormatted strings\n")
	printingFormattedStrings()

}

func printingFormattedStrings() {
	//Print [doesnt add new line]
	fmt.Print("steph,")
	fmt.Print("no new line unless you add \n")

	//Println and outputting variables
	age := 100
	name := "steph"
	fmt.Println("age is", age, "and name is", name)

	//Formatted & Embedded strings
	//1 [%v->variable]
	fmt.Printf("my age is %v and name is %v \n", age, name)
	//%v = variable. It finds the first variable and replaces it with what comes first in the print statement

	//2 [%q->quotations aroun strings. doesnt work with ints]
	fmt.Printf("name is %q and age is %q \n", name, age)

	//3 [%T-> tells us the type of variable it is. It works with strings and ints]
	fmt.Printf("variable age is of type = %T \n", age)

	//4 [%f-> prints out float variables]
	fmt.Printf("the score is = %f \n", 20.36)
	//variation of the above %f but with a limit on the numbers outputted after the decimal
	//0 numbers after the decimal
	fmt.Printf("the score is = %0.0f \n", 20.36)
	//1 number after the decimal
	fmt.Printf("the score is = %0.1f \n", 20.36)
	//2 numbers after the decimal
	fmt.Printf("the score is = %0.2f \n", 20.36)

	//5 [Sprintf -> saving formatted strings. Things makes the whole print statement into 1 variable]
	var str = fmt.Sprintf("my age is %v and name is %v \n", age, name)
	fmt.Println("The saved string is in here:", str)

}
