

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
func learnMultiple(x, y int) (sum, prod int) {
    return x + y, x * y // Return two values.
}

func main() {

    // var result = learnMultiple(5, 6);

    // Function returns two values.
    // quick declaration and assignment for function that returns 2 values
    sum, prod := learnMultiple(5, 6)


    fmt.Println("sum:", sum, "prod:", prod)
}
