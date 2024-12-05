package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func compress(chars []byte) []byte {
	s := strings.Builder{}
	n := 0
	for _, v := range chars {
		if !strings.Contains(s.String(), string(v)) {
			n = strings.Count(string(chars), string(v))
			s.WriteByte(v)
			if n > 1 {
				s.WriteString(strconv.Itoa(n))
			} else {
				s.WriteString("")
			}
		}
	}
	chars = []byte(s.String())
	return chars
}

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
func (l *ListNode) Length() int {
	n := 0
	cur := l.Head
	if cur == nil {
		return 0
	}
	for cur != nil {
		cur = cur.Next
		n++
	}
	return n
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	cur1 := l1.Tail
	cur2 := l2.Tail
	res := &ListNode{}
	carry := 0
	for cur1 != nil || cur2 != nil || carry > 0 {
		val1, val2 := 0, 0

		if cur1 != nil {
			val1 = cur1.Val
			cur1 = cur1.Prev
		}
		if cur2 != nil {
			val2 = cur2.Val
			cur2 = cur2.Prev
		}
		sum := val1 + val2 + carry
		carry = sum / 10
		res.Append(sum % 10)
	}

	return res
}

func getSum(a int, b int) int {
	if a >= 0 && b >= 0 {
		for i := 0; i < b; i++ {
			a++
		}
		return a
	} else if a > 0 && b < 0 {
		for i := 0; i < a; i++ {
			b++
		}
		return b
	} else if a < 0 && b > 0 {
		for i := 0; i < b; i++ {
			a++
		}
		return a
	} else {
		for i := 0; i > b; i-- {
			a--
		}
		return a
	}
}
