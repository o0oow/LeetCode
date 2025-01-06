package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func twoSum(numbers []int, target int) []int {
	var b []int
	i, j := 0, len(numbers)-1
	for i < len(numbers) {
		if (numbers[i] + numbers[j]) == target {
			b = append(b, i+1, j+1)
			return b
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}

	}
	return b
}
func reverseStr(s string, k int) string {
	s1 := []byte(s)
	var i, j int
	x := 0
	n := len(s1) / (k * 2)
	for z := 0; z < n; z++ {
		i = z * 2 * k
		j = z*2*k + (k - 1)
		for x < k/2 {
			s1[i], s1[j] = s1[j], s1[i]
			i++
			j--
			x++
		}
		x = 0

	}
	return string(s1)
}
func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < len(s)/2 {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
func intersection(nums1 []int, nums2 []int) []int {
	c1 := make(map[int]int)
	c2 := make(map[int]int)
	n := make([]int, 0)
	for _, v := range nums1 {
		c1[v]++
	}
	for _, v := range nums2 {
		c2[v]++
	}
	if len(c1) > len(c2) {
		for v, _ := range c1 {
			if c1[v] >= 1 && c2[v] >= 1 {
				n = append(n, v)
			}
		}
	} else {
		for v, _ := range c2 {
			if c1[v] >= 1 && c2[v] >= 1 {
				n = append(n, v)
			}
		}
	}
	return n
}
func uncommonFromSentences(s1 string, s2 string) []string {
	c1 := make(map[string]int)
	c2 := make(map[string]int)
	s11 := strings.Split(s1, " ")
	s22 := strings.Split(s2, " ")
	s33 := make([]string, 0)
	for _, v := range s11 {
		c1[v]++
	}
	for _, v := range s22 {
		c2[v]++
	}
	for v, _ := range c1 {
		if (c1[v] == 1 && c2[v] == 0) || (c1[v] == 0 && c2[v] == 1) {
			s33 = append(s33, v)
		}
	}
	for v, _ := range c2 {
		if (c1[v] == 1 && c2[v] == 0) || (c1[v] == 0 && c2[v] == 1) {
			s33 = append(s33, v)
		}
	}
	return s33
}
func countWords(words1 []string, words2 []string) int {
	c1 := make(map[string]int)
	c2 := make(map[string]int)

	count := 0
	for _, v := range words1 {
		c1[v]++
	}
	for _, v := range words2 {
		c2[v]++
	}
	if len(words1) > len(words2) {
		for v, _ := range c1 {
			if c1[v] == 1 && c2[v] == 1 {
				count++
			}
		}
	} else {
		for v, _ := range c2 {
			if c1[v] == 1 && c2[v] == 1 {
				count++
			}
		}
	}
	return count

}

func kthDistinct(arr []string, k int) string {
	c := make(map[string]int)
	for _, v := range arr {
		c[v]++
	}
	for _, v := range arr {
		if c[v] == 1 {
			k--
		}
		if k == 0 {
			return v
		}
	}
	return ""
}
func maximumNumberOfStringPairs(words []string) int {
	count := 0
	for i := 0; i < len(words)-1; i++ {
		j := i + 1
		for j < len(words) {
			if words[i] == Reverse(words[j]) {
				count++
				break
			} else {
				j++
			}
		}
	}
	return count
}
func Reverse(s string) string {
	var s1 string
	for i := len(s) - 1; i >= 0; i-- {
		s1 += string(s[i])
	}
	return s1
}
func sortPeople(names []string, heights []int) []string {
	m := make(map[int]string)
	for i := 0; i < len(names); i++ {
		m[heights[i]] = names[i]
	}
	heights = quickSortDescending(heights)
	for i := 0; i < len(names); i++ {
		names[i] = m[heights[i]]
	}
	return names
}
func quickSortDescending(arr []int) []int {
	// Базовый случай: массив длиной 0 или 1 уже отсортирован
	if len(arr) <= 1 {
		return arr
	}

	// Опорный элемент
	pivot := arr[len(arr)/2]

	// Три части: больше, равны, меньше
	greater := []int{}
	equal := []int{}
	less := []int{}

	for _, num := range arr {
		if num > pivot {
			greater = append(greater, num)
		} else if num == pivot {
			equal = append(equal, num)
		} else {
			less = append(less, num)
		}
	}

	// Рекурсивно сортируем части и объединяем
	return append(append(quickSortDescending(greater), equal...), quickSortDescending(less)...)
}

func finalPrices(prices []int) []int {

	for i := 0; i < len(prices)-1; i++ {
		j := i + 1
		for j <= len(prices)-1 {
			if prices[j] <= prices[i] {
				prices[i] -= prices[j]
				break
			} else {
				j++
			}
		}
	}
	return prices
}

func getFinalState(nums []int, k int, multiplier int) []int {
	j := 0
	for i := 0; i < k; i++ {
		j = findMinIndex(nums)
		nums[j] *= multiplier
	}
	return nums
}
func isSubsequence(s string, t string) bool {
	if len(s) <= 1 {
		return strings.Contains(t, s)
	}
	count := 1
	s1 := make([]int, 0)
	if len(t) > len(s) {
		s1 = append(s1, strings.Index(t, string(s[0])))
		for i := 1; i < len(s); i++ {

			s1 = append(s1, strings.Index(t, string(s[i])))
		}
		for i := 0; i < len(s1)-1; i++ {
			if s1[i] < s1[i+1] {
				count++
			}
		}
	} else {
		return false
	}
	return count == len(s)
}

func binarySearch(arr []string, target string) bool {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2 // Избегаем переполнения при больших значениях
		if arr[mid] == target {
			return true // Элемент найден
		} else if arr[mid] < target {
			left = mid + 1 // Сужаем поиск к правой половине
		} else {
			right = mid - 1 // Сужаем поиск к левой половине
		}
	}

	return false // Элемент не найден
}
func minimumSize(s []int, maxOperations int) int {
	m := 0
	z := 0
	for z < maxOperations {
		//m = findMinIndex(Index(s))
		if s[m] > 1 {
			part1 := s[m] / 2
			part2 := s[m]/2 + s[m]%2
			if part1 > 0 && part2 > 0 && part1+part2 == s[m] {
				s = append(s[:m], append([]int{part1, part2}, s[m+1:]...)...)
			}
		}
		z++
	}
	min1 := s[0]
	for _, v := range s {
		if min1 > v {
			min1 = v
		}
	}

	return min1
}
func findMinIndex(arr []int) int {
	maxIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false // Найден делитель, значит число не простое
		}
	}
	return true
}
func maxCount1(banned []int, n int, maxSum int) int {

	return 1
}

//	func canChange(start string, target string) bool {
//		i := 0
//		j := len(start) - 1
//		s := strings.Builder{}
//		for k, v := range start {
//			if string(start[k]) == "L" {
//				s.WriteString(string(v) + strings.Repeat("_", k-1))
//			}
//
//		}
//		return true
//	}
func compress(chars []byte) int {
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
	return len(chars)
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

func maxCount(banned []int, n int, maxSum int) int {
	BanMap := make(map[int]bool)
	for _, val := range banned {
		BanMap[val] = true
	}
	var sum, count int
	for i := 1; i <= n; i++ {
		if BanMap[i] {
			continue
		}
		sum += i
		if sum <= maxSum {
			count++
		}
	}
	return count
}
