from typing import List


class NonOverlappingIntervals:
    def __init__(self):
        pass

    def solve_by_earliest_finish_time(self, intervals: List[List[int]]) -> int:
        """Finds the minimum number of intervals to remove to make the rest of the intervals non-overlapping.

        The algorithm is the Interval Scheduling algorithm, also commonly known as the Activity Selection algorithm.
        More specifically, it uses the Earliest Finish Time greedy strategy. This is a classic greedy algorithm used
        for scheduling and resource allocation problems. First, sort all intervals by their end time, then select the
        intervals that don't overlap with the previously selected one.

        Args:
            intervals (List[List[int]]): A list of intervals, where each interval is represented as a list [start, end].

        Returns:
            int: The minimum number of intervals that need to be removed to eliminate all overlaps.

        Raises:
            ValueError: If 'intervals' is empty.

        Time Complexity:
            O(n log n): Where n is the number of intervals, due to sorting.

        Space Complexity:
            O(1): As the algorithm uses only a constant amount of extra space.
        """

        if not intervals:
            raise ValueError("'intervals' must have at least 1 element")

        # Sort by end time (ascending)
        intervals.sort(key=lambda x: x[1])

        kept = 1  # Keep the first earliest interval
        last_end = intervals[0][1]

        for i in range(1, len(intervals)):
            if intervals[i][0] >= last_end:
                # non overlapping
                kept += 1
                last_end = intervals[i][1]

        return len(intervals) - kept


if __name__ == "__main__":
    inputs = {
        "case_1": [[[1, 2], [2, 3], [3, 4], [1, 3]]],
        "case_2": [[[1, 2], [1, 2], [1, 2]]],
        "case_3": [[[1, 2], [2, 3]]],
    }

    outputs = {
        "case_1": 1,
        "case_2": 2,
        "case_3": 0,
    }

    solution = NonOverlappingIntervals()

    for case, input in inputs.items():
        result = solution.solve_by_earliest_finish_time(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")
