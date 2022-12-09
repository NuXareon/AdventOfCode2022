package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type mapPos struct {
	x, y int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	positionsVisited := make(map[mapPos]bool, 0)
	tailKnots := make([]mapPos, 10)

	for scanner.Scan() {
    	line := scanner.Text()
		splitLine := strings.Split(line, " ")
		direction := splitLine[0]
		movement, _ := strconv.ParseInt(splitLine[1], 10, 0)
		x, y := getDirectionMovement(direction)
		for i := 0; int64(i) < movement; i++ {
			tailKnots[0].x += x
			tailKnots[0].y += y
			for i, knotPos := range tailKnots {
				if i == 0 {
					continue
				}
				tailKnots[i] = adjustTailPos(knotPos, tailKnots[i-1])
			}
			positionsVisited[tailKnots[len(tailKnots)-1]] = true
		}
	}

	fmt.Println(len(positionsVisited))
}

func getDirectionMovement(direction string) (int, int) {
	switch direction {
	case "R" :
		return 1, 0
	case "D":
		return 0, -1
	case "L":
		return -1, 0
	case "U":
		return 0, 1
	}
	return 0, 0
}

func adjustTailPos(tailPos mapPos, headPos mapPos) mapPos {
	var distance mapPos
	distance.x = headPos.x-tailPos.x
	distance.y = headPos.y-tailPos.y
	if distance.x > 1 {
		tailPos.x++
		if distance.y > 0 {
			tailPos.y++
		} else if distance.y < 0 {
			tailPos.y--
		}
	} else if distance.x < -1 {
		tailPos.x--
		if distance.y > 0 {
			tailPos.y++
		} else if distance.y < 0 {
			tailPos.y--
		}
	} else if distance.y > 1 {
		tailPos.y++
		if distance.x > 0 {
			tailPos.x++
		} else if distance.x < 0 {
			tailPos.x--
		}
	} else if distance.y < -1 {
		tailPos.y--
		if distance.x > 0 {
			tailPos.x++
		} else if distance.x < 0 {
			tailPos.x--
		}
	}
	return tailPos
}