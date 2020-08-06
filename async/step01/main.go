package main

import "fmt"

func getIds() (ids []int) {
	ids = append(ids, 2)
	ids = append(ids, 3)
	ids = append(ids, 4)
	// "ids" is optional since the return value is named
	// This is known as a "bare return"
	return // ids
}

func main() {
	ids := getIds()
	fmt.Println(ids)
}
