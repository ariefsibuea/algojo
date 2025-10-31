"""
LeetCode Problem : Flood Fill
Topic            : Array, Depth-First Search, Breadth-First Search, Matrix
Level            : Easy
URL              : https://leetcode.com/problems/flood-fill/
Description      : An image is represented by an m x n integer grid image where image[i][j] represents the pixel value
        of the image. You are also given three integers sr, sc, and color. You should perform a flood fill on the
        image starting from the pixel image[sr][sc]. To perform a flood fill, consider the starting pixel, plus any
        pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, plus any
        pixels connected 4-directionally to those pixels (also with the same color), and so on. Replace the color of
        all of the aforementioned pixels with color.
Examples         :
        Example 1:
        Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, color = 2
        Output: [[2,2,2],[2,2,0],[2,0,1]]
        Explanation: From the center of the image with position (sr, sc) = (1, 1) (i.e., the red pixel), all pixels
                connected by a path of the same color as the starting pixel (i.e., the blue pixels) are colored with the
                new color. Note the bottom corner is not colored 2, because it is not 4-directionally connected to the
                starting pixel.

        Example 2:
        Input: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, color = 0
        Output: [[0,0,0],[0,0,0]]
        Explanation: The starting pixel is already colored 0 so no change is made to the image.
"""

from typing import List


class Solution:
    def floodFill(self, image: List[List[int]], sr: int, sc: int, color: int) -> List[List[int]]:
        """Performs flood fill on an image similar to a paint bucket tool in painting applications.

        Args:
            image (List[List[int]]): 2D array representing image where each int is a color.
            sr (int): Starting row coordinate.
            sc (int): Starting column coordinate.
            color (int): New color to fill with.

        Returns:
            List[List[int]]: Modified image after flood fill.

        Time Complexity:
            O(N): Where N is the number of pixels in the image.

        Space Complexity:
            O(N): Space used for the stack in the worst case.
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
