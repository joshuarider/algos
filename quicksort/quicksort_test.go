package quicksort

import "testing"

func TestQuicksortEmpty(t *testing.T) {
	got := quicksort([]int{})

	if len(got) != 0 {
		t.Errorf("expected empty slice, got %v", got)
	}
}

func TestQuicksortOneElement(t *testing.T) {
	want := []int{1}
	got := quicksort([]int{1})

	if !same(got, want) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestQuicksortReversed(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}
	got := quicksort([]int{5, 4, 3, 2, 1})

	if !same(got, want) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestQuicksortSorted(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}
	got := quicksort([]int{1, 2, 3, 4, 5})

	if !same(got, want) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestQuicksortJumbled(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}
	got := quicksort([]int{4, 2, 1, 5, 3})

	if !same(got, want) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestQuicksortDuplicates(t *testing.T) {
	want := []int{1, 2, 2, 3, 4, 4, 5, 5}
	got := quicksort([]int{2, 4, 4, 2, 1, 5, 3, 5})

	if !same(got, want) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func same(a []int, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
