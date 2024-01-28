package floodfill

import (
	"slices"
)

// https://leetcode.com/problems/flood-fill/
func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	root := image[sr][sc]
	visited := [][]int{}

	queue := [][]int{{sr, sc}}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		visited = append(visited, node)
		image[node[0]][node[1]] = color

		for _, neighbor := range neighbors(image, node[0], node[1], root) {
			if !contains(visited, neighbor) {
				queue = append(queue, neighbor)
			}
		}
	}

	return image
}

func neighbors(image [][]int, sr, sc, root int) [][]int {
	possibleNeighbors := [][]int{
		{sr - 1, sc},
		{sr + 1, sc},
		{sr, sc - 1},
		{sr, sc + 1},
	}

	neighbors := [][]int{}

	for _, pN := range possibleNeighbors {
		if isValid(image, pN[0], pN[1]) {
			if root == image[pN[0]][pN[1]] {
				neighbors = append(neighbors, pN)
			}
		}
	}

	return neighbors
}

func isValid(image [][]int, sr, sc int) bool {
	return sr >= 0 && sr < len(image) && sc >= 0 && sc < len(image[0])
}

func contains(visited [][]int, node []int) bool {
	for _, v := range visited {
		if slices.Equal(v, node) {
			return true
		}
	}

	return false
}
