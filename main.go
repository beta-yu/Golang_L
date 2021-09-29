package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	fmt.Println(get(2))
}

func get(index int) int {
	arr := []int{2, 3, 4}
	return arr[index]
}
