from typing import List


def partition(nums: List[int], low: int, high: int) -> int:
    # choose the rightmost element as pivot
    pivot = nums[high]

    # pointer for greater element
    i = low - 1

    # traverse through all elements, compare each element with pivot
    for j in range(low, high):
        if nums[j] <= pivot:
            # if element smaller than pivot is found, swap it with the greater element pointed by i
            i = i + 1
            # swapping element at i with element at j
            nums[i], nums[j] = nums[j], nums[i]

    # swap the pivot element with the greater element specified by i
    nums[i + 1], nums[high] = nums[high], nums[i + 1]

    return i + 1


def quick_sort(nums: List[int], low: int, high: int):
    if low < high:
        pi = partition(nums, low, high)

        # recursive call on the left pivot
        quick_sort(nums, low, pi - 1)
        # recursive call on the right pivot
        quick_sort(nums, pi + 1, high)


if __name__ == "__main__":
    nums = [8, 7, 2, 1, 0, 9, 6]
    print(f"nums = {nums}")
    quick_sort(nums, 0, len(nums) - 1)
    print(f"result = {nums}")
