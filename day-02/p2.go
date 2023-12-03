package day02

import (
	day01 "aoc2023/day-01"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	readFile, err := os.Open("/home/ani/personal/aoc-2023/day-02/input.txt")
	day01.Check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	// only 12 red cubes, 13 green cubes, and 14 blue cubes
	ans := map[string]float64{}
	gamePower := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		game := strings.Split(line, ":")
		sets := strings.Split(game[1], ";")
		for _, e := range sets {
			setInfo := strings.Split(e, ",")
			for _, e := range setInfo {
				color := strings.Split(e, " ")
				num, _ := strconv.Atoi(color[1])
				// max
				x := -1.0
				n, ok := ans[color[2]]
				if ok {
					x = n
				}
				ans[color[2]] = math.Max(float64(num), x)
				fmt.Println(ans)
			}
		}
		power := 1
		for _, v := range ans {
			power *= int(v)
		}
		gamePower += power
		clear(ans)

	}
	fmt.Println("ans is ::", gamePower)

}
