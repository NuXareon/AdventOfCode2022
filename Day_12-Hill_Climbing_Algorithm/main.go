package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type coordinate struct {
	x, y int
}

type prioCoordinate struct {
	coordinate
	priority int
}

type coordinatePrioQueue []*prioCoordinate

func (pq coordinatePrioQueue) Len() int {return len(pq)}
func (pq coordinatePrioQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq coordinatePrioQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *coordinatePrioQueue) Push(x any) {
	coord := x.(*prioCoordinate)
	*pq = append(*pq, coord)
}
func (pq *coordinatePrioQueue) Pop() any {
	old := *pq
	n := len(old)
	coord := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return coord
}

func main() {
	heightMap, _, endPos := readMap()
	shortestPath := math.MaxInt

	for i, row := range heightMap {
		for j, cell := range row {
			if cell == 0 {
				path := doSearch(heightMap, coordinate{j, i}, endPos)
				if len(path) > 0 && len(path)-1 < shortestPath {
					shortestPath = len(path)-1
				}
			}
		}
	}

	fmt.Println(shortestPath)
}

func readMap() ([][]int, coordinate, coordinate){
	heightMap := make([][]int, 0)
	var startPos coordinate
	var endPos coordinate

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
    	line := scanner.Text()
		heightMap = append(heightMap, make([]int, 0))

		for _, val := range line {
			if (val == 'S') {
				startPos.x = len(heightMap[len(heightMap)-1])
				startPos.y = len(heightMap)-1
				val = 'a'
			} else if (val == 'E') {
				endPos.x = len(heightMap[len(heightMap)-1])
				endPos.y = len(heightMap)-1
				val = 'z'
			}
			
			heightMap[len(heightMap)-1] = append(heightMap[len(heightMap)-1], int(val-'a'))
		}
	}

	return heightMap, startPos, endPos
}

func doSearch(heightMap [][]int, startPos coordinate, endPos coordinate) []coordinate {
	previousNode := make(map[coordinate]coordinate, 0)
	nodesToExpand := make(coordinatePrioQueue, 0)
	var startPrioCoordinate prioCoordinate
	startPrioCoordinate.coordinate = startPos
	heap.Push(&nodesToExpand, &startPrioCoordinate)
	nodeScore := make(map[coordinate]int)
	nodeScore[startPos] = 0

	for len(nodesToExpand) > 0 {
		currentCoord := heap.Pop(&nodesToExpand).(*prioCoordinate)

		if currentCoord.coordinate == endPos {
			return reconstructPath(previousNode, endPos)
		}

		neighbours := []coordinate{
			{currentCoord.coordinate.x+1, currentCoord.coordinate.y},
			{currentCoord.coordinate.x, currentCoord.coordinate.y+1},
			{currentCoord.coordinate.x-1, currentCoord.coordinate.y},
			{currentCoord.coordinate.x, currentCoord.coordinate.y-1},
		}

		for _, neighbourCoord := range neighbours {
			if neighbourCoord.y < 0 || neighbourCoord.y > len(heightMap)-1 ||
			neighbourCoord.x < 0 || neighbourCoord.x > len(heightMap[neighbourCoord.y])-1 {
				// coordinate out of bounds
				continue
			}

			if heightMap[neighbourCoord.y][neighbourCoord.x] - heightMap[currentCoord.y][currentCoord.x] > 1 {
				// height different too high (only for higher values)
				continue
			}

			neighbourScoreFromCurrent := nodeScore[currentCoord.coordinate]+1
			if neighbourScore, found := nodeScore[neighbourCoord]; !found || neighbourScoreFromCurrent < neighbourScore  {
				previousNode[neighbourCoord] = currentCoord.coordinate
				nodeScore[neighbourCoord] = neighbourScoreFromCurrent
				var nodeToExpand prioCoordinate
				nodeToExpand.coordinate = neighbourCoord
				nodeToExpand.priority = neighbourScoreFromCurrent + heightMapHeuristic(neighbourCoord, endPos)
				heap.Push(&nodesToExpand, &nodeToExpand)
			}
		}
	}

	return []coordinate{}
}

func reconstructPath(previousNode map[coordinate]coordinate, end coordinate) []coordinate {
	path := []coordinate{end}
	currentCoord := end
	for val, found := previousNode[currentCoord]; found; val, found = previousNode[currentCoord] {
		path = append(path, val)
		currentCoord = val
	}
	return path
}

// Why tf does go not provide math operations for ints
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func heightMapHeuristic(currentPos coordinate, endPos coordinate) int {
	// Manhattan distance, since we operate in a grid. 
	// That being said, I think the conditions on which nodes we can expand kind of make this a bad heuristic (but still admisible)
	return abs(endPos.x - currentPos.x)+abs(endPos.y - currentPos.y)
}