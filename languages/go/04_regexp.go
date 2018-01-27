package main

import "regexp"

func main() {

    // s := []string{"The thing for Jamis", "James is a great king", "The king is named James but he sin't great"}

    s := "testing whether we can perform a match on this string thing..."

    bool, err := regexp.MatchString("perform", s)


    println("result", bool)
    println("err", err) // err will be null.
}
