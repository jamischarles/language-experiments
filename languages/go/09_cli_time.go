package main

import "fmt"
import "time"
import "os"
import "strconv"

func main() {
	arg := os.Args[1]
	i, _ := strconv.ParseInt(arg, 10, 64)
	t := time.Unix(0, i)
	fmt.Println("Unix time: "+arg+" equals", t.Format(time.RFC822Z))
	// fmt.Println("Unix time: "+arg+"=", t.Format(time.RFC822Z))
}
