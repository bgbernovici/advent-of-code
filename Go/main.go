package main

import (
	"fmt"

	"github.com/bgbernovici/advent-of-code/2023/day1"
	"github.com/bgbernovici/advent-of-code/2023/day2"
	"github.com/bgbernovici/advent-of-code/2023/day3"
	"github.com/bgbernovici/advent-of-code/2023/day4"
	"github.com/bgbernovici/advent-of-code/2023/day5"
	"github.com/bgbernovici/advent-of-code/2023/day5_opti"
	"github.com/bgbernovici/advent-of-code/2023/day6"
	"github.com/bgbernovici/advent-of-code/2023/day7"
)

func main() {
	fmt.Println("## DAY 1 ##")
	day1.Execute()
	fmt.Println("## DAY 2 ##")
	day2.Execute()
	fmt.Println("## DAY 3 ##")
	day3.Execute()
	fmt.Println("## DAY 4 ##")
	day4.Execute()
	fmt.Println("## DAY 5 NAIVE ##")
	day5.Execute()
	fmt.Println("## DAY 5 OPTIMIZED ##")
	day5_opti.Execute()
	fmt.Println("## DAY 6 ##")
	day6.Execute()
	fmt.Println("## DAY 7 ##")
	day7.Execute()
}
