package main                                     
import "fmt"
func mergesort(arr[]int)[]int{
	if len(arr)<=1{
		return arr
	}
	mid :=len(arr)/2
	left := mergesort(arr[:mid])
	right := mergesort(arr[mid:])
	return merge(left,right)
}
func merge(left,right []int)[]int{
	result :=make([]int,len(left)+len(right))
	i,j := 0,0
	for k := 0; k < len(result); k++ {
		if i >= len(left){
			result[k]=right[j]
			j++	
		}else if j >= len(right){
			result[k]=left[i]
			i++
		}else if left[i] < right[j]{
			result[k]=left[i]
			i++
		}else{
			result[k]=right[j]
			j++
		}
	}
	return result
}
func main() {
	arr := []int{23, 11, 41, 55, 22, 4, 7}
	fmt.Println(mergesort(arr))
}