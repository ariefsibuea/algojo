"""
LeetCode Problem : Flood Fill
Topic            : Array, Depth-First Search, Breadth-First Search, Matrix
Level            : Easy
URL              : https://leetcode.com/problems/flood-fill/description/
"""

from typing import List


class Solution:
    def floodFill(self, image: List[List[int]], sr: int, sc: int, color: int) -> List[List[int]]:
        """
        Performs flood fill on an image by changing the color of a pixel and all adjacent pixels
        of the same original color, similar to the "paint bucket tool" in painting applications.
        Args:
            image (List[List[int]]): 2D array representing image where each int is a color
            sr (int): Starting row coordinate
            sc (int): Starting column coordinate
            color (int): New color to fill with
        Returns:
            List[List[int]]: Modified image after flood fill
        Example:
            >>> floodFill([[1,1,1],[1,1,0],[1,0,1]], 1, 1, 2)
            [[2,2,2],[2,2,0],[2,0,1]]
        Solution:
            Iterative DFS using stack
        Time Complexity:
            O(N) where N is the number of pixels in the image
        Space Complexity:
            O(N) where N is the number of pixels in the image (due to stack)
        """

        original_color = image[sr][sc]
        if original_color == color:
            return image

        stack = [(sr, sc)]
        while stack:
            r, c = stack.pop()
            if 0 <= r < len(image) and 0 <= c < len(image[0]) and image[r][c] == original_color:
                image[r][c] = color
                stack.extend([(r, c - 1), (r, c + 1), (r - 1, c), (r + 1, c)])

        return image


def run_tests():
    inputs = {"case_1": [[[1, 1, 1], [1, 1, 0], [1, 0, 1]], 1, 1, 2], "case_2": [[[0, 0, 0], [0, 0, 0]], 0, 0, 0]}
    outputs = {"case_1": [[2, 2, 2], [2, 2, 0], [2, 0, 1]], "case_2": [[0, 0, 0], [0, 0, 0]]}

    solution = Solution()

    for case, input in inputs.items():
        result = solution.floodFill(input[0], input[1], input[2], input[3])
        assert result == outputs[case], f"{case}: expected {outputs[case]}, got {result}"

    print("✅ All tests passed!")


if __name__ == "__main__":
    run_tests()
