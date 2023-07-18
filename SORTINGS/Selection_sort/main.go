package main

import "fmt"

func selectionsort(arr []int) []int {
	var temp int
	for i := 0; i < len(arr); i++ {
		min :=i
		for j := 0; j < len(arr); j++ {
			if arr[min] < arr[j]{
				arr[min],arr[j]=arr[j],arr[min]
			}
		}
		temp = arr[min]
		arr[min]=arr[i]
		arr[i]=temp
	}
	return arr
}
func main() {
	arr := []int{23, 11, 41, 55, 22, 4, 7}
	fmt.Println(selectionsort(arr))
}
