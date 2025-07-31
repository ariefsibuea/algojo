from typing import List


def binary_search(nums: List[int], target: int) -> int:
    low = 0
    high = len(nums) - 1

    while low <= high:
        mid = (high + low) // 2

        if nums[mid] == target:
            return mid
        elif nums[mid] < target:
            low = mid + 1
        else:
            high = mid - 1

    return -1


if __name__ == "__main__":
    nums = [3, 4, 5, 6, 7, 8, 9]
    target = 4
    result = binary_search(nums, target)
    print(f"nums = {nums}, target = {target}, result = {result}")
