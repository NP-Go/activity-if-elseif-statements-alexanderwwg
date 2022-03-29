package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//Insert your code here
	number := 17
	//Hint if a:= ?? ; condition {  }
	oddOrEven(number)
	noOfDigits(number)
}
func oddOrEven(number int) {
	if number%2 == 0 {
		fmt.Printf("%v"+" is even.\n", number)
	} else if number%2 == 1 {
		fmt.Printf("%v"+" is odd. \n", number)
	}
}
func noOfDigits(number int) {
	number = int(math.Abs(float64(number)))
	numToString := strconv.Itoa(number)
	lengthOfString := len(numToString)
	fmt.Printf("Number of digits : %v.\n", lengthOfString)
}
