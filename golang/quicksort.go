/*
 *  K largest element 
 *  
 *  Quick Sort -
 *    - Select the pivot element - rightmost element
 *    - Partition - place the pivot element in the correct position
 *    - Quick Sort the othe elements
 *    - Complexity - O(nlogn) 
 */

package main

import "fmt"

func partition(arr []int, low int, high int) int {
  pivot_index := high - 1
  pivot := arr[pivot_index]
  for i := low; i < high - 1 ; i++ {
    if (arr[i] < pivot) {
      arr[i], arr[pivot_index] = arr[pivot_index], arr[i]
      pivot_index++
    }
  }
  arr[pivot_index], arr[high - 1] = arr[high - 1], arr[pivot_index]
  return pivot_index
}

func kLargestElement(arr []int, low int, high int, k int) int {
  if (low < high) {
    pivot_index := partition(arr, low, high)

    fmt.Printf("Array: %v\n", arr)
    fmt.Printf("Pivot Index: %v\n", pivot_index)

    kLargestElement(arr, low, pivot_index - 1, k)
    kLargestElement(arr, pivot_index + 1, high, k)
  }
  return 0
}

func main() {
  arr := []int{7, 89, 11, 23, 46, 56}
  fmt.Printf("kth largest element: %v %v \n", kLargestElement(arr, 0, len(arr), 2), arr)
}
