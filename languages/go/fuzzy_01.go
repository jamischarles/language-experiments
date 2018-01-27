// First attempt at fuzzy finder in Go...

// MVP1:
// We need to start with an input list
// then based on keyboard input, search against that list
// then show the results of that search...

// MVP2:
// make it work with command line input (like folders)
// - have it give output
// - add timing

// MVP3:
// - try to add concurrency to make it a whole lot faster...


// if there's no input args, maybe just read folders from the current directory as input?
// os.readdirrnames

// does that mean this is the 'main' module?
package main

import "fmt"
import "regexp"
import "bufio"
import "os"

func main() {
    s := []string{"The thing for Jamis", "James is a great king", "The king is named James but he sin't great"}

    // pipe input
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	// fmt.Println(text)

    slice := []string{}

    // iterate over an array with a range
    for _, v := range s {
        // is there a match?
        // if yes, add it to the slice
        // TODO: figure out how to watch stdin, or how to watch for typing event...

	fmt.Println(text)

    // WHY won't the match work from input? Does it have to be cast first?
        isMatch, _ := regexp.MatchString(text, v)

        if isMatch ==true {
            // fmt.Printf("INSIDE")
            slice = append(slice, v)
        }

        // fmt.Printf("isMatch", isMatch)

        // fmt.Printf("err\n", err)

        // fmt.Printf("2**%d = %d\n", i, v)

        // if there's a match add it to the result set...
    }

    // TODO: we need to have a ontype handler

    // then, feed the slice back into the algorithm...
    // stdin is only used the first time around, then we control what the input/output is...

    // iterate through it...

    // regexp against it

    // fmt.Println("hello world")

    fmt.Printf("%q\n", slice)


go func() {
    consolescanner := bufio.NewScanner(os.Stdin)

    // by default, bufio.Scanner scans newline-separated lines
    for consolescanner.Scan() {
        input := consolescanner.Text()
        fmt.Println(input)
    }

    // check once at the end to see if any errors
    // were encountered (the Scan() method will
    // return false as soon as an error is encountered)
    if err := consolescanner.Err(); err != nil {
         fmt.Println(err)
         os.Exit(1)
    }
}()

}


// have a list of values...
