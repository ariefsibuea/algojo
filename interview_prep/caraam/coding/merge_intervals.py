from typing import List


class MergeIntervals:
    def __init__(self):
        pass

    def solve_by_interval_merging(self, intervals: List[List[int]]) -> List[List[int]]:
        """Merges overlapping intervals in a list

        Merge the overlapping intervals using interval merging (sort and scan). First, sorting them and then performing
        a linear scan to combine overlaps.

        Args:
            intervals (List[List[int]]): A list of intervals, where each interval is represented as a list of two
            integers [start, end].

        Returns:
            List[List[int]]: A list of merged, non-overlapping intervals.

        Raises:
            ValueError: If the input list 'intervals' is empty.

        Time Complexity:
            O(n log n): Where n is the number of intervals, due to the sorting step.

        Space Complexity:
            O(n): For storing the result list of merged intervals.
        """

        if not intervals:
            raise ValueError("'intervals' must have at least 1 element")

        intervals.sort(key=lambda x: x[0])

        result = []
        current_start, current_end = intervals[0]

        for interval in intervals[1:]:
            next_start, next_end = interval

            if next_start <= current_end:
                # overlap
                current_end = max(current_end, next_end)
            else:
                result.append([current_start, current_end])
                current_start, current_end = next_start, next_end

        # Add the latest interval
        result.append([current_start, current_end])

        return result


if __name__ == "__main__":
    inputs = {
        "case_1": [[[1, 3], [2, 6], [8, 10], [15, 18]]],
        "case_2": [[[1, 4], [4, 5]]],
        "case_3": [[[1, 4], [2, 3]]],
        "case_4": [[[1, 4], [1, 2]]],
        "case_5": [[[1, 4]]],
    }

    outputs = {
        "case_1": [[1, 6], [8, 10], [15, 18]],
        "case_2": [[1, 5]],
        "case_3": [[1, 4]],
        "case_4": [[1, 4]],
        "case_5": [[1, 4]],
    }

    solution = MergeIntervals()

    for case, input in inputs.items():
        result = solution.solve_by_interval_merging(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")
