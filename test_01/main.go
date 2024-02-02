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
	
	var intArr [3]int32
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	intArr[1] = 22
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intArray [4]int32 = [4]int32{1, 2, 3, 4}
	fmt.Println(intArray)
	intArray2 := [4]float32{3, 2, 1} 
	fmt.Println(intArray2)
	intArray3 := [...]float32{3, 2, 1} 
	fmt.Println(intArray3)

	// SLICES
	// SLICES

	var intSlice []float64 = []float64{7, 8, 9}
	fmt.Printf("intSlice: %v\n", intSlice)
	fmt.Printf("Length = %v\nCapacity is %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 10)
	fmt.Printf("intSlice: %v\n", intSlice)
	fmt.Printf("Length = %v\nCapacity is %v\n", len(intSlice), cap(intSlice))
	// fmt.Println(intSlice[4]) // runtime error: index out of range [4] with length 4

	var intSlice2 []float64 = []float64{55, 56, 57}
	fmt.Println(intSlice2)
	intSlice2 = append(intSlice2, intSlice... )
	fmt.Println(intSlice2)

	// MAP
	// MAP

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)
	
	var myMap2 = map[string]uint8{"Jon" : 17, "Knut": 1}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Jon"])
	fmt.Println(myMap2["finnesIkke"])
	var age, ok = myMap2["finnesIkke"]
	if ok {
		fmt.Println(age, ok)
		} else {
			fmt.Printf("age: %v - doesExist: %v\n", age, ok)
	}	
	
	for name, age := range myMap2{
		fmt.Printf("forLoopMap - name: %v, age: %v\n", name, age)
	}

	delete(myMap2, "Knut")
	
	for name, age := range myMap2{
		fmt.Printf("forLoopMap(after delete()) - name: %v, age: %v\n", name, age)
	}

	for index, value := range intArr{
		fmt.Printf("forLoopArr - index: %v, value: %v\n", index, value)
	}

	// WHILE LOOP
	// WHILE LOOP

	var i uint8 = 0
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 100; i <= 105; i++ {
		fmt.Println(i)
	}

	// MATHS
	// MATHS

	number := 10
	number--
	fmt.Println(number)
	number++
	fmt.Println(number)
	number += 10
	fmt.Println(number)
	number -= 10
	fmt.Println(number)
	number *= 10
	fmt.Println(number)
	number /= 10
	fmt.Println(number)
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
















