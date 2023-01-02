package main

import "fmt"

func main() {

	fmt.Print("\nArrays\n")
	arrays()

	fmt.Print("\nSlices\n")
	slices()

}

func arrays() {
	//1 [Empty - Explicitly tell it it will be an array]
	var ages [5]int
	fmt.Println("Empty Array = ", ages)

	//2 [Filled - Explicitly tell it to be an array]
	var ageTwo [3]int = [3]int{1, 2, 5}
	fmt.Println("Filled Array = ", ageTwo)

	//3 [Filled and shorthand of the above]
	var ageShortHand = [4]int{5, 8, 9, 3}
	fmt.Println("Short hand and Filled array = ", ageShortHand)

	//Printing out the length of the arrays
	fmt.Println("Empty array length = ", len(ages))
	fmt.Println("Filled array length = ", len(ageTwo))
	fmt.Println("Shorthanded array length = ", len(ageShortHand))

	//4 shorthanded with inffered type using :=
	names := [2]string{"xaa", "kll"}
	fmt.Printf("string array contains: %v and the length of the array is %v", names, len(names))
	//changing the value of a string in the array to a different one
	names[0] = "changedName"
	fmt.Println("the changed string array: ", names)
	fmt.Println()
}

//slices are arrays that can be manipulated and changed. They are not fixed liked regular arrays
func slices() {
	//1 [Slice example with way to change a variable inside the array]
	var scores = []int{8, 9, 0}
	fmt.Println("Original Slice = ", scores)
	scores[1] = 4
	fmt.Println("Changed Slice = ", scores)

	//Appending data to a slice. This changes the orginial slice and appends
	scores = append(scores, 120, 10)
	fmt.Println("appended data to original slice = ", scores)
	fmt.Println("The length of the new appended slice is = ", len(scores))

	//Slice Ranges
	//1 [Print out a certain range]
	rangeOne := scores[1:2]
	fmt.Println("Print out range 1->2 of the array = ", rangeOne)
	//[1:2] The way you read this is [Include this Position : Up to this Position but do not include what is found here]

	//2 [Print out from a point to no end]
	rangeTwo := scores[1:]
	fmt.Println("Print out range 1->end of the array = ", rangeTwo)

	//3 [Print out from no beginning to a point]
	rangeThree := scores[:4]
	fmt.Println("Print out range noBeginning->4 of the array = ", rangeThree)

	//appending to a slice
	fmt.Println("Original Slice of range 1 = ", rangeOne)
	rangeOne = append(rangeOne, 20)
	fmt.Println("Appending to range 1 slice = ", rangeOne)

}
