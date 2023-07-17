// Bubble sort
package main

import "fmt"

func bubblesort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
func main() {
	arr := []int{23, 11, 41, 55, 22, 4, 7}
	fmt.Println(bubblesort(arr))
}
