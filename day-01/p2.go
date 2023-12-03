package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numArr []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func Part2() {
	readFile, err := os.Open("/home/ani/personal/aoc-2023/day-01/input.txt")
	Check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var ans []int

		strArr := strings.Split(fileScanner.Text(), "")
		for i, e := range line {
			if num, err := strconv.Atoi(string(e)); err == nil {
				ans = append(ans, num)
			}

			for j, e := range numArr {
				dummy := strings.Join(strArr[i:], "")
				if strings.HasPrefix(dummy, e) {
					ans = append(ans, j+1)
				}

			}
		}
		sum += ans[0]*10 + ans[len(ans)-1]

	}
	fmt.Println(sum)

}
