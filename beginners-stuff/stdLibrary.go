package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Print("\nString Library/Package\n\n")
	stringPkg()

	fmt.Print("\nSort Library/Package\n\n")
	sortPkg()

}

func stringPkg() {

	//uses fmt
	greeting := "hello everyone how is it going y'all"
	fmt.Println(greeting)

	//1 strings.Contains [looks for string in a variable that you give it then returns true or false if it cotains it]
	fmt.Println(strings.Contains(greeting, "poop"))

	//2 strings.ReplaceAll [replaces one term with another one]
	fmt.Println(strings.ReplaceAll(greeting, "everyone", "poop"))
	//the original string greetings is not altered
	fmt.Println(greeting)

	//3 strings.ToUpper [makes the string all upper case]
	fmt.Println(strings.ToUpper(greeting))

	//4 strings.Index [returns the index of a substring we give it. Gives you the strating index of the substring]
	fmt.Println(strings.Index(greeting, "ev"))
	//returns the index of the first one it finds
	fmt.Println(strings.Index(greeting, "e"))

	//5 strings.Split [splits/removes a string on substring/letter you give it and returns an array with letter/substring removed]
	fmt.Println(strings.Split(greeting, "h"))

	fmt.Println()
}

func sortPkg() {

	ages := []int{2, 50, 4, 88, 238, 45, 1, 0, 56, 33333}

	//1 sort.Ints [sorts the slice or array you give it in ascendign order]
	sort.Ints(ages)
	fmt.Println("Sorted slice in ascending order = ", ages)

	//2 sort.SearchInts [searched for a number in a slice and returns the index if it finds it.
	// If its not there it gives you the index of where to place it]
	index := sort.SearchInts(ages, -9)
	fmt.Println(index)

	//3 sort.Strings [sorts the string in alphabetical order and places it in a array/slice]
	names := []string{"a", "b", "x", "ll"}
	sort.Strings(names)
	fmt.Println(names)

	//4 sort.SearchString [searched the string for a letter/substring you give it and outputs the index. If it isnt there
	// it tells you at what index you should place it]
	fmt.Println(sort.SearchStrings(names, "name"))

}
