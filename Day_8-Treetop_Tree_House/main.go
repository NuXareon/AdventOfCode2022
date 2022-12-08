package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	forest := make([][]int, 0)
	
	for scanner.Scan() {
    	line := scanner.Text()
		forest = append(forest, make([]int, 0))
		forestIdx := len(forest)-1

		for _, tree := range line {
			forest[forestIdx] = append(forest[forestIdx], int(tree-'0'))
		}
	}

	//visibleTrees := len(forest)*2 + len(forest[0])*2 - 4
	bestScenicScore := int64(0)

	for i, treeLine := range forest {
		if i == 0 || i == len(forest)-1 {
			continue
		}
		
		for j := range treeLine {
			if j == 0 || j == len(treeLine)-1 {
				continue
			}

			//isTreeVisible := checkTreeVisibility(forest, i, j)
			scenicScore := calculateScenicScore(forest, i, j)

			if scenicScore > bestScenicScore {
				bestScenicScore = scenicScore
			}
		}
	}

	fmt.Println(bestScenicScore)
}

// FIRST PART
func checkTreeVisibility(forest [][]int, i int, j int) bool {
	// Up
	if (checkTreeVisibilityDirection(forest, i, j, 0, 1)) { 
		return true
	}
	// Right
	if (checkTreeVisibilityDirection(forest, i, j, 1, 0)) { 
		return true
	}
	// Down
	if (checkTreeVisibilityDirection(forest, i, j, 0, -1)) { 
		return true
	}
	// Left
	if (checkTreeVisibilityDirection(forest, i, j, -1, 0)) { 
		return true
	}
	return false
}

func checkTreeVisibilityDirection(forest [][]int, i int, j int, x int, y int) bool {
	treeVal := forest[i][j]
	i += y
	j += x
	for i >= 0 && i < len(forest) && j >= 0 && j < len(forest[i]) {
		if (forest[i][j] >= treeVal) {
			return false
		}
		i += y
		j += x
	}
	return true
}

// SECOND PART
func calculateScenicScore(forest [][]int, i int, j int) int64 {
	// Up
	scenicScore := calculateScenicScoreDirection(forest, i, j, 0, 1)
	// Right
	scenicScore *= calculateScenicScoreDirection(forest, i, j, 1, 0)
	// Down
	scenicScore *= calculateScenicScoreDirection(forest, i, j, 0, -1)
	// Left
	scenicScore *= calculateScenicScoreDirection(forest, i, j, -1, 0)

	return scenicScore
}

func calculateScenicScoreDirection(forest [][]int, i int, j int, x int, y int) int64 {
	scenicScore := int64(0)
	treeVal := forest[i][j]
	i += y
	j += x
	for i >= 0 && i < len(forest) && j >= 0 && j < len(forest[i]) {
		scenicScore++
		if (forest[i][j] >= treeVal) {
			break
		}
		i += y
		j += x
	}
	return scenicScore
}