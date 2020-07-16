package main

import "fmt"

//二分查找法

func binarySearch(arr *[10]int, leftIndex, rightIndex, findVal int) bool {

	if leftIndex > rightIndex {
		return false
	}
	//先找到中间的下标
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		//说明我们要查找的数，应该在left
		return binarySearch(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		return binarySearch(arr, middle+1, rightIndex, findVal)
	} else {
		return true
	}
	return false
}

//二分查找法基于循环的实现
func binaryLoop(arr *[10]int, val int) bool {
	leftIndex := 0
	rightIndex := len(arr) - 1
	for leftIndex <= rightIndex {
		midIndex := (leftIndex + rightIndex) / 2
		if (*arr)[midIndex] > val { //中间值大于要查找的值则要查找的值的范围在 leftIndex (midIndex - 1)
			rightIndex = midIndex - 1
		} else if (*arr)[midIndex] < val { //中间值小于要查找的值则要查找的值的范围在 (midIndex + 1) rightIndex
			leftIndex = midIndex + 1
		} else {
			return true
		}
	}
	return false
}
func main() {
	arr := [10]int{1, 11, 23, 55, 59, 123, 220, 223, 566, 1234}
	res := binarySearch(&arr, 0, len(arr)-1, 223)
	fmt.Println(res)
}
