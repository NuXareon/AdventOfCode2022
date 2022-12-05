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
	stacks := readInitialStacks(scanner)

	for scanner.Scan() {
    	line := scanner.Text()

		if (len(line) == 0) {
			continue
		}

		lineSplit := strings.Split(line, " from ")
		numMoveStr := lineSplit[0]
		fromToStr := lineSplit[1]
		numMoveStr = strings.TrimPrefix(numMoveStr, "move ")
		numMove, _ := strconv.ParseInt(numMoveStr, 10, 0)
		fromToSplit := strings.Split(fromToStr, " to ")
		from, _ := strconv.ParseInt(fromToSplit[0], 10, 0)
		to, _ := strconv.ParseInt(fromToSplit[1], 10, 0)
		from--
		to--

		lenFrom := int64(len(stacks[from]))
		firstElemIdx := lenFrom-numMove
		stacks[to] = append(stacks[to], stacks[from][firstElemIdx:lenFrom]...)
		stacks[from] = stacks[from][:firstElemIdx]
	}

	var ans string
	for _, val := range stacks {
		ans += string(val[len(val)-1])
	}

	fmt.Println(ans)
}

/*
[N]         [C]     [Z]            
[Q] [G]     [V]     [S]         [V]
[L] [C]     [M]     [T]     [W] [L]
[S] [H]     [L]     [C] [D] [H] [S]
[C] [V] [F] [D]     [D] [B] [Q] [F]
[Z] [T] [Z] [T] [C] [J] [G] [S] [Q]
[P] [P] [C] [W] [W] [F] [W] [J] [C]
[T] [L] [D] [G] [P] [P] [V] [N] [R]
 1   2   3   4   5   6   7   8   9 
 
 position of containers: 1,5,9,13...
 i = 1+(4*(n-1))
 n = 1+(i-1)/4
 */ 
func readInitialStacks(scanner *bufio.Scanner) [][]rune {
	initInput := make([]string, 0)
	for scanner.Scan() {
    	line := scanner.Text()

		finishedInput := false
		for _, val := range line {
			if val >= '1' && val <= '9' {
				// If we find a number we are done reading the stacks.
				finishedInput = true
				break; 
			} else if val == '[' {
				break;
			}
		}

		if (finishedInput) {
			// Finished reading input. Parse it and return. 
			return parseInitialInput(initInput)
		} else {
			// Store input data
			initInput = append(initInput, line)
		}
	}

	// Failed to read a valid input.
	return make([][]rune, 0)
}

func parseInitialInput(initInput [] string) [][]rune {
	stacks := make([][]rune, 0)
	for inputIdx := len(initInput)-1; inputIdx >= 0; inputIdx-- {
		for i, val := range initInput[inputIdx] {
			if val >= 'A' && val <= 'Z' {
				// Note: Containers start at 0
				stackIdx := (i-1)/4 

				// Grow the number of stacks if needed
				for len(stacks) <= stackIdx {
					stacks = append(stacks, make([]rune, 0))
				}

				// Insert new container
				stacks[stackIdx] = append(stacks[stackIdx], val)
			}
		}
	}
	return stacks
}