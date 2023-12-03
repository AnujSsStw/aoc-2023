import os
current_directory = os.getcwd()

files_in_current_directory = os.listdir(current_directory)
days =  []

for file_name in files_in_current_directory:
    daysFolder  = file_name.startswith("day")
    if daysFolder:
        l = file_name.split("-")
        days.append(l[1])
        
days.sort()
file_names = ["p2.go", "p1.go", "input.txt"]

if len(days) > 0:
    package_name = "day"+str(int(days[len(days)-1])+1)
    directory_name = "day-"+str(int(days[len(days)-1])+1)
    os.makedirs(directory_name)
    for file_name in file_names:
        file_path = os.path.join(directory_name, file_name)
        with open(file_path, 'w') as file:
            if file_name == "p2.go":
                go_code = f"""package {package_name}
import (
day01	"aoc2023/day-1"
"os"
"bufio"
)

func Part2() {{
    readFile, err := os.Open("/home/ani/personal/aoc-2023/{directory_name}/input.txt")
	day01.Check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {{
		line := fileScanner.Text()

    }}

}}
"""
                file.write(go_code)
            if file_name == "p1.go":
                go_code = f"""package {package_name}
import (
day01	"aoc2023/day-1"
"os"
"bufio"
)

func Part1() {{
    readFile, err := os.Open("/home/ani/personal/aoc-2023/{directory_name}/input.txt")
	day01.Check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {{
		line := fileScanner.Text()

    }}

}}
"""
                file.write(go_code)
    
    with open("./main.go", "w") as file:
        file.write(f"""
    package main

import {package_name} "aoc2023/{directory_name}"

func main() {{
        {package_name}.Part1()
}}

""")

                