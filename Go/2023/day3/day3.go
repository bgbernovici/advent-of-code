package day3

/*
   https://adventofcode.com/2023/day/3
   Bogdan Bernovici
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type PartNumber struct {
	value  string
	start  int
	end    int
	lineNo int
	valid  bool
}

type Symbol struct {
	value    string
	index    int
	lineNo   int
	attached []int
}

func Execute() {
	file, err := os.Open("../Inputs/2023/Day3_.txt")
	if err != nil {
		fmt.Println("Error opening input file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	lineNo := 0
	partNumbers := []PartNumber{}
	symbols := []Symbol{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		partNumber := PartNumber{value: "", valid: false, lineNo: lineNo}
		for index, c := range line {
			if unicode.IsDigit(rune(c)) {
				partNumber.value += string(c)
				if len(partNumber.value) == 1 {
					partNumber.start = index
				}
				if index == len(line)-1 && len(partNumber.value) > 0 {
					partNumber.end = index
					partNumbers = append(partNumbers, partNumber)
					partNumber = PartNumber{value: "", valid: false, lineNo: lineNo}
				}
			} else {
				if len(partNumber.value) > 0 {
					if index == len(line)-1 && len(partNumber.value) > 0 {
						partNumber.end = index
					} else {
						partNumber.end = index - 1
					}
					partNumbers = append(partNumbers, partNumber)
					partNumber = PartNumber{value: "", valid: false, lineNo: lineNo}
				}
			}
			if !unicode.IsDigit(rune(c)) && rune(c) != '.' {
				s := Symbol{value: string(c), index: index, lineNo: lineNo, attached: []int{}}
				symbols = append(symbols, s)
			}
		}
		lineNo += 1
	}

	for i, s := range symbols {
		for j, pn := range partNumbers {
			if pn.lineNo <= s.lineNo+1 && pn.lineNo >= s.lineNo-1 {
				if (pn.start >= s.index-1 && pn.start <= s.index+1) || (pn.end >= s.index-1 && pn.end <= s.index+1) {
					partNumbers[j].valid = true
					v, _ := strconv.Atoi(pn.value)
					symbols[i].attached = append(symbols[i].attached, v)
				}
			}
		}
	}

	sumParts := 0
	sumGears := 0
	for _, s := range symbols {
		ratio := 0
		if len(s.attached) > 1 {
			ratio = 1
		}
		for _, v := range s.attached {
			sumParts += v
			if len(s.attached) > 1 {
				ratio *= v
			}
		}
		sumGears += ratio
	}
	fmt.Println("Part 1: ", sumParts)
	fmt.Println("Part 2: ", sumGears)
}
