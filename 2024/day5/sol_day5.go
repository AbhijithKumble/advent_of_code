package day5

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(filename string) int {
	// Read file content
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Unable to read file:", err)
	}

	// Split rules and updates
	sections := strings.Split(string(data), "\n\n")
	if len(sections) < 2 {
		log.Fatal("Invalid input format")
	}
	rules := strings.Split(sections[0], "\n")
	updateStrings := strings.Split(sections[1], "\n")

	// Parse rules into a "before" map
	before := make(map[int]map[int]struct{})
	for _, rule := range rules {
		if rule == "" {
			continue
		}
		parts := strings.Split(rule, "|")
		a, err1 := strconv.Atoi(parts[0])
		b, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			log.Fatal("Invalid rule format:", rule)
		}
		if before[b] == nil {
			before[b] = make(map[int]struct{})
		}
		before[b][a] = struct{}{}
	}

	// Parse updates
	var updates [][]int
	for _, line := range updateStrings {
		if line == "" {
			continue
		}
		nums := strings.Split(line, ",")
		var upd []int
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal("Invalid update number:", num)
			}
			upd = append(upd, n)
		}
		updates = append(updates, upd)
	}

	// Calculate the answer
	ans := 0
	for _, upd := range updates {
		// Check if the update is in valid order
		cupd := make([]int, len(upd))
		copy(cupd, upd)
		sort.Slice(cupd, func(i, j int) bool {
			return compare(cupd[i], cupd[j], before)
		})
		if equal(upd, cupd) {
			ans += upd[len(upd)/2]
		}
	}
	return ans
}

// compare returns true if x should come before y based on the rules in "before".
func compare(x, y int, before map[int]map[int]struct{}) bool {
	if _, exists := before[y][x]; exists {
		return true
	}
	if _, exists := before[x][y]; exists {
		return false
	}
	return x < y // Fallback to natural order
}

// equal checks if two slices are equal.
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Part1Day5() int {
	// Read the file
	return part1("./day5/input_day5.txt")
}

func part2(filename string) int {
	// Read file content
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Unable to read file:", err)
	}

	// Split rules and updates
	sections := strings.Split(string(data), "\n\n")
	if len(sections) < 2 {
		log.Fatal("Invalid input format")
	}
	rules := strings.Split(sections[0], "\n")
	updateStrings := strings.Split(sections[1], "\n")

	// Parse rules into a "before" map
	before := make(map[int]map[int]struct{})
	for _, rule := range rules {
		if rule == "" {
			continue
		}
		parts := strings.Split(rule, "|")
		a, err1 := strconv.Atoi(parts[0])
		b, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			log.Fatal("Invalid rule format:", rule)
		}
		if before[b] == nil {
			before[b] = make(map[int]struct{})
		}
		before[b][a] = struct{}{}
	}

	// Parse updates
	var updates [][]int
	for _, line := range updateStrings {
		if line == "" {
			continue
		}
		nums := strings.Split(line, ",")
		var upd []int
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal("Invalid update number:", num)
			}
			upd = append(upd, n)
		}
		updates = append(updates, upd)
	}

	// Process incorrectly ordered updates
	sum := 0
	for _, upd := range updates {
		cupd := make([]int, len(upd))
		copy(cupd, upd)
		sort.Slice(cupd, func(i, j int) bool {
			return compare(cupd[i], cupd[j], before)
		})
		if !equal(upd, cupd) {
			middle := cupd[len(cupd)/2]
			sum += middle
		}
	}
	return sum
}

func Part2Day5() int {
	// Read the file
	return part2("./day5/input_day5.txt")
}
