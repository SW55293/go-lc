package main

import "fmt"

func main() {

	fmt.Print("\nBooleans\n\n")
	bools()

	fmt.Print("\nConditionals\n\n")
	conds()

}

func bools() {

	//1
	age := 45
	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)
}

func conds() {

	//1 [if else statements]
	score := 15
	if score < 10 {
		fmt.Println("less than 10")
	} else if score < 30 {
		fmt.Println("less than 30")
	} else {
		fmt.Println("Passed")
	}

	//2 [continue and break example]
	names := []string{"xxx", "yy", "pool", "time", "sky", "black"}
	for index, value := range names {
		if index == 1 {
			fmt.Printf("Not printing out the string at position %v we are continuing\n", index)
			continue
		}
		if index > 3 {
			fmt.Printf("Going to break at position %v ", index)
			break
		}
		fmt.Printf("the value at position %v is %v \n", index, value)
	}

}
