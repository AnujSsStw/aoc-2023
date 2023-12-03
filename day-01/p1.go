package day01

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Part1() {

	readFile, err := os.Open("/home/ani/personal/aoc-2023/day-01/input.txt")
	Check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		strArr := strings.Split(fileScanner.Text(), "")
		var twoDigitNoArr []string

		for _, e := range strArr {
			if _, err := strconv.Atoi(e); err == nil {
				twoDigitNoArr = append(twoDigitNoArr, e)
				break
			}
		}
		slices.Reverse(strArr)
		for _, e := range strArr {
			if _, err := strconv.Atoi(e); err == nil {
				twoDigitNoArr = append(twoDigitNoArr, e)
				break
			}
		}

		twoDigit := strings.Join(twoDigitNoArr, "")
		if num, err := strconv.Atoi(twoDigit); err == nil {
			sum += num
		}

	}
	fmt.Println(sum)

}
