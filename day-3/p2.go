package day3
import (
day01	"aoc2023/day-1"
"os"
"bufio"
)

func Part2() {
    readFile, err := os.Open("/home/ani/personal/aoc-2023/day-3/input.txt")
	day01.Check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()

    }

}
