package main

import "fmt"
func quicksort(arr[]int)[]int{
   if len(arr)<=1{
	return arr
   }
   pivot := arr[len(arr)-1]
   left ,right:= []int{},[]int{}
   for i := 0; i < len(arr)-1; i++ {
	 if arr[i]< pivot{
		left =append(left, arr[i])
	 }else{
		right =append(right, arr[i])
	 }
   }
   sortedleft :=quicksort(left)
   sortedright := quicksort(right)
   return append(append(sortedleft,pivot),sortedright...)
}
 
func main() {
	arr := []int{14, 6, 11, 7, 22, 18, 3}
	fmt.Println(quicksort(arr))
}
