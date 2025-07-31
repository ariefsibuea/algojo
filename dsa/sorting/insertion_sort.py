from typing import List


def insertion_sort(nums: List[int]):
    for i in range(1, len(nums)):
        key = nums[i]
        j = i - 1

        while j >= 0 and nums[j] > key:
            nums[j + 1] = nums[j]
            j -= 1

        nums[j + 1] = key


if __name__ == "__main__":
    nums = [9, 5, 1, 4, 3]
    print(f"nums = {nums}")
    insertion_sort(nums)
    print(f"result = {nums}")
