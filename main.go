package main

import "fmt"

func main() {
	// fmt.Println("Hello world")
	// fmt.Println(get(2))
	stu := Student{
		name: "Tom",
	}
	fmt.Println(stu.hello("Jack"))
}

func get(index int) int {
	arr := []int{2, 3, 4}
	return arr[index]
}

type Student struct {
	name string
}

func (stu *Student) hello(person string) string {
	return fmt.Sprintf("hello %s, I am %s.", person, stu.name)
}
