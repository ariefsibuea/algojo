package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/*
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
*/

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if color == image[sr][sc] {
		return image
	}
	return changeColor(image, sr, sc, color, image[sr][sc])
}

func changeColor(image [][]int, sr int, sc int, color, originalColor int) [][]int {
	// check row
	if sr < 0 || sr >= len(image) {
		return image
	}
	// check column
	if sc < 0 || sc >= len(image[0]) {
		return image
	}
	// check original color
	if image[sr][sc] != originalColor {
		return image
	}

	image[sr][sc] = color
	// top
	image = changeColor(image, sr-1, sc, color, originalColor)
	// left
	image = changeColor(image, sr, sc-1, color, originalColor)
	// bottom
	image = changeColor(image, sr+1, sc, color, originalColor)
	// right
	image = changeColor(image, sr, sc+1, color, originalColor)

	return image
}

func RunTestFloodFill() {
	testCases := map[string]struct {
		image  [][]int
		sr     int
		sc     int
		color  int
		expect [][]int
	}{
		"case-1": {
			image: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},
			sr:    1,
			sc:    1,
			color: 2,
			expect: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
		"case-2": {
			image: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
			sr:    0,
			sc:    0,
			color: 0,
			expect: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := floodFill(testCase.image, testCase.sr, testCase.sc, testCase.color)

		for i, expectRow := range testCase.expect {
			if !cmp.EqualSlices(expectRow, result[i]) {
				fmt.Printf("==== FAILED: expect = %v - got = %v in the index-%d\n", expectRow, result[i], i)
				os.Exit(1)
			}
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
