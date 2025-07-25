from typing import List


class MinimumTrainPlatforms:
    def __init__(self):
        pass

    def find_minimum_platform(self, arrivals: List[str], departures: List[str]) -> int:
        """Returns the minimum number of train platforms required so that no train has to wait for a platform.

        Args:
            arrivals (List[str]): List of train arrival times in "HH:MM" format.
            departures (List[str]): List of train departure times in "HH:MM" format.

        Returns:
            int: The minimum number of platforms required to accommodate all trains without waiting.

        Raises:
            ValueError: If arrivals or departures are missing, or if their lengths do not match.

        Time Complexity:
            O(n log n): sort() uses Timsort algorithm which has O(n log n) worst-case complexity where n is the number
            of trains.

        Space Complexity:
            O(n): We need to store the converted arrival and departure times in minutes.
        """
        if not arrivals or not departures:
            raise ValueError("arrivals and departures are required")

        if len(arrivals) != len(departures):
            raise ValueError("arrivals and departures must have the same length")

        # convert string-format time into minutes
        arrivals_minute = [self.time_to_minutes(t) for t in arrivals]
        departures_minute = [self.time_to_minutes(t) for t in departures]

        # sort both arrivals and departures
        arrivals_minute.sort()
        departures_minute.sort()

        platform_needed = 0
        max_platforms = 0

        # two pointers
        i = 0
        j = 0
        n = len(arrivals)

        while i < n and j < n:
            if arrivals_minute[i] < departures_minute[j]:
                platform_needed += 1
                max_platforms = max(max_platforms, platform_needed)
                i += 1
            else:
                platform_needed -= 1
                j += 1

        # NOTE: Process remaining arrivals in case the departure finish first.
        while i < j:
            platform_needed += 1
            max_platforms = max(max_platforms, platform_needed)
            i += 1

        return max_platforms

    def time_to_minutes(self, time_str: str) -> int:
        hours, minutes = map(int, time_str.split(":"))
        return hours * 60 + minutes


if __name__ == "__main__":
    inputs = {
        "case_1": [
            ["9:00", "9:40", "9:50", "11:00", "15:00", "18:00"],
            ["9:10", "12:00", "11:20", "11:30", "19:00", "20:00"],
        ],
    }

    outputs = {"case_1": 3}

    solution = MinimumTrainPlatforms()

    for case, input in inputs.items():
        result = solution.find_minimum_platform(input[0], input[1])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")
