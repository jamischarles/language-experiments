// simple experiments reading / writing json data in go https://gobyexample.com/json
package main

import "encoding/json"
import "fmt"

// import "os"

func main() {

	// First we’ll look at encoding basic data types to JSON strings.
	// Here are some examples for atomic values.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// And here are some for slices and maps, which encode to JSON
	// arrays and objects as you’d expect.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// Now let’s look at decoding JSON data into Go values. Here’s an
	// example for a generic data structure.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	// fmt.Println(string(byt))

	// We need to provide a variable where the JSON package can put
	// the decoded data. This map[string]interface{} will hold a map of
	// strings to arbitrary data types.
	var dat map[string]interface{}

	// Here’s the actual decoding, and a check for associated errors.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// In order to use the values in the decoded map, we’ll need to cast
	// them to their appropriate type. For example here we cast the value in
	// num to the expected float64 type.
	num := dat["num"].(float64)
	fmt.Println("num", num)

	// Accessing nested data requires a series of casts.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println("str1", str1)

	// https://gobyexample.com/range
	// range iterates over elements in a variety of data structures.
	// Let’s see how to use range with some of the data structures we’ve
	// already learned.

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// long custom json obj I need to itearate over
	byt2 := []byte(`{ "data": [ { "datapoints": [ [ 1.0, 1508107740000 ], [ 1.0, 1508146620000 ], [ 1.0, 1508160840000 ] ], "logViewLinks": [], "target": "Q3:Event Name-dcg13-p2pnodeweb-ERROR-UNCAUGHTEXCEPTION-COUNT" }, { "datapoints": [ [ 1.0, 1508084700000 ] ], "logViewLinks": [], "target": "Q3:Event Name-dcg12-p2pnodeweb-ERROR-UNCAUGHTEXCEPTION-COUNT" }, { "datapoints": [ [ 1.0, 1508124600000 ], [ 1.0, 1508167260000 ] ], "logViewLinks": [], "target": "Q3:Event Name-dcg11-p2pnodeweb-ERROR-UNCAUGHTEXCEPTION-COUNT" } ], "empty": false, "reachedMaxRecords": false }`)
	// fmt.Println(string(byt2))

	// map of arbitrary data types (we can prolly optimize this later...)
	var data map[string]interface{}

	// decode json and store in data map. handle err
	if err := json.Unmarshal(byt2, &data); err != nil {
		panic(err)
	}

	// loop over the data using range
	// access the props we want...
	arr := data["data"].([]interface{})

	// loop over the json set, and print each one...
	for _, num := range arr {
		// fetch the obj and print the prop we want...
		obj := num.(map[string]interface{})
		fmt.Println("target", obj["target"])
	}

}
