package day12_uncacheable

/*
   https://adventofcode.com/2023/day/12
   Bogdan Bernovici
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Execute() {
	start := time.Now()
	file, err := os.Open("../Inputs/2023/Day12.txt")
	if err != nil {
		fmt.Println("Error opening input file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	criterias := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		spring := strings.Split(line, " ")[0] + "."
		criteria := strings.Split(strings.Split(line, " ")[1], ",")
		tempCriterias := []int{}
		for _, r := range criteria {
			cr, _ := strconv.Atoi(r)
			tempCriterias = append(tempCriterias, cr)
		}
		lines = append(lines, spring)
		criterias = append(criterias, tempCriterias)

	}
	arrangements := 0
	for i, line := range lines {
		traverse(rune(line[0]), line, 0, criterias[i], &arrangements, 0, "")
	}
	fmt.Println("Part 1: ", arrangements)

	unfoldedLines := []string{}
	unfoldedCriterias := [][]int{}
	for _, line := range lines {
		line = line[:len(line)-1] // remove the appended dot from the end
		factor := line
		for i := 0; i < 4; i++ {
			line += "?" + factor
		}
		line += "."
		unfoldedLines = append(unfoldedLines, line)
	}
	for _, criteria := range criterias {
		factor := criteria
		for i := 0; i < 4; i++ {
			criteria = append(criteria, factor...)
		}
		unfoldedCriterias = append(unfoldedCriterias, criteria)
	}

	arrangements = 0
	for i, line := range unfoldedLines {
		traverse(rune(line[0]), line, 0, unfoldedCriterias[i], &arrangements, 0, "")
	}
	fmt.Println("Part 2: ", arrangements)

	elapsed := time.Since(start)
	fmt.Println("Total execution took ", elapsed)
}

func traverse(spring rune, line string, i int, criteria []int, arrangements *int, groupSpan int, constructedString string) {
	clonedCriteria := make([]int, len(criteria))
	copy(clonedCriteria, criteria)
	for p := i; p < len(line); p++ {

		// constructedString was used for debugging purposes
		if line[p] != '?' {
			constructedString += string(line[p])
		}
		if p == i && spring != '?' {
			constructedString += string(spring)
		}

		if (p == i && spring == '#') || line[p] == '#' {
			groupSpan += 1
		} else if (p == i && spring == '.') || line[p] == '.' {
			if len(clonedCriteria) > 0 && groupSpan == clonedCriteria[0] {
				copy(clonedCriteria[0:], clonedCriteria[1:])
				clonedCriteria = clonedCriteria[:len(clonedCriteria)-1]
				groupSpan = 0
			} else if groupSpan != 0 {
				break // abandon arrangement if we are in a group and it doesn't match the criteria
			}
		} else if line[p] == '?' {
			traverse('#', line, p, clonedCriteria, arrangements, groupSpan, constructedString)
			traverse('.', line, p, clonedCriteria, arrangements, groupSpan, constructedString)
			break
		}

		if len(line)-1 == p && groupSpan == 0 && len(clonedCriteria) == 0 {
			*arrangements += 1
		}
	}
}
