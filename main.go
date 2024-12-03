package main

import "fmt"

func main() {
	list := &ListNode{}
	list1 := &ListNode{}
	list.Append(2)
	list.Append(4)
	list.Append(3)
	list1.Append(5)
	list1.Append(6)
	list1.Append(4)
	l := addTwoNumbers(list, list1)

	fmt.Printf("%s \n", l)
	var i []int = []int{1, 2, 3}
	addSpaces("hdsvbhdfv", i)

}
