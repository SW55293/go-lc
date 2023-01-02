package bitsNmemory

import "fmt"

// := cannot be used outside of a function

func main() {

	fmt.Print("\nBits and Memory\n")
	bitsNmemory()

}

func bitsNmemory() {
	//bits
	//1  [int8 ranges from -128 -> 127]
	var firstBit int8 = 15
	fmt.Println(firstBit)

	//2
	var secondBit int8 = -128
	fmt.Println(secondBit)

	//Unsigned Ints
	//1 [Cannot be negative numbers only positives so 0->infinity]
	var unsignedOne uint = 2233
	fmt.Println(unsignedOne)

	//Floats/Decimals
	//1
	var scoreOne float32 = 21.36
	fmt.Println(scoreOne)

	//2
	var scoreTwo float64 = 333121212121212121.3
	fmt.Println(scoreTwo)
	fmt.Println()

}
