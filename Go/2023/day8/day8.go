package day8

/*
   https://adventofcode.com/2023/day/8
   Bogdan Bernovici
*/

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
	"time"
)

type Node struct {
	v string
	l string
	r string
}

func Execute() {
	start := time.Now()
	file, err := os.Open("../Inputs/2023/Day8_.txt")
	if err != nil {
		fmt.Println("Error opening input file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	network := make(map[string]Node)
	lineCount := 0
	instructions := ""
	for scanner.Scan() {

		line := scanner.Text()
		lines = append(lines, line)

		if lineCount == 0 {
			instructions = line
		}
		lineCount += 1
		regex := regexp.MustCompile("[A-Z]{3}")
		matches := regex.FindAllStringSubmatch(line, -1)
		for i, match := range matches {
			if i == 0 {
				network[match[0]] = Node{v: match[0]}
			} else if i == 1 {
				n := network[matches[0][0]]
				n.l = strings.TrimSpace(match[0])
				network[matches[0][0]] = n
			} else if i == 2 {
				n := network[matches[0][0]]
				n.r = strings.TrimSpace(match[0])
				network[matches[0][0]] = n
			}
		}
	}

	steps := countSteps(instructions, network["AAA"], network["ZZZ"], &network)
	fmt.Println("Part 1: ", steps)

	lcm := 0
	for _, nodeX := range network {
		if nodeX.v[2] == 'A' {
			steps := countSteps(instructions, nodeX, Node{v: "Z"}, &network)
			if lcm == 0 {
				lcm = steps
			} else {
				lcm = compute_lcm(lcm, steps)
			}
		}

	}
	fmt.Println("Part 2: ", lcm)
	elapsed := time.Since(start)
	fmt.Println("Total execution took ", elapsed)
}

func countSteps(instructions string, startNode Node, endNode Node, network *map[string]Node) int {
	currentNode := startNode
	instIndex := 0
	steps := 0
	for {
		if instructions[instIndex] == 'L' {
			currentNode = (*network)[currentNode.l]
		} else if instructions[instIndex] == 'R' {
			currentNode = (*network)[currentNode.r]
		}
		steps++
		if len(endNode.v) == 3 {
			if currentNode.v == endNode.v && instIndex == len(instructions)-1 {
				break
			}
		} else {
			if currentNode.v[2:] == endNode.v && instIndex == len(instructions)-1 {
				break
			}
		}
		if instIndex+1 <= len(instructions)-1 {
			instIndex += 1
		} else {
			instIndex = 0
		}
	}
	return steps
}

func compute_lcm(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(compute_gcd(a, b)))
}

func compute_gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
