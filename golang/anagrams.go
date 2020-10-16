/*
  Check for anagrams
*/

package main

import "fmt"

func isAnagram(str1 string, str2 string)  bool {
    var charSet = make(map[string]int)
    for _, c := range str1 {
        charSet[string(c)] ++
    }
    for _, c := range str2 {
        charSet[string(c)] --
    }
    for _, val := range charSet {
        if val != 0 {
            return false
        }
    }
    return true
}

func main() {
    fmt.Printf("fired/fried: %v\n", isAnagram("fired", "fried"));
    fmt.Printf("listen/silent: %v\n", isAnagram("listen", "silent"));
    fmt.Printf("yolo/bolo: %v\n", isAnagram("yolo", "bolo"));
}
