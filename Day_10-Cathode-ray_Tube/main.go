package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type CPUState int

const (
	ReadingInput CPUState = iota
	ProcessingAdd
)

func main () {
	scanner := bufio.NewScanner(os.Stdin)
	x := int64(1)
	ans := int64(0)
	CPUState := ReadingInput
	previousValue := int64(0)

	for cycle := int64(1);; cycle++ {
		// Calculate ans
		if (cycle-20)%40 == 0 {
			ans+=cycle*x
		}

		// Draw CRT
		// Note: It technillay prints one extra pixel at the end, since we don't check for input finished until later
		pixelDrawn := (cycle-1)%40
		if pixelDrawn >= x-1 && pixelDrawn <= x+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if pixelDrawn == 39 {
			fmt.Print("\n")
		}

		// Process instructions
		if CPUState == ReadingInput {
			hasInstruction := scanner.Scan()
			if !hasInstruction {
				break	// we are done with all input
			}

			line := scanner.Text()
			instruction := strings.Split(line, " ")
			if instruction[0] == "addx" {
				previousValue, _ = strconv.ParseInt(instruction[1], 10, 0)
				CPUState = ProcessingAdd
			}
		} else if (CPUState == ProcessingAdd) {
			x += previousValue
			CPUState = ReadingInput
		}
	}
	fmt.Println(ans)
}