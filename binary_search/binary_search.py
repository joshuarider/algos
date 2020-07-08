#!/usr/bin/python3

from typing import List, Optional


def search(needle: int, haystack: List[int]) -> Optional[int]:
    low = 0
    high = len(haystack) - 1

    while low <= high:
        mid = (low + high) // 2

        if haystack[mid] == needle:
            return mid
        elif haystack[mid] < needle:
            low = mid + 1
        elif haystack[mid] > needle:
            high = mid - 1

    return None

