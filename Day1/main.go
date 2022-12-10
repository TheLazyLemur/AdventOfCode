package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	var elvesCount, max int
	elvesMap := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			elvesCount++
			elvesMap[elvesCount] = 0
			continue
		}

		value, _ := strconv.Atoi(line)
		elvesMap[elvesCount] += value

		if elvesMap[elvesCount] > max {
			max = elvesMap[elvesCount]
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(max)
}

