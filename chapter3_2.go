package main

import "fmt"

func manyReturns() (int, bool, string) {
	return 1, true, "hello"
}
func main2() {
	myInt, myBool, myString := manyReturns()
	fmt.Println(myInt, myBool, myString)
}
