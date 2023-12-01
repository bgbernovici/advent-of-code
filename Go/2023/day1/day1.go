package day1

/*
   https://adventofcode.com/2023/day/1
   Bogdan Bernovici
*/

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Execute() {
	file, err := os.Open("../Inputs/2023/Day1_.txt")
	if err != nil {
		fmt.Println("Error opening input file: ", err)
	}
	defer file.Close()
	sum := solve(file, false)
	println("Part 1: ", sum)

	file, err = os.Open("../Inputs/2023/Day1_.txt")
	sum = solve(file, true)
	println("Part 2: ", sum)

}

func solve(file *os.File, needsRecalibration bool) int {
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if needsRecalibration {
			line = recalibrate(line)
		}
		value, err := calibrate(line)
		if err == nil {
			sum += value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file: ", err)
	}
	return sum
}

func calibrate(line string) (int, error) {
	re := regexp.MustCompile(`[0-9]`)
	digits := re.FindAllString(line, -1)
	if len(digits) == 0 {
		return 0, errors.New("No digits in string")
	}
	f, _ := strconv.Atoi(digits[0])
	l, _ := strconv.Atoi(digits[len(digits)-1])
	return f*10 + l, nil
}

var domain = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func recalibrate(line string) string {
	var newLine = ""
	for k, v := range domain {
		line = strings.Replace(line, fmt.Sprint(v), k, -1)
	}
	for i := 0; i < len(line); i++ {
		for k, v := range domain {
			if strings.Index(line[i:minInt(i+len(k), len(line))], k) != -1 {
				newLine += fmt.Sprint(v)
				break
			}
		}
	}
	return newLine
}
