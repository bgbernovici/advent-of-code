package day4

/*
   https://adventofcode.com/2023/day/4
   Bogdan Bernovici
*/

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	id         int
	winNumbers []int
	numbers    []int
	score      int
}

func Execute() {
	file, err := os.Open("../Inputs/2023/Day4_.txt")
	if err != nil {
		fmt.Println("Error opening input file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	cards := []Card{}
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		card := Card{winNumbers: []int{}, numbers: []int{}, score: 0}
		firstSplit := strings.Split(line, ":")
		reNumbers := regexp.MustCompile(`(\d+)`)
		idMatches := reNumbers.FindAllStringSubmatch(firstSplit[0], -1)
		secondSplit := strings.Split(firstSplit[1], "|")
		winNumMatches := reNumbers.FindAllStringSubmatch(secondSplit[0], -1)
		numMatches := reNumbers.FindAllStringSubmatch(secondSplit[1], -1)
		for _, match := range idMatches {
			card.id, err = strconv.Atoi(match[1])
		}
		for _, match := range winNumMatches {
			num, _ := strconv.Atoi(match[1])
			card.winNumbers = append(card.winNumbers, num)
		}
		for _, match := range numMatches {
			num, _ := strconv.Atoi(match[1])
			card.numbers = append(card.numbers, num)
			for _, winNum := range card.winNumbers {
				if num == winNum {
					if card.score == 0 {
						card.score = 1
					} else {
						card.score *= 2
					}
				}
			}
		}
		cards = append(cards, card)
		sum += card.score
	}

	counter := 0
	for i := range cards {
		goThroughCards(i, cards, &counter)
	}

	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", counter)
}

func goThroughCards(i int, cards []Card, count *int) {
	matchingNumbers := int(math.Log2(float64(cards[i].score))) + 1
	*count += 1
	for j := i + 1; j < i+matchingNumbers+1; j++ {
		goThroughCards(j, cards, count)
	}
}
