package selection_sort

func selectionSort(in []int) []int {
	out := make([]int, len(in))
	outIndex := 0

	for len(in) > 0 {
		low := in[0]
		lowIndex := 0

		for i := range in {
			if in[i] < low {
				low = in[i]
				lowIndex = i
			}
		}

		out[outIndex] = low
		outIndex++

		in = append(in[:lowIndex], in[lowIndex+1:]...)
	}

	return out
}
