package main

import "fmt"

// search element in array
func main() {
	arr := []int{1, 2, 3, 4}
	search_element := 2
	fmt.Println(linearSearch(arr, search_element))                //works for unsorted array too - time complexity of linear search is O(n)
	fmt.Println(binarySearch(arr, 0, len(arr)-1, search_element)) //works for sorted array - time complexity of binary search is O(log n)
}

func linearSearch(arr []int, search_element int) (bool, int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == search_element {
			return true, i
		}
	}
	return false, -1
}
func binarySearch(arr []int, l, r, search_element int) int {
	if r >= l {
		mid := l + (r-l)/2
		if arr[mid] == search_element {
			return mid
		}
		if arr[mid] > search_element {
			return binarySearch(arr, l, mid-1, search_element)
		}
		return binarySearch(arr, mid+1, r, search_element)
	}
	return -1
}
