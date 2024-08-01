// The main package will ensure that the program will be compiled to an executable file
package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// The function with the name 'main' is the entry point of the application
// There must be only 1 'main' function

// Defining functions
func sayGreeting(n string) {
	fmt.Printf("Good Morning %v \n", n)
}

// Another variation
func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func function() (string, string) {
	val := 10
	if val > 10 {
		return "String 1", "String 2"
	} else {
		return "String 1", "_"
	}
}

// Function returning a value, note that the return type has to be specified
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	fmt.Println("Hello, there")

	//-----------------------VARIABLES-------------------------------
	// Remember that go is a statically typed language

	// strings
	var nameOne string = "hermon" // Double quotes for strings
	var nameTwo = "Fenu"          // Go can also infer variable types
	var nameThree string          // Variable declaration

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "Messi"
	nameFour := "mom" // same as [var nameFour string = "mom"]
	// You can use this shorthand outside functions

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// ints
	var age1 int = 20
	var age2 = 30
	age3 := 40

	var num1 int8 = 50   // specify the max number of bits, here [-128 -> 127 allowed]
	var num2 uint8 = 255 // unsigned int, here [0 -> 255], negatives not allowed

	fmt.Println(age1, age2, age3, num1, num2)

	// floats
	var scoreOne float32 = -198.42 // bit size MUST be specified
	var scoreTwo float64 = 3343.99 // better precision comaperd with float32
	scoreThree := 1.5              // float64 is the default if bit num not specified

	fmt.Println(scoreOne, scoreTwo, scoreThree)

	//-----------------------The FMT Package------------------------

	// Print
	fmt.Print("hello,")   // Not newline added at the end of 'Print', unlike 'Println'
	fmt.Print("world \n") // You would have to manually add '\n'

	// Printf (formated strings), %_ -> format specifiers
	var name string = "Hermon"
	var age int = 21

	fmt.Printf("my age is %v and my name is %v \n", age, name) // Notice that the order matters
	fmt.Printf("My name is %q \n", name)                       // Adding quotes
	fmt.Printf("age is of type %T \n", age)                    // Adding variable type
	fmt.Printf("You scored %0.2f points ! \n", 422.321)        // Format floats and limit the num of decimal places displayed

	// Sprintf (save formatted strings)
	var myStr = fmt.Sprintf("%v is my favorite number!", 10)
	fmt.Println(myStr)

	//-----------------------ARRAYS---------------------------

	// Defining arrays
	var ages1 [3]int = [3]int{20, 23, 84} // Notice length and type of arrays are static
	var ages2 = [3]int{20, 23, 84}        // Shorthand method
	ages3 := [3]int{20, 23, 84}           // Even more shorter

	ages1[0] = 45 // You can change values of elements of arrays

	fmt.Println(ages1, ages2, ages3, len(ages1))

	// Slices (use arrays under the hood)
	var scores = []int{123, 452}

	scores[1] = 100             // Changing values in slices
	scores = append(scores, 73) // Notice that the append() fun doesn't automatically append values

	fmt.Println(scores, len(scores))

	// slice ranges
	var arr = [5]int{0, 1, 2, 3, 4}

	rangeOne := arr[1:3] // slicing properties are just like python
	rangeTwo := arr[1:]
	rangeThree := arr[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	//-----------------------The Standard Library--------------------

	// 'string' package
	greeting := "Hello there people!"

	fmt.Println(strings.Contains(greeting, "Hello"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, " th")) // if not found, returns -1
	fmt.Println(strings.Split(greeting, " "))   // returns a slice splitted by the given character

	// 'sort' package
	nums := []int{5, 1, 3, 6, 2, 1, 5, 7, 3, 2} // Defining a slice

	sort.Ints(nums)   // this sorting operation doesn't work on arrays
	fmt.Println(nums) // and unlike the prev 'string' methods. 'sort' will alter the original data

	index1 := sort.SearchInts(nums, 6)   // returns the index
	index2 := sort.SearchInts(nums, 100) // returns the length if value not found
	fmt.Println(index1, index2)

	names := []string{"Hermon", "Kidus", "Bisrat", "Mele", "Tize"}
	sort.Strings(names)

	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "Tize"))

	//-----------------------Loops------------------------
	// Check out the different variations of loops, (all 'for' loops, apparently no 'while' loops)

	// Just like a 'while' loop
	x := 0
	for x < 5 {
		fmt.Println("Value of x is:", x)
		x++
	}

	for z := 0; z < 5; z++ {
		fmt.Println("Value of z is:", z)
	}

	siblings := []string{"Hermon", "Fenu", "Beamlak"}
	for i := 0; i < len(siblings); i++ {
		fmt.Println("Current Sibling-", siblings[i])
	}

	for index, value := range siblings {
		fmt.Printf("Position %v, Sibling %v \n", index, value)
	}

	// if you want to take just the value
	for _, value := range siblings {
		fmt.Printf("Sibling %v \n", value)
	}

	//-----------------------Conditionals------------------------
	// It's just like you know it!!
	current := 50
	if current < 30 {
		fmt.Println("Age is less that 30")
	} else {
		fmt.Println("Age is greater that 30")
	}

	//------------------------------
	sayHello("Hermon")

	//------------maps------------------
	menu := map[string]float64{
		"soup": 4.99,
		"pie":  7.99,
	}

	fmt.Println(menu, menu["pie"])

	// Looping through maps
	for key, value := range menu {
		fmt.Println(key, value)
	}

	//Pointers
	curr_name := "Hermon"

	fmt.Println("mem add: ", &curr_name)
	pointer := &curr_name

	*pointer = "Fenu"
	fmt.Println(curr_name)

	//STRUCT
	myBill := newBill("Hermon's bill")
	fmt.Println(myBill)
}
