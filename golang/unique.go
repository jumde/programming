/*
 * Return unique from two slices
 *
 */
package main

import "fmt"

func uniqueNames(a, b []string) []string {
	var result []string
	var nameMap map[string]int

	nameMap = make(map[string]int)

	for _, name := range a {
		if _, ok := nameMap[name]; !ok {
			nameMap[name] = 1
		}
	}

	for _, name := range b {
		if _, ok := nameMap[name]; !ok {
			nameMap[name] = 1
		}
	}

	for key := range nameMap {
		result = append(result, key)
	}

	return result
}

func main() {
	// should print Ava, Emma, Olivia, Sophia
	fmt.Println(uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"}))
}
