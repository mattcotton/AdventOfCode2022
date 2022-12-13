package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func Run() {
	file, err := os.Open("Day 1/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var elves []int
	totalCals := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			elves = append(elves, totalCals)
			totalCals = 0
		} else {
			cals, err := strconv.Atoi(scanner.Text())

			if err == nil {
				totalCals += cals
			} else {
				fmt.Println("Error: ", err)
			}
		}
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	fmt.Println("Part 1: ", elves[0])
	fmt.Println("Part 2: ", elves[0]+elves[1]+elves[2])
}
