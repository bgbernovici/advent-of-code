package day11_dijkstra

/*
   https://adventofcode.com/2023/day/11
   Bogdan Bernovici
*/

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sync"
	"time"
)

var moves = [][2]int{
	{0, 1},  //up
	{0, -1}, //down
	{-1, 0}, //left
	{1, 0},  //right
}

func Execute() {
	start := time.Now()
	file, err := os.Open("../Inputs/2023/Day11.txt")
	if err != nil {
		fmt.Println("Error opening input file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	m := [][]rune{}

	for scanner.Scan() {

		line := scanner.Text()
		lines = append(lines, line)

		l := []rune{}
		for _, tile := range line {
			l = append(l, tile)
		}
		m = append(m, l)
	}

	tempM := [][]rune{}
	for i := range m {
		tempM = append(tempM, m[i])
		allDots := true
		for _, c := range m[i] {
			if c != '.' {
				allDots = false
			}
		}
		if allDots {
			tempM = append(tempM, m[i])
		}
	}
	m = make([][]rune, len(tempM))
	copy(m, tempM)
	extended := 0
	for i := range m[0] {
		allDots := true
		for j := range m {
			if m[j][i] != '.' {
				allDots = false
			}
		}
		if allDots {
			for j := range m {
				newRow := make([]rune, len(tempM[j])+1)
				copy(newRow, tempM[j][:i+extended])
				newRow[i+extended] = '.'
				copy(newRow[i+1+extended:], tempM[j][i+extended:])
				tempM[j] = newRow
			}
			extended += 1
		}
	}
	m = tempM
	galaxies := [][2]int{}
	for i := range m {
		for j := range m[i] {
			if tempM[i][j] == '#' {
				galaxies = append(galaxies, [2]int{i, j})
			}
		}
	}

	sum := 0
	sumCh := make(chan int)
	var wg sync.WaitGroup
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			wg.Add(1)
			go dijkstra(&m, [2]int{galaxies[i][0], galaxies[i][1]}, [2]int{galaxies[j][0], galaxies[j][1]}, sumCh, &wg)
		}
	}

	go func() {
		wg.Wait()
		close(sumCh)
	}()

	for s := range sumCh {
		sum += s
	}
	fmt.Println("Part 1: ", sum)

	elapsed := time.Since(start)
	fmt.Println("Total execution took ", elapsed)
}

func dijkstra(m *[][]rune, source [2]int, dest [2]int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	dist := map[[2]int]int{}
	prev := map[[2]int][2]int{}
	queue := [][2]int{}
	for i := range *m {
		for j := range (*m)[i] {
			dist[[2]int{i, j}] = math.MaxInt
			queue = append(queue, [2]int{i, j})
		}
	}
	dist[[2]int{source[0], source[1]}] = 0

	for len(queue) > 0 {
		minDist := math.MaxInt
		var minVertex [2]int
		for key, value := range dist {
			deleted := true
			for _, v := range queue {
				if v[0] == key[0] && v[1] == key[1] {
					deleted = false
				}
			}
			if value <= minDist && !deleted {
				minDist = value
				minVertex = key
			}
		}
		var u [2]int
		for i, v := range queue {
			if v[0] == minVertex[0] && v[1] == minVertex[1] {
				u = v
				copy(queue[i:], queue[i+1:])
				queue = queue[:len(queue)-1]
				break
			}
		}
		if u[0] == dest[0] && u[1] == dest[1] {
			break
		}
		for _, move := range moves {
			alt := dist[[2]int{u[0], u[1]}] + 1

			if alt < dist[[2]int{u[0] + move[0], u[1] + move[1]}] {
				dist[[2]int{u[0] + move[0], u[1] + move[1]}] = alt
				prev[[2]int{u[0] + move[0], u[1] + move[1]}] = [2]int{u[0], u[1]}
			}
		}
	}
	ch <- dist[[2]int{dest[0], dest[1]}]
}
