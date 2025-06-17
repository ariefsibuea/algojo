"""
LeetCode Problem : Valid Palindrome
Topic            : Two Pointers, String
Level            : Easy
URL              : https://leetcode.com/problems/valid-palindrome/description/
"""


class Solution:
    def isPalindrome(self, s: str) -> bool:
        """
        Check if a string is a palindrome considering only alphanumeric characters and ignoring case.
        A string is a palindrome if it reads the same forward and backward after converting all uppercase letters
        into lowercase letters and removing all non-alphanumeric characters.
        Args:
            s (str): The input string to check.
        Returns:
            bool: True if the input string is a palindrome, False otherwise.
        Example:
            >>> solution = Solution()
            >>> solution.isPalindrome("A man, a plan, a canal: Panama")
            True
            >>> solution.isPalindrome("race a car")
            False
            >>> solution.isPalindrome(" ")
            True
        Note:
            - For non-alphanumeric characters (spaces, punctuation), they are ignored in the palindrome check
            - The comparison is case-insensitive
            - An empty string is considered a palindrome
        Solution:
            Two-pointer approach
        Time Complexity:
            O(n) where n is the length of the string
        Space Complexity:
            O(1) only using two pointers
        """

        i, j = 0, len(s) - 1
        while i <= j:
            while i < j and not s[i].isalnum():
                i += 1
            while i < j and not s[j].isalnum():
                j -= 1

            if s[i].lower() != s[j].lower():
                return False

            i += 1
            j -= 1

        return True


def run_tests():
    inputs = {"case_1": ["A man, a plan, a canal: Panama"], "case_2": ["race a car"]}
    outputs = {
        "case_1": True,
        "case_2": False,
    }

    solution = Solution()

    for case, input in inputs.items():
        result = solution.isPalindrome(input[0])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
