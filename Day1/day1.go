package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func run() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}
}
