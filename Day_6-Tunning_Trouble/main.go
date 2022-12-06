package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	
	ans := findUniqueSequence(input, 14)
	fmt.Println(ans)
}

func findUniqueSequence(input string, length int) int {
	segmentStart := 0
	for i, val := range input {
		if i == segmentStart {
			continue
		}

		foundIdx := strings.LastIndex(input[segmentStart:i], string(val))

		if (foundIdx == -1) {
			if (i - segmentStart == length-1) {
				return i + 1
			}
		} else {
			segmentStart += foundIdx+1
		}
	}

	return -1
}