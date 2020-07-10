#!/usr/bin/python3

from typing import List


def selection_sort(nums: List[int]) -> List[int]:
    out = []

    while nums:
        lowest = nums[0]
        lowest_idx = 0

        for i in range(0, len(nums)):
            if nums[i] < lowest:
                lowest = nums[i]
                lowest_idx = i

        out.append(nums.pop(lowest_idx))

    return out
