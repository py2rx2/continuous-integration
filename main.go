package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	result := Add(5, 3)
	fmt.Println("Result:", result) 
}
