package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"sort"
)

func main(){
    scanner := bufio.NewScanner(os.Stdin)
    topCalories := [3]int64{};
	currentCalories := int64(0)

	for scanner.Scan(){
		text := scanner.Text()

		if len(text) != 0 {
			value, _ := strconv.ParseInt(text, 10, 0)
			currentCalories += value
		} else {
			for index, value := range topCalories{
				if currentCalories > value {
					topCalories[index] = currentCalories
					sort.Slice(topCalories[:], func(i, j int) bool { return topCalories[i] < topCalories[j] })
					break;
				}
			}
			currentCalories = 0
		}
	}
	// Check last result
	for index, value := range topCalories{
		if currentCalories > value {
			topCalories[index] = currentCalories
			break;
		}
	}

	maxCalories := int64(0)
	for _, value := range topCalories{
		maxCalories += value
	}

    fmt.Println(maxCalories)
}