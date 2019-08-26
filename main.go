package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "1568141291110137"

	fmt.Println(s)
	fmt.Println(yankOut(s, "15"))

	//l := 10

	// for x := 1; x < 100-l; x++ {	// loop thru all possible ranges that are of length l
	// 	r := []int{}
	// 	// todo : run the make here to allocate what we need

	// 	for y := x; y < x+l; y++ {
	// 		fmt.Printf("%d ", y)
	// 		r = append(r, y)
	// 	}

	// 	// now we have a built range, we want to see if the string s belongs to that range
	// 	if belongsToRange(s, r) {
	// 		fmt.Println("YES!")
	// 	}
	// 	fmt.Println()
	// }

}

func belongsToRange(s string, r []int) bool {
	ret := false

	//s1 := s
	//idxs := []int{}
	for x := 0; x < len(r); x++ {
		//idxs[x] := strings.Index(s1,

	}

	return ret
}

// yankOut removes a substring (yo) from a string s and returns the result
func yankOut(s string, yo string) string {
	fmt.Printf("yanking out %v\n", yo)
	i := strings.Index(s, yo)
	if i < 0 {
		return s
	}

	ret := s[:i] + s[i+len(yo):]
	return ret
}
