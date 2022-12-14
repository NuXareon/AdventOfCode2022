package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type terrainType int

const (
	Air terrainType = iota
	Rock
	Sand
)

func min(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	input, minCoord, maxCoord := readInput()

	minCoord[0] = min(minCoord[0], 500)
	minCoord[1] = min(minCoord[1], 0)
	maxCoord[0] = max(maxCoord[0], 500)
	maxCoord[1] = max(maxCoord[1], 0)

	grid := generateGrid(input, minCoord, maxCoord)
	grid = append(grid, make([]terrainType, maxCoord[0]-minCoord[0]+1))
	grid = append(grid, make([]terrainType, maxCoord[0]-minCoord[0]+1))
	maxCoord[1] += 2
	startPos := []int64{500, 0}

	// Resize grid to fit the whole pile of sand
	gridXLeftSize := 1+(maxCoord[1]-minCoord[1])-(startPos[0]-minCoord[0])
	gridXRightSize := 1+(maxCoord[1]-minCoord[1])-(maxCoord[0]-startPos[0])
	if gridXLeftSize > 0 {
		resizedGrid := make([][]terrainType, 0, len(grid))

		for _, val := range grid {
			resizedGrid = append(resizedGrid, make([]terrainType, gridXLeftSize))
			resizedGrid[len(resizedGrid)-1] = append(resizedGrid[len(resizedGrid)-1], val...)
		}

		minCoord[0] -= gridXLeftSize
		grid = resizedGrid
	}
	if gridXRightSize > 0 {
		for i := range grid {
			grid[i] = append(grid[i], make([]terrainType, gridXRightSize)...)
		}
	}

	// Add floor
	for i := range grid[len(grid)-1] {
		grid[len(grid)-1][i] = Rock
	}

	grid, sand := processGrid(grid, minCoord)

	for _, val := range grid {
		fmt.Println(val)
	}
	fmt.Println(sand)
}

func readInput() ([][][]int64, []int64, []int64) {
	input := make([][][]int64, 0)
	var minCoord []int64
	var maxCoord []int64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
    	wallSegment := scanner.Text()
		input = append(input, make([][]int64, 0))
		wallSplit := strings.Split(wallSegment, " -> ")

		for _, point := range wallSplit {
			pointSplit := strings.Split(point, ",")
			var pointSplitInt []int64
			for _, val := range pointSplit {
				valInt, _ := strconv.ParseInt(val, 10, 0)
				pointSplitInt = append(pointSplitInt, valInt)
			}

			if minCoord == nil {
				minCoord = make([]int64, len(pointSplitInt))
				copy(minCoord, pointSplitInt)
			} else {
				minCoord[0] = min(minCoord[0], pointSplitInt[0])
				minCoord[1] = min(minCoord[1], pointSplitInt[1])
			}
			if maxCoord == nil {
				maxCoord = make([]int64, len(pointSplitInt))
				copy(maxCoord, pointSplitInt)
			} else {
				maxCoord[0] = max(maxCoord[0], pointSplitInt[0])
				maxCoord[1] = max(maxCoord[1], pointSplitInt[1])
			}

			input[len(input)-1] = append(input[len(input)-1], pointSplitInt)
		}
	}

	return input, minCoord, maxCoord
}

func generateGrid(input [][][]int64, minPos []int64, maxPos []int64) [][]terrainType{
	grid := make([][]terrainType, 0, maxPos[1]-minPos[1])
	for i := int64(0); i <= maxPos[1]-minPos[1]; i++ {
		grid = append(grid, make([]terrainType, maxPos[0]-minPos[0]+1))
	}

	for _, line := range input {
		var previousPoint []int64
		for _, point := range line {
			point[0] -= minPos[0]
			point[1] -= minPos[1]
			if previousPoint == nil {
				previousPoint = point
				continue
			}

			// Down
			direction := []int64{0,1}
			if point[0] < previousPoint[0] {
				direction = []int64{-1,0}
			} else if point[0] > previousPoint[0] {
				direction = []int64{1,0}
			} else if point[1] < previousPoint[1] {
				direction = []int64{0,-1}
			}

			for currentPoint := previousPoint; currentPoint[0] != point[0] || currentPoint[1] != point[1]; currentPoint[0], currentPoint[1] = currentPoint[0]+direction[0], currentPoint[1]+direction[1] {
				grid[currentPoint[1]][currentPoint[0]] = Rock
			}
			grid[point[1]][point[0]] = Rock

			previousPoint = point
		}
	}
	return grid
}

func processGrid(grid [][]terrainType, minPos []int64) ([][]terrainType, int) {
	sandGrains := 0
	startPos := []int64{500, 0}
	startPos[0] -= minPos[0]
	startPos[1] -= minPos[1]

	for {
		if grid[startPos[1]][startPos[0]] == Sand {
			return grid, sandGrains
		}

		pos := make([]int64, len(startPos))
		copy(pos, startPos)

		pos = simulateSand(pos, grid)
		
		if pos[1] < 0 || pos[1] >= int64(len(grid)) || pos[0] < 0 || pos[0] >= int64(len(grid[pos[1]])) {
			// Fell to the void
			break
		}

		grid[pos[1]][pos[0]] = Sand
		sandGrains++
	}

	return grid, sandGrains
}

func simulateSand(pos []int64, grid [][]terrainType) []int64 {
	for {
		// check down
		pos[1] += 1
		if pos[1] >= int64(len(grid)) {
			// Fell to the void
			return pos
		}
		if grid[pos[1]][pos[0]] == Air {
			continue
		}

		// check left (and down)
		pos[0] -= 1
		if pos[0] < 0 {
			// Fell to the void
			return pos
		}
		if grid[pos[1]][pos[0]] == Air {
			continue
		}

		// check right (and down)
		pos[0] += 2
		if pos[0] >= int64(len(grid[pos[1]])) {
			// Fell to the void
			return pos
		}
		if grid[pos[1]][pos[0]] == Air {
			continue
		}

		// return to previous position
		pos[0] -= 1
		pos[1] -= 1
		return pos
	}
}