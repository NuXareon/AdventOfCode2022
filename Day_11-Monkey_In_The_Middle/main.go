package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	itemWorry []uint64
	operation func(uint64)uint64
	conditionVal uint64
	monkeyTrue, monkeyFalse int64
	inspectedItems uint64
}

func (m *monkey) test(i uint64) int64 {
	if i%m.conditionVal == 0 {
		return m.monkeyTrue
	}
	return m.monkeyFalse
}

func main () {
	scanner := bufio.NewScanner(os.Stdin)
	monkeys := make([]monkey, 0)

	for scanner.Scan() {
    	line := scanner.Text()

		if strings.HasPrefix(line, "Monkey ") {
			monkeys = append(monkeys, monkey{})
		} else if strings.HasPrefix(line, "  Starting items: ") {
			processStartingItems(monkeys, line)
		} else if strings.HasPrefix(line, "  Operation: new = old ") {
			processOperation(monkeys, line)
		} else if strings.HasPrefix(line, "  Test: divisible by ") {
			processTest(monkeys, line)
		} else if strings.HasPrefix(line,"    If true: throw to monkey ") {
			processMonkeyTrue(monkeys, line)
		} else if strings.HasPrefix(line,"    If false: throw to monkey ") {
			processMonkeyFalse(monkeys, line)
		}
	}

	maxWorry := uint64(1) 
	for _, monkey := range monkeys {
		maxWorry *= monkey.conditionVal
	}
	
	for i := 0; i < 10000; i++ {
		for i := range monkeys {
			monkeyRef := &monkeys[i]
			for _, item := range monkeyRef.itemWorry {
				item = monkeyRef.operation(item)%maxWorry
				targetMonkey := monkeyRef.test(item)
				monkeys[targetMonkey].itemWorry = append(monkeys[targetMonkey].itemWorry, item)
				monkeyRef.inspectedItems++
			}
			monkeyRef.itemWorry = nil
		}
	}

	inspectedItems := make([]uint64,0)
	for _, monkeyDude := range monkeys {
		inspectedItems = append(inspectedItems, monkeyDude.inspectedItems)
	}
	
	sort.Slice(inspectedItems, func(i, j int) bool {
		return inspectedItems[i] > inspectedItems[j]
	})

	ans := inspectedItems[0] * inspectedItems[1]

	fmt.Println(ans)
}

func processStartingItems(monkeys []monkey, line string) {
	line = strings.TrimPrefix(line, "  Starting items: ")
	items := strings.Split(line, ", ")
	for _, item := range items {
		worry, _ := strconv.ParseUint(item, 10, 0)
		monkeys[len(monkeys)-1].itemWorry = append(monkeys[len(monkeys)-1].itemWorry, worry)
	}
}

func processOperation(monkeys []monkey, line string) {
	line = strings.TrimPrefix(line, "  Operation: new = old ")
	if line[0] == '+' {
		line = strings.TrimPrefix(line, "+ ")
		if line == "old" {
			monkeys[len(monkeys)-1].operation = func(i uint64) uint64 {
				return i + i
			}
		} else {
			worry, _ := strconv.ParseUint(line, 10, 0)
			monkeys[len(monkeys)-1].operation = func(i uint64) uint64 {
				return i + worry
			}
		}
	} else if line[0] == '*' {
		line = strings.TrimPrefix(line, "* ")
		if line == "old" {
			monkeys[len(monkeys)-1].operation = func(i uint64) uint64 {
				return i * i
			}
		} else {
			worry, _ := strconv.ParseUint(line, 10, 0)
			monkeys[len(monkeys)-1].operation = func(i uint64) uint64 {
				return i * worry
			}
		}
	}
}

func processTest(monkeys []monkey, line string) {
	line = strings.TrimPrefix(line, "  Test: divisible by ")
	conditionVal, _ := strconv.ParseUint(line, 10, 0)
	monkeys[len(monkeys)-1].conditionVal = conditionVal
}

func processMonkeyTrue(monkeys []monkey, line string) {
	line = strings.TrimPrefix(line, "    If true: throw to monkey ")
	monkey, _ := strconv.ParseInt(line, 10, 0)
	monkeys[len(monkeys)-1].monkeyTrue = monkey
}

func processMonkeyFalse(monkeys []monkey, line string) {
	line = strings.TrimPrefix(line, "    If false: throw to monkey ")
	monkey, _ := strconv.ParseInt(line, 10, 0)
	monkeys[len(monkeys)-1].monkeyFalse = monkey
}