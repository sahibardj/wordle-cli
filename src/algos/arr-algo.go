package algos

import (
	"math/rand"
	"time"
)

// Fischer-Yates in-place shuffle.
func Shuffle(arr []string) {
	rand.Seed(time.Now().UnixNano())
	arr_len := len(arr)
	for i := arr_len - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[j], arr[i] = arr[i], arr[j]
	}
}

// Return index of element, or -1 if not found.
func BinarySearch(arr []string, val string) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) >> 1
		if arr[mid] == val {
			return mid
		} else if arr[mid] > val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
