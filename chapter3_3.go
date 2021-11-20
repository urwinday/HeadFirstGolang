package main

import (
	"fmt"
)

func main() {
	//amount := 6
	//double(&amount)
	//fmt.Println(amount)

	name := "Stefan"
	//pName := &name

	//fmt.Println(name)
	//fmt.Println(pName)
	//fmt.Println(*pName)

	// Размименование
	//*pName = "Степан"
	//fmt.Println(name)

	setName(&name)
	fmt.Println(name)

	var emptyPointer *string
	fmt.Println(emptyPointer)

}

// Размименование
func setName(name *string) {
	*name = "Степан"
}

func double(number *int) {
	*number *= 2
}

func funcType() {
	//var myInt int
	//fmt.Println(reflect.TypeOf(&myInt))
	//var myFloat float64
	//fmt.Println(reflect.TypeOf(&myFloat))
	//var myBool bool
	//fmt.Println(reflect.TypeOf(&myBool))
	//
	////var myInt int
	//var myIntPointer *int
	//myIntPointer = &myInt
	//fmt.Println(myIntPointer)
	////var myFloat float64
	//var myFloatPointer *float64
	//myFloatPointer = &myFloat
	//fmt.Println(myFloatPointer)
}
