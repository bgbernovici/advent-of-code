package day9

/*
   https://adventofcode.com/2023/day/9
   Bogdan Bernovici
*/

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

type History struct {
	values               []int64
	placeholder          int64
	backwardsPlaceholder int64
}

func Execute() {
	start := time.Now()
	file, err := os.Open("../Inputs/2023/Day9_.txt")
	if err != nil {
		fmt.Println("Error opening input file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	histories := []History{}
	for scanner.Scan() {

		line := scanner.Text()
		lines = append(lines, line)

		regex := regexp.MustCompile(`(-?\d+)`)
		matches := regex.FindAllStringSubmatch(line, -1)
		history := History{}
		for _, match := range matches {
			v, _ := strconv.ParseInt(match[1], 10, 64)
			history.values = append(history.values, v)
		}
		histories = append(histories, history)
	}

	for k, history := range histories {
		tempHistory := []History{history}

		// Compute differences
		ok := true
		for ok {
			ok = false
			h := History{values: []int64{}}
			th := tempHistory[len(tempHistory)-1].values
			for j := 1; j < len(th); j++ {
				diff := th[j] - th[j-1]
				if diff != 0 {
					ok = true
				}
				h.values = append(h.values, diff)
			}
			tempHistory = append(tempHistory, h)
		}

		// Compute placeholders
		for i := len(tempHistory) - 1; i >= 0; i-- {
			if i == len(tempHistory)-1 {
				tempHistory[i].placeholder = 0
				tempHistory[i].backwardsPlaceholder = 0
				continue
			}
			tempHistory[i].placeholder = tempHistory[i+1].placeholder + tempHistory[i].values[len(tempHistory[i].values)-1]
			tempHistory[i].backwardsPlaceholder = tempHistory[i].values[0] - tempHistory[i+1].backwardsPlaceholder
		}

		histories[k] = tempHistory[0]
	}

	sum := int64(0)
	for _, history := range histories {
		sum += history.placeholder
	}
	fmt.Println("Part 1: ", sum)

	sum = int64(0)
	for _, history := range histories {
		sum += history.backwardsPlaceholder
	}
	fmt.Println("Part 2: ", sum)
	elapsed := time.Since(start)
	fmt.Println("Total execution took ", elapsed)
}
