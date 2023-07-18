// INSERTON SORT
package main

import "fmt"

func insertionsort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return arr
}
func main() {
	arr := []int{23, 11, 41, 55, 22, 4, 7}
	fmt.Println(insertionsort(arr))
}
