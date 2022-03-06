package main

import "unicode/utf8"

import "fmt"

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	const (
		j = 1
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	const (
		FlagUp           = 1 << (10 * iota) // is up
		FlagBroadcast                       // supports broadcast access capability
		FlagLoopback                        // is a loopback interface
		FlagPointToPoint                    // belongs to a point-to-point link
		FlagMulticast                       // supports multicast access capability
	)
	var f1 int = 212.2
	fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, f1)
}
