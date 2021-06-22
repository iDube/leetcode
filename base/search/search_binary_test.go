package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var lookingFor int = 6
	var sortedArray []int = []int{1, 3, 4, 6, 7, 9, 10, 11, 13}
	index := BinarySearch(sortedArray, lookingFor)
	fmt.Println("the array", sortedArray)
	fmt.Println("the lookingFor", lookingFor)
	if index >= 0 {
		fmt.Println("Find the index:", index)
	} else {
		fmt.Println("Not Find the Data!")
	}
}

func BinarySearch(sortedArray []int, lookingFor int) int {
	low, high := 0, len(sortedArray)-1

	for low <= high {
		mid := low + (high-low)>>1
		midValue := sortedArray[mid]

		if midValue == lookingFor {
			return mid
		} else if midValue > lookingFor {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	return -1
}
