package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 1
	// var intNum16 int16 = 32767
	// var intNum16 int16 = 32767 + 1
	var intNum8 uint8 = 127
	var uIntNum8 uint8 = 128
	fmt.Println(intNum, intNum8, uIntNum8)

	var floatNum32 float32 = 123.45677895574
	fmt.Println(floatNum32)

	// var floatNum64 float64 = 1.0012
	// var intNum64 int64 = 123456
	// var result float32 = floatNum64 + intNum64
	// fmt.Println(result)

	var myString string = "Hello" + " " + "World!"
	fmt.Println(myString)
	fmt.Println(len("☯"))
	fmt.Println(utf8.RuneCountInString("☯"))

	var myRune rune = 'a' // ?
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var defaultInt int
	var defaultString string
	var defaultBoolean bool
	fmt.Printf("defaultInt: %d\n", defaultInt)
	fmt.Printf("defaultString: %s\n", defaultString)
	fmt.Printf("defaultBoolean: %t\n", defaultBoolean)

	myVar := "myVar" // why tho
	fmt.Println(myVar)

	const myConst string = ""
	fmt.Println("myConst: " + myConst)

	const pi float32 = 3.14159
	fmt.Printf("const pi: %f\n", pi)

	// functions 
	// functions 

	var printValue string = "Hello World!"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2

	var result,remainder, error = intDivision(numerator, denominator)

	if error != nil {
		fmt.Println(error.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v\n", result)
	} else {
		fmt.Printf("The result of the integer division is %v and the remainder is %v\n", result, remainder)
	}

	// SWITCH STATEMENT
	// SWITCH STATEMENT

	switch {
		case error != nil:
			fmt.Println(error.Error())
		case remainder == 0:
			fmt.Printf("SWITCH: The result of the integer division is %v\n", result)
		default:
			fmt.Printf("SWITCH: The result of the integer division is %v and the remainder is %v\n", result, remainder)
	}

	// CONDITIONAL SWITCH STATEMEN
	// CONDITIONAL SWITCH STATEMEN

	switch remainder {
		case 9:
			fmt.Println("CONDITIONAL SWITCH: The division was exact, %modulus=0")
		case 1, 2: 
			fmt.Println("CONDITIONAL SWITCH: The division was close, %modulus=1||2")
		default:
			fmt.Println("CONDITIONAL SWITCH: The division was far off, %modulus>=3")
	}
	
	// AND-OPERATOR
	// AND-OPERATOR
	if 1 == 1 && 2 == 2 {
		fmt.Println("1 == 1 && 2 == 2")
	}

	// OR-OPERATOR
	// OR-OPERATOR
	if 3 == 2 || 3 == 3 {
		fmt.Println("3 == 2 || 3 == 3")
	}

	// ARRAYS - fixed length, same type, indexable, contiguoous in memory 
	// ARRAYS - fixed length, same type, indexable, contiguoous in memory
	
	var intArry [3]int32
	fmt.Println(intArry)
}

// FUNCTIONS
// FUNCTIONS

func printMe(printValue string) {
	fmt.Println("func: " + printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var error error

	if denominator == 0 {
		error = errors.New("Cannot divide by zero.") 
		return 0, 0, error
	}

	var result int = numerator / denominator
	var remainder int = numerator%denominator
	

	return result, remainder, error
}
















