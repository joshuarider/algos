#!/usr/bin/python3

import unittest

from quicksort import quicksort


class TestQuicksort(unittest.TestCase):
    def test_sort_empty(self):
        self.assertEqual([], quicksort([]))

    def test_sort_one_element(self):
        self.assertEqual([2], quicksort([2]))

    def test_sort_reversed(self):
        self.assertEqual([0, 1, 2, 3, 4, 5], quicksort([5, 4, 3, 2, 1, 0]))

    def test_sort_sorted(self):
        self.assertEqual([0, 1, 2, 3, 4, 5], quicksort([0, 1, 2, 3, 4, 5]))

    def test_sort_jumble(self):
        self.assertEqual([0, 1, 2, 3, 4, 5], quicksort([3, 2, 4, 1, 0, 5]))

    def test_sort_duplicates(self):
        self.assertEqual([0, 1, 1, 2, 3, 3, 4, 4, 5], quicksort([1, 3, 2, 4, 1, 0, 5, 4, 3]))


if __name__ == "__main__":
    unittest.main()
