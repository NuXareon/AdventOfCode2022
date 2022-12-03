package main

import (
    "fmt"
    "bufio"
    "os"
)

func main(){
    scanner := bufio.NewScanner(os.Stdin)
	priority := int64(0)

    for scanner.Scan() {
        rucksack1 := scanner.Text()
		itemsInRucksack1 := processItemsInRucksack(rucksack1)

		scanner.Scan()
        rucksack2 := scanner.Text()
		itemsInRucksack2 := processItemsInRucksack(rucksack2)

		scanner.Scan()
        rucksack3 := scanner.Text()
		for _, item := range rucksack3 {
			_, found1 := itemsInRucksack1[item]
			_, found2 := itemsInRucksack2[item]
			if (found1 && found2) {
				priority += getItemPrio(item)
				break
			}
		}
	}

	fmt.Println(priority)
}

func processItemsInRucksack(rucksack string) map[rune]bool {
	itemsInRucksack := make(map[rune]bool)
	for _, item := range rucksack {
		itemsInRucksack[item] = true
	}
	return itemsInRucksack
}

func getItemPrio(item rune) int64 {
	if item >= 'a' && item <= 'z' {
		return int64(item - 'a') + 1
	} else if item >= 'A' && item <= 'Z' {
		return int64(item - 'A') + 27
	}
	return 0
}