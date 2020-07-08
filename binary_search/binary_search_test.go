package binary_search

import "testing"

var numbers []int = []int{1, 2, 3, 4, 6, 7, 8}

func TestSearchNotFoundLow(t *testing.T) {
	idx, err := search(0, numbers)

	if idx != -1 {
		t.Errorf("got %d, want -1", idx)
	}

	if err != notFoundError {
		t.Errorf("got %v, want %s", err, notFoundError)
	}
}

func TestSearchNotFoundHigh(t *testing.T) {
	idx, err := search(10, numbers)

	if idx != -1 {
		t.Errorf("got %d, want -1", idx)
	}

	if err != notFoundError {
		t.Errorf("got %v, want %s", err, notFoundError)
	}
}

func TestSearchNotFoundMid(t *testing.T) {
	idx, err := search(5, numbers)

	if idx != -1 {
		t.Errorf("got %d, want -1", idx)
	}

	if err != notFoundError {
		t.Errorf("got %v, want %s", err, notFoundError)
	}
}

func TestSearchFound(t *testing.T) {
	for i, num := range numbers {
		idx, err := search(num, numbers)

		if idx != i {
			t.Errorf("got %d, want %d", idx, i)
		}

		if err != nil {
			t.Errorf("got unexpected error %v for search(%d, numbers)", err, notFoundError)
		}
	}
}
