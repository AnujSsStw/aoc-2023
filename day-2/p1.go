package day02

import (
	day01 "aoc2023/day-1"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	readFile, err := os.Open("/home/ani/personal/aoc-2023/day-02/input.txt")
	day01.Check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	// only 12 red cubes, 13 green cubes, and 14 blue cubes
	ans := map[string]int{}
	gamesum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		game := strings.Split(line, ":")
		sets := strings.Split(game[1], ";")
		mark := 0
		for _, e := range sets {
			setInfo := strings.Split(e, ",")
			for _, e := range setInfo {
				color := strings.Split(e, " ")
				num, _ := strconv.Atoi(color[1])
				ans[color[2]] += num
				for k, v := range ans {
					if k == "blue" && v <= 14 {
						continue
					} else if k == "red" && v <= 12 {
						continue
					} else if k == "green" && v <= 13 {
						continue
					} else {
						mark = 1
					}
				}
				clear(ans)
			}
		}
		if mark == 0 {
			num, _ := strconv.Atoi(strings.Split(game[0], " ")[1])
			gamesum += num
			fmt.Println(num)
		}

	}
	fmt.Println("ans is ::", gamesum)

}
