package main

import "fmt"

func main1() {
	var originalCount int = 10
	fmt.Println("I started with", originalCount, "apples.")

	var eatenCount int = 4
	fmt.Println("Some jake ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apples left.")
}
