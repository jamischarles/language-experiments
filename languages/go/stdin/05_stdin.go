// practice taking and processing stdin
package main

import "os"
// import "byte"

import "fmt"

import "bufio"

// import "io"

func main() {



	// read one line in... you can either pipe something to it, or leave stdin blank, and then we can get it from the cmd line
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)


	// go through multiple lines...
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	println("result", os.Args[1])
}
