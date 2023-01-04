package main

import "fmt"

func main() {

	fmt.Print("\nFor Loops\n\n")
	forLoops()

}

func forLoops() {

	//1 Simple for loop that prints out 0->4
	x := 0
	for x < 5 {
		fmt.Println("value of x is: ", x)
		x++
	}
	fmt.Println()

	//2 Different way to do the same as above but in 1 line
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is: ", i)
	}
	fmt.Println()

	//3 Cycle thru a slice of strings with a for loop and print out the position of x
	names := []string{"voo", "look", "cat", "potato", "dog"}
	for x := 0; x < len(names); x++ {
		fmt.Println("Names in the slice:", names[x])
	}
	fmt.Println()

	//4 Another variation but using the Range term in the for loop using index and value
	namesTwo := []string{"juno", "ruby", "sunny", "annoying", "benji", "bubba"}
	for index, value := range namesTwo {
		fmt.Printf("index = %v value = %v \n", index, value)
	}
	fmt.Println()

	//5 Printing out just the value while ignoring the index
	namesThree := []string{"juno", "ruby", "sunny", "annoying", "benji", "bubba"}
	for _, value := range namesThree {
		fmt.Printf("value is %v \n", value)
	}
}
