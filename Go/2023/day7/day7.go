package day7

/*
   https://adventofcode.com/2023/day/7
   Bogdan Bernovici
*/

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var cardRanks = map[rune]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

type Hand struct {
	cards []rune
	bid   int
	kind  int
}

type CardFunc func(h Hand) bool

var priorityFuncs = []CardFunc{
	func(h Hand) bool {
		ok := true
		for i := 0; i < len(h.cards)-1; i++ {
			if h.cards[i] != h.cards[i+1] {
				ok = false
			}
		}
		return ok
	},
	func(h Hand) bool {
		counts := countOccurrences(h.cards)
		ok := false
		for _, count := range counts {
			if count == 4 {
				ok = true
			}
		}
		return ok
	},
	func(h Hand) bool {
		counts := countOccurrences(h.cards)
		ok2 := false
		ok3 := false
		for _, count := range counts {
			if count == 3 {
				ok3 = true
			} else if count == 2 {
				ok2 = true
			}
		}
		return ok2 && ok3
	},
	func(h Hand) bool {
		return h.pair(3)
	},
	func(h Hand) bool {
		return h.twoPairs()
	},
	func(h Hand) bool {
		return h.pair(2)
	},
	func(h Hand) bool {
		return h.pair(1)
	},
}

func (h Hand) twoPairs() bool {
	counts := countOccurrences(h.cards)
	hasTwo := false
	hasAnotherTwo := false
	for _, count := range counts {
		if count == 2 {
			if hasTwo {
				hasAnotherTwo = true
			}
			hasTwo = true
		}
	}
	return hasTwo && hasAnotherTwo
}

func (h Hand) pair(numInPair int) bool {
	counts := countOccurrences(h.cards)
	ok := false
	for _, count := range counts {
		if count == numInPair {
			ok = true
			break
		}
	}
	return ok
}

func countOccurrences(cards []rune) map[rune]int {
	counts := make(map[rune]int)
	for _, card := range cards {
		counts[card]++
	}
	return counts
}

func Execute() {
	file, err := os.Open("../Inputs/2023/Day7_.txt")
	if err != nil {
		fmt.Println("Error opening input file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	hands := []Hand{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		lineSplit := strings.Split(line, " ")
		hand := Hand{cards: []rune{}}
		for i := 0; i < len(lineSplit[0]); i++ {
			hand.cards = append(hand.cards, rune(lineSplit[0][i]))
		}
		bid, _ := strconv.Atoi(lineSplit[1])
		hand.bid = bid
		hands = append(hands, hand)
	}

	for i, hand := range hands {
		for j, f := range priorityFuncs {
			if f(hand) == true {
				hands[i].kind = j
				break
			}
		}
	}

	inplaceSortSlice(&hands)

	sum := 0
	for i, hand := range hands {
		sum += hand.bid * (i + 1)
	}
	fmt.Println("Part 1: ", sum)

	cardRanks['J'] = 0

	for i, hand := range hands {
		tempHands := []Hand{}
		regex := regexp.MustCompile("[A-Z0-9]")
		matches := regex.FindAllString(string(hand.cards), -1)
		uniqueLetters := make(map[string]bool)
		for _, match := range matches {
			uniqueLetters[match] = true
		}
		for key := range uniqueLetters {
			newCards := strings.ReplaceAll(string(hand.cards), "J", key)
			h := Hand{bid: hand.bid, cards: []rune(newCards)}
			tempHands = append(tempHands, h)
		}
		minKind := len(priorityFuncs) - 1
		for _, h := range tempHands {
			for k, f := range priorityFuncs {
				if f(h) == true {
					if k < minKind {
						minKind = k
					}
					break
				}
			}
		}
		hands[i].kind = minKind
	}

	inplaceSortSlice(&hands)

	sum = 0
	for i, hand := range hands {
		sum += hand.bid * (i + 1)
	}
	fmt.Println("Part 2: ", sum)
}

func inplaceSortSlice(hands *[]Hand) {
	sort.Slice(*hands, func(i, j int) bool {
		var pred1 = (*hands)[i].kind > (*hands)[j].kind
		if pred1 {
			return pred1
		}
		var pred2 = (*hands)[i].kind == (*hands)[j].kind
		if pred2 {
			for k := 0; k < 5; k++ {
				if cardRanks[(*hands)[i].cards[k]] == cardRanks[(*hands)[j].cards[k]] {
					continue
				} else {
					return cardRanks[(*hands)[i].cards[k]] < cardRanks[(*hands)[j].cards[k]]
				}
			}
		}
		return pred2
	})
}
