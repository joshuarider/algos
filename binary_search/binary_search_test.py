#!/usr/bin/python3

import unittest

from binary_search import search


NUMBERS = [1, 2, 3, 5, 6, 7, 8]

class TestBinarySearch(unittest.TestCase):
    def test_search_not_found_low(self):
        self.assertIsNone(search(0, NUMBERS))

    def test_search_not_found_high(self):
        self.assertIsNone(search(9, NUMBERS))

    def test_search_not_found_mid(self):
        self.assertIsNone(search(4, NUMBERS))

    def test_search_found(self):
        for i, num in enumerate(NUMBERS):
            self.assertEqual(i, search(num, NUMBERS))


if __name__ == "__main__":
    unittest.main()
