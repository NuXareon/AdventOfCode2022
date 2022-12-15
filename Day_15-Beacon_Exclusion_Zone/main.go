package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxBeaconCoord = 4000000

type point struct {
	x,y int64
}

type sensor struct {
	pos point
	dist int64
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	intersectionY := int64(2000000)
	intersectionPoints := make(map[point]bool, 0)
	sensorPoints := make([]point, 0)
	beaconPoints := make([]point, 0)

	sensors := make([]sensor, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
    	line := scanner.Text()

		line = strings.TrimPrefix(line, "Sensor at x=")
		sensorX, line, _ := strings.Cut(line, ",")
		sensorXCoord, _ := strconv.ParseInt(sensorX, 10, 0)

		line = strings.TrimPrefix(line, " y=")
		sensorY, line, _ := strings.Cut(line, ":")
		sensorYCoord, _ := strconv.ParseInt(sensorY, 10, 0)
		
		line = strings.TrimPrefix(line, " closest beacon is at x=")
		beaconX, line, _ := strings.Cut(line, ",")
		beaconXCoord, _ := strconv.ParseInt(beaconX, 10, 0)
		
		line = strings.TrimPrefix(line, " y=")
		beaconYCoord, _ := strconv.ParseInt(line, 10, 0)

		sensorPoints = append(sensorPoints, point{sensorXCoord, sensorYCoord})
		beaconPoints = append(beaconPoints, point{beaconXCoord, beaconYCoord})

		distance := abs(sensorXCoord-beaconXCoord) + abs(sensorYCoord - beaconYCoord)

		// Part 1
		intersectionPoints = processIntersection(intersectionPoints, distance, sensorXCoord, sensorYCoord, intersectionY)

		// part 2
		sensors = append(sensors, sensor{point{sensorXCoord, sensorYCoord}, distance})
	}

	// Part 1
	printPart1(intersectionPoints, sensorPoints, beaconPoints)

	// Part 2
	for _, s := range sensors {
		intersectionPoint := checkSensorIntersection(sensors, s)

		if intersectionPoint != nil {
			tunningFrequency := intersectionPoint.x * 4000000 + intersectionPoint.y
			fmt.Println(tunningFrequency)
			break;
		}
	}
}

func processIntersection(intersectionPoints map[point]bool, distance int64, sensorXCoord int64, sensorYCoord int64, intersectionY int64) map[point]bool {
	if distance - abs(sensorYCoord - intersectionY) < 0 {
		// Too far on y
		return intersectionPoints
	}

	intersectionX1 := -distance + abs(sensorYCoord - intersectionY) + sensorXCoord
	intersectionX2 := distance - abs(sensorYCoord - intersectionY) + sensorXCoord

	for ;intersectionX1 <= intersectionX2; intersectionX1++ {
		intersectionPoints[point{intersectionX1, intersectionY}] = true
	}

	return intersectionPoints
}

func printPart1(intersectionPoints map[point]bool, sensorPoints []point, beaconPoints []point) {
	for _, val := range sensorPoints {
		delete(intersectionPoints, val)
	}

	for _, val := range beaconPoints {
		delete(intersectionPoints, val)
	}

	fmt.Println(len(intersectionPoints))
}

// From the point of view of a c++ programmer, returning a pointer to a local variable is such a travesty. But let's roll with it since GO handles it (and doesn't have built in optional support anyway).
func checkSensorIntersection(sensors []sensor, s sensor) *point {
	// Check all diagonals of the area of the sensor
	for i := int64(0); i <= s.dist+1; i++ {
		xOffset := i
		yOffset := s.dist+1 - i
		xPos := s.pos.x + xOffset
		yPos := s.pos.y + yOffset
		if xPos < 0 || xPos > maxBeaconCoord || yPos < 0 || yPos > maxBeaconCoord {
			continue
		}

		if !checkIntersection(sensors, xPos, yPos) {
			return &point{xPos, yPos}
		}
	}

	for i := int64(0); i <= s.dist+1; i++ {
		xOffset := i
		yOffset := i - s.dist+1
		xPos := s.pos.x + xOffset
		yPos := s.pos.y + yOffset
		if xPos < 0 || xPos > maxBeaconCoord || yPos < 0 || yPos > maxBeaconCoord {
			continue
		}

		if !checkIntersection(sensors, xPos, yPos) {
			return &point{xPos, yPos}
		}
	}

	for i := int64(0); i <= s.dist+1; i++ {
		xOffset := s.dist+1 - i
		yOffset := i
		xPos := s.pos.x + xOffset
		yPos := s.pos.y + yOffset
		if xPos < 0 || xPos > maxBeaconCoord || yPos < 0 || yPos > maxBeaconCoord {
			continue
		}

		if !checkIntersection(sensors, xPos, yPos) {
			return &point{xPos, yPos}
		}
	}

	for i := int64(0); i <= s.dist+1; i++ {
		xOffset := i - s.dist+1
		yOffset := i
		xPos := s.pos.x + xOffset
		yPos := s.pos.y + yOffset
		if xPos < 0 || xPos > maxBeaconCoord || yPos < 0 || yPos > maxBeaconCoord {
			continue
		}

		if !checkIntersection(sensors, xPos, yPos) {
			return &point{xPos, yPos}
		}
	}

	return nil
}

func checkIntersection(sensors []sensor, posX int64, posY int64) bool {
	for _, s := range sensors {
		if abs(s.pos.x - posX) + abs(s.pos.y - posY) <= s.dist {
			return true
		}
	}

	return false
}