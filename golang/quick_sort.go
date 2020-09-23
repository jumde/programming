/*
    Quick Sort:

    [8, 7, 6, 1, 0, 9, 2]
*/

package main

import "fmt"

func partition(arr []int, low int, high int) int {
   fmt.Printf("Array: %v\n", arr[low:high])
   pivot := arr[high - 1]
   i := low
   for j := low; j< high-1; j++ {
     if (arr[j] < pivot) {
       arr[j], arr[i] = arr[i], arr[j]
       i++
     }
   }
   arr[high - 1], arr[i] = arr[i], arr[high - 1]
   return i
}

func quickSort(array []int, low int, high int) {
   if (low < high) {
     pivot := partition (array, low, high)
     fmt.Printf("Pivot: %v\n", pivot)
     quickSort(array, low, pivot)
     quickSort(array, pivot+1, high)
   }
}

func main() {
  arr := []int{8, 7, 6, 1, 0, 9, 2}
  quickSort(arr, 0, len(arr))
  fmt.Printf("Sorted: %v \n", arr)
}
