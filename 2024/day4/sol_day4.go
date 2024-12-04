package day4

import (
	"log"
	"os"
	"strings"
)

func Part1Day4() int {
	// Read the file
	data, err := os.ReadFile("./day4/input_day4.txt")
	if err != nil {
		log.Fatal("Unable to read input data")
	}

	// Parse the data into a 2D slice
	wordSlice := strings.Split(strings.TrimSpace(string(data)), "\n")
	var d [][]rune
	for _, line := range wordSlice {
		d = append(d, []rune(line))
	}

	// Dimensions of the grid
	R := len(d)
	C := len(d[0])

	// Word to find
	word := "XMAS"
	wordLen := len(word)
	ans := 0

	// Iterate through the grid
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			// Check all 8 directions
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					// Skip the case where there's no movement
					if dx == 0 && dy == 0 {
						continue
					}

					// Check if the word matches in the direction
					good := true
					for k := 0; k < wordLen; k++ {
						nx := i + k*dx
						ny := j + k*dy

						// Check boundaries
						if nx < 0 || nx >= R || ny < 0 || ny >= C || d[nx][ny] != rune(word[k]) {
							good = false
							break
						}
					}

					// If all characters matched, increment count
					if good {
						ans++
					}
				}
			}
		}
	}

	return ans
}
func Part2Day4() int {
	// Read the file
	data, err := os.ReadFile("./day4/input_day4.txt")
	if err != nil {
		log.Fatal("Unable to read input data")
	}

	// Parse the data into a 2D slice (grid)
	wordSlice := strings.Split(strings.TrimSpace(string(data)), "\n")
	var grid [][]string
	for _, line := range wordSlice {
		grid = append(grid, strings.Split(line, "")) // Split each line into characters
	}

	rightPatterns := []string{"MAS", "SAM"}

	// Dimensions of the grid
	rows := len(grid)
	cols := len(grid[0])

	// Count the number of X-MAS patterns
	count := 0
	diag := [][]int{{-1, -1}, {0, 0}, {1, 1}}    // Main diagonal movements
	offdiag := [][]int{{1, -1}, {0, 0}, {-1, 1}} // Off-diagonal movements

	// Check for X-MAS patterns
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Check if the current character is 'A' to start the check for X-MAS patterns
			if grid[i][j] == "A" {
				// Initialize variables to store characters for diagonal and off-diagonal patterns
				ofd := ""
				dia := ""

				// Check for the diagonal pattern (X)
				for _, d := range diag {
					nx := i + d[0]
					ny := j + d[1]
					if nx >= 0 && nx < rows && ny >= 0 && ny < cols {
						dia = dia + grid[nx][ny]
					} else {
						break
					}
				}

				// If diagonal pattern length is less than 3, skip to the next iteration
				if len(dia) < 3 {
					continue
				}

				// Check for the off-diagonal pattern (M)
				for _, d := range offdiag {
					nx := i + d[0]
					ny := j + d[1]
					if nx >= 0 && nx < rows && ny >= 0 && ny < cols {
						ofd = ofd + grid[nx][ny]
					} else {
						break
					}
				}

				// If off-diagonal pattern length is less than 3, skip to the next iteration
				if len(ofd) < 3 {
					continue
				}

				// Check if the patterns match the rightPatterns
				if (ofd == rightPatterns[0] || ofd == rightPatterns[1]) &&
					(dia == rightPatterns[0] || dia == rightPatterns[1]) {
					count += 1
				}
			}
		}
	}

	return count
}
