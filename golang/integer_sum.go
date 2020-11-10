/*
 *  Determine if the sum of two integers is equal to the given value
 */

package main

import "fmt"

func sumExists(arr []int, sum int)  bool {
  set := make(map[int]int);
  for _, val := range arr {
    set[val] = 1;   
  }
  
  for _, val := range arr {
    diff := sum - val
    if _, ok := set[diff]; ok {
      return true
    }
  }
  return false
}

func main() {
  fmt.Printf("Sum Exists: %v\n", sumExists([]int{5, 7, 1, 2, 8, 4, 3}, 10));
  fmt.Printf("Sum Exists: %v\n", sumExists([]int{5, 7, 1, 2, 8, 4, 3}, 19));
  fmt.Printf("Sum Exists: %v\n", sumExists([]int{5, 7, 1, 2, 8, 4, 3}, 11));
}
