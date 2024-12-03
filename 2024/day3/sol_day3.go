package day3

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

func Part1Day3() int {

	data, err := os.ReadFile("./day3/input_day3.txt")

	if err != nil {
		log.Fatal("Unable to read input data")
	}

	d := string(data)
	pattern := `mul\((\d+),(\d+)\)`

	re, err := regexp.Compile(pattern)

	if err != nil {
		log.Fatal("Unable to compile regex pattern", err)
	}

	matches := re.FindAllStringSubmatch(d, -1)
	ans := 0
	for _, match := range matches {
		a, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal("Unable to convert a")
		}
		b, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal("Unable to convert b")
		}
		ans += (a * b)
	}

	return ans
}

func Part2Day3() int {

	data, err := os.ReadFile("./day3/input_day3.txt")

	if err != nil {
		log.Fatal("Unable to read input data")
	}

	d := string(data)
	pattern := `(mul\((\d+),(\d+)\))|(do\(\))|(don't\(\))`

	re, err := regexp.Compile(pattern)

	if err != nil {
		log.Fatal("Unable to compile regex pattern", err)
	}

	matches := re.FindAllStringSubmatch(d, -1)
	ans := 0
	isMulEnabled := true
	for _, match := range matches {
		// If the match is a mul() instruction and mul is enabled, print the digits
		if match[1] != "" {
			if isMulEnabled {
        a, err := strconv.Atoi(match[2])
        if err != nil {
            log.Fatal("unable to convert to a")
        }
        b, err := strconv.Atoi(match[3])
        if err != nil {
            log.Fatal("unable to convert to b")
        }
        ans += (a*b)
			}
		}

		// If the match is a do() instruction, enable mul instructions
		if match[0] == "do()" {
			isMulEnabled = true
		}

		// If the match is a don't() instruction, disable mul instructions
		if match[0] == "don't()" {
			isMulEnabled = false
		}
	}
	return ans
}
