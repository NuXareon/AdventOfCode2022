package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type comparsionResult int

const (
	CompEqual comparsionResult = iota
	CompFirst
	CompSecond
)

// Why does go not have a generic math.min function??
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	ans := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
    	packet := scanner.Text()
		if (len(packet) == 0) {
			continue
		}

		ans, _ = insertOrdered(ans, packet)
	}

	x, y := 0, 0
	ans, x = insertOrdered(ans, "[[2]]")
	_, y = insertOrdered(ans, "[[6]]")

	fmt.Println((x+1)*(y+1))
}

func insertOrdered(ans []string, packet string) ([]string, int) {
	packetSplit := splitList(packet)

	// This sorting could be much more efficient
	for i, val := range ans {	
		valSplit := splitList(val)
		compResult := processListSplit(packetSplit, valSplit)
		if compResult == CompFirst {
			ans = append(ans[:i+1], ans[i:]...)
			ans[i] = packet
			return ans, i
		}
	}

	ans = append(ans, packet)
	return ans, len(ans)-1
}

func splitList(list string) []string {
	if len(list) <= 2 {
		// [] is an empty list
		return make([]string, 0)
	}

	listLevel := 0
	segmentStart := 1
	split := make([]string, 0)
	for i, val := range list {
		if i == 0 || i == len(list) - 1 {
			continue
		}

		if val == '[' {
			listLevel++
		} else if val == ']' {
			listLevel--
		} else if val == ',' {
			if listLevel > 0 {
				// This comma is part of a sublist, skip
				continue
			}

			if (segmentStart == i) {
				// I don't think this is a valid input anyway
				split = append(split, "")
			} else {
				split = append(split, list[segmentStart:i])
			}
			segmentStart = i+1
		}
	}
	split = append(split, list[segmentStart:len(list)-1])

	return split
}

func processListSplit(first []string, second []string) (comparsionResult) {
	minLen := min(len(first), len(second))

	for i := 0; i < minLen; i++ {
		valFirst := first[i]
		valSecond := second[i]
		var firstValueParse int64
		var secondValueParse int64
		var firstSplit []string
		var secondSplit []string

		if valFirst[0] >= '0' && valFirst[0] <= '9' {
			firstValueParse, _ = strconv.ParseInt(valFirst, 10, 0)
		} else if valFirst[0] == '[' {
			firstSplit = splitList(valFirst)
		}

		if valSecond[0] >= '0' && valSecond[0] <= '9' {
			secondValueParse, _ = strconv.ParseInt(valSecond, 10, 0)
		} else if valSecond[0] == '[' {
			secondSplit = splitList(valSecond)
		}

		// At least one side is a list
		if firstSplit != nil || secondSplit != nil {
			// Convert mixed types into lists
			if firstSplit == nil {
				firstSplit = append(firstSplit, valFirst)
			}
			if secondSplit == nil {
				secondSplit = append(secondSplit, valSecond)
			}

			compResult := processListSplit(firstSplit, secondSplit)

			if compResult != CompEqual {
				return compResult
			}
		} else { // Both sides have numbers
			if firstValueParse > secondValueParse {
				return CompSecond
			} else if firstValueParse < secondValueParse {
				return CompFirst
			}
		}
	}

	// All values are equal, compare sizes
	if len(first) == len(second) {
		return CompEqual
	} else if len(first) < len(second) {
		return CompFirst
	}
	return CompSecond
}