package quicksort

func quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[0]
	lower := make([]int, 0)
	higher := make([]int, 0)

	for i := 1; i < len(nums); i++ {
		if nums[i] <= pivot {
			lower = append(lower, nums[i])
		} else if nums[i] > pivot {
			higher = append(higher, nums[i])
		}
	}

	return append(quicksort(lower), append([]int{pivot}, quicksort(higher)...)...)
}
