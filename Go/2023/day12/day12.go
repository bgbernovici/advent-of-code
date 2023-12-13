package day12

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
	file, err := os.Open("../Inputs/2023/Day12_.txt")
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
		memo := make(Cache)
		arrangements += traverse(line, criterias[i], 0, &memo)
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
		cache := make(Cache)
		arrangements += traverse(line, unfoldedCriterias[i], 0, &cache)
	}
	fmt.Println("Part 2: ", arrangements)

	elapsed := time.Since(start)
	fmt.Println("Total execution took ", elapsed)
}

type Cache map[string]int

func traverse(line string, criteria []int, groupSpan int, cache *Cache) int {

	if val, ok := (*cache)[fmt.Sprintf("%s|%v|%d", line, criteria, groupSpan)]; ok {
		return val
	}

	if len(line) == 0 && groupSpan == 0 && len(criteria) == 0 {
		return 1
	} else if len(line) == 0 {
		return 0
	}

	count := 0

	if line[0] == '#' || line[0] == '?' {
		count += traverse(line[1:], criteria, groupSpan+1, cache)
	}

	if line[0] == '.' || line[0] == '?' {
		if len(criteria) > 0 && groupSpan == criteria[0] {
			count += traverse(line[1:], criteria[1:], 0, cache)
		} else if groupSpan != 0 {
			return count
		} else {
			count += traverse(line[1:], criteria, 0, cache)
		}
	}

	(*cache)[fmt.Sprintf("%s|%v|%d", line, criteria, groupSpan)] = count

	return count
}
