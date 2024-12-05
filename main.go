package main

import "fmt"

func main() {
	var char string = "aabbccc"
	fmt.Println(string(compress([]byte(char))))

}
