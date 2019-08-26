package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//s := "1568141291110137"
	//l := 10

	//fmt.Println(MysteryRange(s, 10))

	fmt.Println(MysteryRange("1721532418565922162558663126649136347436733301144143236653738464135820194215516155541239452852623450572927602348104049", 60))
	//fmt.Println(MysteryRange("13161820142119101112917232215", 15))

	//s = "1721532418565922162558663126649136347436733301144143236653738464135820194215516155541239452852623450572927602348104049"
	//l = 60

	//fmt.Println(MysteryRange(s, l))

	//r := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//r := []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	//fmt.Printf("belongsToRange returns : %v\n", belongsToRange(s, r))

	/*fmt.Println(s)
	fmt.Println(yankOut(s, "15"))

	tmp := yankOut("15", "15")
	if tmp == "" {
		fmt.Println("done")
	}*/

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

func MysteryRange(s string, l int) [2]int {
	fmt.Println("MysteryRange() called")
	fmt.Printf("s=%s\n", s)
	fmt.Printf("l=%d\n", l)
	ret := [2]int{0, 0}
	for x := 1; x < 100+l-l; x++ { // loop thru all possible ranges that are of length l

		r := make([]int, l)

		cnt := 0
		for y := x; y < x+l; y++ {
			//fmt.Printf("%d ", y)
			r[cnt] = y
			cnt++
		}

		// now we have a built range, we want to see if the string s belongs to that range
		if belongsToRange(s, r) {
			fmt.Println("YES!")
			fmt.Println(r)
			ret[0] = r[0]
			ret[1] = r[len(r)-1]
			return ret
		}
		fmt.Println()
	}
	return ret
}

func belongsToRange(s string, r []int) bool {
	fmt.Printf("belongsToRange() s=%v r=%v\n", s, r)

	tmp := s
	rx := make([]bool, len(r))

	for x := len(r) - 1; x >= 0; x-- { // start yanking out, start from the largest number
		ok := false

		yo := strconv.Itoa(r[x])

		tmp, ok = yankOut(yo, tmp)
		if !ok {
			return false
		}

		rx[x] = true
		fmt.Println(tmp)
	}

	// if tmp is empty and all the rx[] items are true, return true

	fmt.Println("belongsToRange, tmp should be empty")
	fmt.Printf("tmp=%v\n", tmp)
	fmt.Println(rx)

	if tmp != "" {
		return false
	}

	for x := 0; x < len(rx); x++ {
		if rx[x] == false {
			return false
		}
	}

	return true
}

// yankOut removes a substring (yo) from within a string (s) and returns the result
func yankOut(yo string, s string) (string, bool) {
	i := strings.Index(s, yo)
	if i < 0 {
		return s, false
	}

	ret := s[:i] + s[i+len(yo):]
	return ret, true
}
