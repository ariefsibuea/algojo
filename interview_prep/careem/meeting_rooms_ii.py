# Definition of Interval:
from typing import List


class Interval(object):
    def __init__(self, start, end):
        self.start = start
        self.end = end


class MeetingRoomsII:
    def __init__(self):
        pass

    def solve_by_interval_sorting(self, intervals: List[Interval]):
        """Determines the minimum number of meeting rooms required to accommodate all given meeting intervals using
        interval sorting.

        The algorithm sorts start and end times separately and uses two pointers to track room usage.

        Args:
            intervals (List[Interval]): A list of Interval objects, each with 'start' and 'end' attributes representing
            meeting times.

        Returns:
            int: The minimum number of meeting rooms required.

        Raises:
            ValueError: If 'intervals' is empty.

        Time Complexity:
            O(n log n): Where n is the number of intervals, due to sorting the start and end times.

        Space Complexity:
            O(n): For storing the sorted start and end times.
        """

        if not intervals:
            raise ValueError("'intervals' must have at least 1 element")

        start_times = sorted([i.start for i in intervals])
        end_times = sorted([i.end for i in intervals])

        start_ptr, end_ptr = 0, 0
        max_rooms = 0
        rooms_needed = 0

        while start_ptr < len(intervals):
            if start_times[start_ptr] < end_times[end_ptr]:
                rooms_needed += 1
                max_rooms = max(max_rooms, rooms_needed)
                start_ptr += 1
            else:
                rooms_needed -= 1
                end_ptr += 1

        return max_rooms


if __name__ == "__main__":
    inputs = {
        "case_1": [[Interval(0, 30), Interval(5, 10), Interval(15, 20)]],
        "case_2": [[Interval(7, 10), Interval(2, 4)]],
        "case_3": [[Interval(0, 10), Interval(5, 15), Interval(10, 20), Interval(15, 25)]],
    }

    outputs = {
        "case_1": 2,
        "case_2": 1,
        "case_3": 2,
    }

    solution = MeetingRoomsII()

    for case, input in inputs.items():
        result = solution.solve_by_interval_sorting(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")
