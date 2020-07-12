package selection_sort

import "testing"

func TestSelectionSortEmpty(t *testing.T) {
	got := selectionSort([]int{})

	if len(got) != 0 {
		t.Errorf("expected empty slice, got %v", got)
	}
}

func TestSelectionSortOneElement(t *testing.T) {
	want := []int{1}
	got := selectionSort([]int{1})

	if !same(got, want) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestSelectionSortReversed(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}
	got := selectionSort([]int{5, 4, 3, 2, 1})

	if !same(got, want) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestSelectionSortSorted(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}
	got := selectionSort([]int{1, 2, 3, 4, 5})

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
