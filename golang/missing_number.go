/*
 - One missing number - n * (n - 1)/2 - (sum of elements in array)
 - Multiple missing numbers - 
    - create an array of n and set all values to 0, set values at index to 1 for elements in array
*/

package main

import "fmt"

func findOneMissingElement(arr []int) int {
    maxElement := len(arr) + 1
    sum := maxElement * (maxElement + 1) / 2
    arrSum := 0
    for _, val := range arr {
       arrSum += val
    }
    return sum - arrSum
}

/*
  args:
    * arr - input array
    * num - number of elements
*/
func findMultipleMissingElements(arr []int, num int) {
    fmt.Printf("Missing numbers: ")
    bits := make([]int, num + 1)
    for _, val := range arr {
        bits[val] = 1
    }
    for i, val := range bits {
       if val == 0 {
           fmt.Printf("%d, ", i)
       }
    }
    fmt.Println()
}

func main() {
   fmt.Printf("Missing number: %d\n", findOneMissingElement([]int{1, 3, 4, 5}))
   findMultipleMissingElements([]int{1, 3, 4, 5}, 5);
}
