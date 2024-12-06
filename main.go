package main

import "fmt"

func main() {
	var char string = ""
	char1 := []byte(char)
	fmt.Println(compress(char1))
	fmt.Println(string(char1))
}
