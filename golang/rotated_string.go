/*
 * Rotated String:
 *          hello ---> ohell
 */

package main

import (
  "fmt"
  "strings"
)

func isRotated(original string, newString string) bool {
  if (len(original) != len(newString)) {
    return false
  }

  doubleString := original + original
  return strings.Contains(doubleString, newString)
}

func main() {
  fmt.Printf("Hello -- oHell: %v\n", isRotated("Hello", "oHell"))
  fmt.Printf("IndiaUSAEngland -- USAEnglandIndia: %v\n", isRotated("IndiaUSAEngland", "USAEnglandIndia"))
}
