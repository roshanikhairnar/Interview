package main

import (
	"fmt"
	"sync"
)

type stack []string

// reverse string
func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	s := "roshanik"
	fmt.Println("Reverse the string", s)
	go reverse1(s, &wg) //usign rune and two pointers
	go reverse2(s, &wg) //using byte and two pointers
	go reverse3(s, &wg) //using stack
	wg.Wait()

}
func reverse1(s string, wg *sync.WaitGroup) {
	// strings are immutable in Go.
	// Therefore, we must first convert the string to a mutable array of runes ([]rune),

	s_rune := []rune(s)
	left := 0
	right := len(s) - 1
	for left < right {
		s_rune[left], s_rune[right] = s_rune[right], s_rune[left]
		left = left + 1
		right = right - 1
	}
	fmt.Println("reverse1", string(s_rune))
	// Time Complexity: O(N)  Only one traversal to rune array.
	// Auxiliary Space: O(N) Extra for reverse string.
	wg.Done()
}
func reverse2(s string, wg *sync.WaitGroup) {
	// strings are immutable in Go.
	// Therefore, we can convert to byte array and then swap,
	byteString := []byte(s)
	left := 0
	right := len(s) - 1
	for left < right {
		byteString[left], byteString[right] = byteString[right], byteString[left]
		left = left + 1
		right = right - 1
	}
	fmt.Println("reverse2", string(byteString))
	// Time Complexity: O(N)  Only one traversal to byte array.
	// Auxiliary Space: O(N) Extra for reverse string.
	wg.Done()
}

// function to push the value to stack
func (s stack) Push(v string) stack {
	return append(s, v)
}

// function to pop the value from stack
func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

// function to get size of stack
func (s stack) Size() int {
	l := len(s)
	return l
}
func reverse3(s string, wg *sync.WaitGroup) {
	//using stack
	s_stack := make(stack, 0)
	for i := 0; i < len(s); i++ {
		s_stack = s_stack.Push(string(s[i])) //pushing char by char to a stack
	}
	var s_reverse string
	ch := ""
	for s_stack.Size() != 0 {
		s_stack, ch = s_stack.Pop() //pop and append to the new string
		s_reverse = s_reverse + ch
	}
	fmt.Println("reverse3", s_reverse)
	// Time Complexity: O(N)  Only one traversal to push and pop so O(n)+O(n)==O(n).
	// Auxiliary Space: O(N) Extra for Stack.
	wg.Done()
}
