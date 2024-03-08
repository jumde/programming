package main

import (
	"fmt"
	"math"
)

func PadLeft(str string, length int) string {
	for len(str) < length {
	   str = "0" + str
	}
	return str
 }

func Reverse(s string) (result string) {
	for _,v := range s {
		result = string(v) + result
	}
	return result
}

func addBinaryStrings(a string, b string) string {
	lenA := len(a)
	lenB := len(b)

	res := ""
	carry := 0

	maxLen := int(math.Max(float64(lenA), float64(lenB)))
	padA := PadLeft(a, maxLen);
	padB := PadLeft(b, maxLen);

	for i:=maxLen-1; i>=0; i-- { 
		r := carry
		if (padA[i] == '1') {	
			r += 1 
		}
		if (padB[i] == '1') { 
			r += 1
		}
		if (r%2 == 1) {
			res = res + "1"
		} else {
			res = res + "0"
		}

		if (r >= 2) {
			carry = 1
		}
	}
	return Reverse(res)
}

func main() {
	a := 1001
	b := 10
	fmt.Printf("Sum of %v and %v is %v\n", a, b, addBinaryStrings("1001", "10"));
}