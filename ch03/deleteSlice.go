package main

import "fmt"

func main() {
	aSlice := []int{1, 2, 4, 5, 6, 7}
	fmt.Println(aSlice)

	// Delete element at index i
	i := 3
	if i > len(aSlice) {
		fmt.Println("Cannot delete element", i)
		return
	}
	aSlice = append(aSlice[:i], aSlice[i+1:]...)
	fmt.Println(aSlice)

	// Delete element at index i
	i = 0
	if i > len(aSlice) {
		fmt.Println("Cannot delete element", i)
		return
	}
	// Replace element at index i with last element
	copy(aSlice[i:], aSlice[i+1:])
	// Remove last element
	aSlice = aSlice[:len(aSlice)-1]
	fmt.Println(aSlice)
}
