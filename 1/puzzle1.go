package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	elves := make([]int, 0, 10)
	currentElf := 0
	elves = append(elves, 0)
	for fileScanner.Scan() {
		num, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			currentElf += 1
		} else {
			elves = append(elves, 0)
			elves[currentElf] += num
		}
	}
	sort.Ints(elves)
	top3 := elves[len(elves)-3:]
	sumTop3 := 0
	for _, v := range top3 {
		sumTop3 += v
	}
	fmt.Println("Top 1:", elves[len(elves)-1])
	fmt.Println("Top 3:", sumTop3)
}
