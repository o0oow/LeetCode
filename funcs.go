package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func reverse(x int) int {
	str := ""
	cur := 0
	if x < 0 {
		str = "-"
		x = (-1) * x
	}
	for x != 0 {
		cur = x % 10
		str += strconv.Itoa(cur)
		x = x / 10
	}

	n, _ := strconv.Atoi(str)
	return range1(n)
}

func MyAtoi(s string) int {
	s1 := ""
	for i, v := range s {
		if unicode.IsDigit(rune(v)) || rune(v) == rune('-') || rune(v) == rune('+') {
			if string(v) == "-" && len(s1) > 0 ||
				string(v) == "+" && len(s1) > 0 {
				n, _ := strconv.Atoi(string(s1))
				return range1(n)
			}
			s1 += string(rune(v))
		} else if i != 0 && unicode.IsDigit(rune(s[i-1])) {
			n, _ := strconv.Atoi(string(s1))
			return range1(n)
		} else if unicode.IsSpace(rune(v)) && strings.Contains(s1, "-") ||
			unicode.IsSpace(rune(v)) && strings.Contains(s1, "+") {
			n, _ := strconv.Atoi(string(s1))
			return range1(n)

		} else if unicode.IsSpace(rune(v)) {
			continue
		} else {
			n, _ := strconv.Atoi(string(s1))
			return range1(n)
		}
	}
	n, _ := strconv.Atoi(string(s1))
	return range1(n)
}
func range1(n int) int {
	if n < -2147483648 {
		return -2147483648
	} else if n > 2147483647 {
		return 2147483647
	}
	return n
}

func addSpaces(s string, spaces []int) string {
	s1 := ""
	n := 0
	for _, indexval := range spaces {
		s1 += s[n:indexval] + " "
		n = indexval
	}
	s1 += s[n:]
	return s1
}

func addSpaces1(s string, spaces []int) string {
	var result strings.Builder
	prev := 0

	for _, index := range spaces {
		result.WriteString(s[prev:index])
		result.WriteByte(' ')
		prev = index
	}

	result.WriteString(s[prev:])
	return result.String()
}

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}
type ListNode struct {
	Head *Node
	Tail *Node
}

func (l *ListNode) Append(val int) {
	newEl := &Node{Val: val}
	if l.Head == nil {
		l.Head = newEl
		l.Tail = newEl
		return
	}
	l.Tail.Next = newEl
	newEl.Prev = l.Tail
	l.Tail = newEl

}
func (list *ListNode) Reverse() *ListNode {
	if list == nil {
		return list
	}
	res := &ListNode{}
	cur := list.Tail
	for cur != nil {
		res.Append(cur.Val)
		cur = cur.Prev
	}
	return res

}
func (l *ListNode) Print() {
	cur := l.Head
	if cur == nil {
		fmt.Print("nil")
	}
	for cur != nil {
		fmt.Printf("%v ", cur.Val)
		cur = cur.Next
	}
	fmt.Print("nil")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) int {
	if l1 == nil {
		//	return l2
	} else if l2 == nil {
		//	return l1
	}
	var res1, res2 strings.Builder
	//res := &ListNode{}
	cur1 := l1.Tail
	cur2 := l2.Tail
	for cur1 != nil {
		res1.WriteString(string(cur1.Val))
		cur1 = cur1.Prev
	}
	for cur2 != nil {
		res2.WriteString(string(cur2.Val))
		cur2 = cur2.Prev
	}
	r, _ := strconv.Atoi(res1.String())
	r2, _ := strconv.Atoi(res2.String())
	return r + r2
}
