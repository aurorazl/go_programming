package main

import "fmt"

func ReverseString(s1, s2 string) bool {
	if len(s1) != len(s2) || len(s1) == 1 {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if i+1 < len(s1) && s2[i] == s1[len(s1)-1] && s2[i+1] == s1[0] {
			return s2[i+1:] == s1[0:len(s1)-i-1] && s2[0:i+1] == s1[len(s1)-i-1:]
		}
	}
	return false
}

func main() {
	var s1 = "aba"
	var s2 = "baa"
	fmt.Print(ReverseString(s1, s2))
}
