#!/usr/bin/python3

from typing import List

def quicksort(nums: List[int]) -> List[int]:
    if len(nums) < 2:
        return nums

    pivot = nums[0]

    lower = []
    higher = []

    for i in range(1, len(nums)):
        if nums[i] <= pivot:
            lower.append(nums[i])
        elif nums[i] > pivot:
            higher.append(nums[i])

        i += 1

    return quicksort(lower) + [pivot] + quicksort(higher)
