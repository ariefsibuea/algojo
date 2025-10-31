"""
LeetCode Problem : Minimum Window Substring
Topic            : Hash Table, String, Sliding Window
Level            : Hard
URL              : https://leetcode.com/problems/minimum-window-substring
"""


class Solution:
    def solve(self, s: str, t: str) -> str:
        target = {}
        for c in t:
            target[c] = 1 + target.get(c, 0)

        required_len = len(target)

        source = {}
        matched_len = 0
        min_window_len = float("inf")

        start = 0
        final_start = 0

        for i, c in enumerate(s):
            if c in target:
                source[c] = 1 + source.get(c, 0)
                if source[c] == target[c]:
                    matched_len += 1

            while matched_len == required_len:
                if (i - start + 1) < min_window_len:
                    min_window_len = i - start + 1
                    final_start = start

                rm_c = s[start]
                if rm_c in target:
                    source[rm_c] = source.get(rm_c, 0) - 1
                    if source[rm_c] < target[rm_c]:
                        matched_len -= 1

                start += 1

        if min_window_len == float("inf"):
            return ""

        return s[final_start : final_start + min_window_len]


def run_tests():
    inputs = {
        "case_1": ["ADOBECODEBANC", "ABC"],
        "case_2": ["a", "a"],
        "case_3": ["a", "aa"],
    }
    outputs = {
        "case_1": "BANC",
        "case_2": "a",
        "case_3": "",
    }

    solution = Solution()

    for case, input in inputs.items():
        result = solution.solve(input[0], input[1])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
