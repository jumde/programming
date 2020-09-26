/*
   Palindrome ignoring punctuation and casing
*/
package main

import "fmt"
import "unicode"

func isAlphaNumeric(c rune) bool {
  return unicode.IsLetter(c) || unicode.IsNumber(c)
}

func palindrome(input string) bool {
  inputRune := []rune(input)
  j := len(inputRune) - 1
  for i:= 0; i < j; i++ {
    for !isAlphaNumeric(inputRune[i]) {
      i++
    }
    
    for !isAlphaNumeric(inputRune[j]) {
      j--
    }

    //fmt.Printf("i: %v, ", unicode.ToLower(inputRune[i]))
    //fmt.Printf("j: %v\n", unicode.ToLower(inputRune[j]))

    if i > j ||  (unicode.ToLower(inputRune[i]) != unicode.ToLower(inputRune[j])) {
      return false;
    }
    j--
  }
  return true
}

func main() {
  fmt.Printf("Palindrome: (I) - %v\n", palindrome("I"))
  fmt.Printf("Palindrome: (I') - %v\n", palindrome("I'"))
  fmt.Printf("Palindrome: (I'm) - %v\n", palindrome("I'm"))
  fmt.Printf("Palindrome: (I'm a ami) - %v\n", palindrome("I'm a ami"))
  fmt.Printf("Palindrome: (Eva, can I see bees in a cave?) - %v\n", palindrome("Eva, can I see bees in a cave?"));
  fmt.Printf("Palindrome: (Was it a rat I saw?) - %v\n", palindrome("Was it a rat I saw?"));
}
