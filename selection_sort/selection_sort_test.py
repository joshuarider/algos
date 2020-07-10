#!/usr/bin/python3

import unittest

from selection_sort import selection_sort


class TestSelectionSort(unittest.TestCase):
    def test_sort_empty(self):
        self.assertEqual([], selection_sort([]))

    def test_sort_one_element(self):
        self.assertEqual([2], selection_sort([2]))

    def test_sort_reversed(self):
        self.assertEqual([0, 1, 2, 3, 4, 5], selection_sort([5, 4, 3, 2, 1, 0]))

    def test_sort_sorted(self):
        self.assertEqual([0, 1, 2, 3, 4, 5], selection_sort([0, 1, 2, 3, 4, 5]))


if __name__ == "__main__":
    unittest.main()
