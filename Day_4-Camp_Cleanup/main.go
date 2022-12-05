package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	overlaps := 0

    for scanner.Scan() {
    	rangePairStr := scanner.Text()

		rangePairs := strings.Split(rangePairStr, ",")
		firstPairValues := strings.Split(rangePairs[0], "-")
		secondPairValues := strings.Split(rangePairs[1], "-")

		firstPairStart, _:= strconv.ParseInt(firstPairValues[0], 10, 0)
		firstPairEnd, _ := strconv.ParseInt(firstPairValues[1], 10, 0)
		secondPairStart, _ := strconv.ParseInt(secondPairValues[0], 10, 0)
		secondPairEnd, _ := strconv.ParseInt(secondPairValues[1], 10, 0)

		if firstPairStart >= secondPairStart && firstPairStart <= secondPairEnd {
			// Overlap at the start of first
			overlaps++
		} else if secondPairStart >= firstPairStart && secondPairStart <= firstPairEnd {
			// Overlat at the start of second
			overlaps++
		}
	}

	fmt.Println(overlaps)
}