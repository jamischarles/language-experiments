package main

import "fmt"

/* <- multiline comment
Functions can have parameters and (multiple!) return values.
Here `x`, `y` are the arguments and `sum`, `prod` is the signature (what's returned).
Note that `x` and `sum` receive the type `int`.
*/
// func learnMultiple(x, y int) (sum, prod int) {
//     return x + y, x * y // Return two values.
// }
// func learnMultiple(x, y int) (sum, prod int) {
//     return x + y, x * y // Return two values.
// }

func main() {
	// declare var before using it
	var x int

	x = 4

	fmt.Println("x", x)

	// "Short" declarations use := to infer the type, declare, and assign.
	y := 5
	fmt.Println("y", y)

	// arrays (slices)

	s := []string{"The thing for Jamis", "James is a great king", "The king is named James but he sin't great"}
	// var result = learnMultiple(5, 6);

	// Function returns two values.
	// quick declaration and assignment for function that returns 2 values
	// sum, prod := learnMultiple(5, 6)

	// fmt.Println("%q\n", s)
	fmt.Printf("%q\n", s)

	// test := s[0:4]
	// print a slice of the array (starting at index 0, and ending at index 1
	fmt.Printf("%q\n", s[0:1])

	// append
	s = append(s, "The newes Jamison")
	fmt.Printf("%q\n", s)

	// iterate over an array with a range
	for i, v := range s {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// Map like a JS object
	m := map[string]int{"three": 3, "four": 4}
	map2 := map[string]string{"three": "testing3", "four": "testing4"}

	fmt.Printf("%q\n", m)
	fmt.Printf("%q\n", map2)

	// iterate over a map
	for key, val := range map2 {
		fmt.Printf("2**%d = %d\n", key, val)
	}
}
