package binary_search

import "errors"

var notFoundError error = errors.New("not found")

func search(needle int, haystack []int) (int, error) {
	low := 0
	high := len(haystack) - 1

	for low <= high {
		mid := (low + high) / 2

		if needle == haystack[mid] {
			return mid, nil
		} else if haystack[mid] < needle {
			low = mid + 1
		} else if haystack[mid] > needle {
			high = mid - 1
		}
	}

	return -1, notFoundError
}
