/*
  https://leetcode.com/explore/interview/card/top-interview-questions-easy/97/dynamic-programming/576/
*/

package main

import "fmt"

func rob(nums []int) int {
  exclude := 0
  
  if (len(nums) == 0) {
    return exclude;
  }

  include := nums[0]

  for i := 1; i < len(nums); i++ {
    // total including the current number
    sum := exclude + nums[i]

    // total excluding the current number
    exclude = include

    // update include if the sum including the 
    // current number is greater than the previous
    // include
    if (include < sum) {
      include = sum
    }
  }
  if (exclude < include) {
    exclude = include
  }
  return exclude
}

func main() {
   nums := []int{1,2,3,1}
   fmt.Printf("Rob: %v\n", rob(nums));

   nums = []int{2,7,9,3,1}
   fmt.Printf("Rob: %v\n", rob(nums));
   
   nums = []int{2,1,1,2}
   fmt.Printf("Rob: %v\n", rob(nums));

   nums = []int{1}
   fmt.Printf("Rob: %v\n", rob(nums));

   nums = []int{2, 1}
   fmt.Printf("Rob: %v\n", rob(nums));

   nums = []int{4,1,2,7,5,3,1}
   fmt.Printf("Rob: %v\n", rob(nums));
   
   nums = []int{}
   fmt.Printf("Rob: %v\n", rob(nums));
}
