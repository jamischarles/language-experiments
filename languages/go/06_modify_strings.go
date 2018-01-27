// Read a file with a bunch of strings in it
// Modify each line in some  way...
// read input, make it upper case and save to output
// or just output it...

// https://gobyexample.com/reading-files

// MVP2:
// Ideally we'd do something like l2met where we take unstructured logs and count the metrics
// and turn them into structured counts and metrics
// this is what mtail does... We should write a simple version of mtail

package main

// import "regexp"
import "io/ioutil"
import "strings"

// import "fmt"

func main() {

	data, err := ioutil.ReadFile("./06_input.txt") // read the entire file into memory

	println("data", data) // Q: What does this output? [237/749]0xc420082000
	// println("data", strings.toUpper(string(data))) // cast as a string...
	// upper case the whole file...
	println("data", strings.ToUpper(string(data)))
	println("err", err) // err will be null.

	//
	d1 := []byte(data)                                    // convers  the data into bytes?
	err2 := ioutil.WriteFile("./06_output.txt", d1, 0644) // write the bytes to the output file

	println("err2", err2) // any errros?

}
