package day2

import (
	"log"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func checkFollowsTheRule(nums []int) bool {
	if len(nums) < 1 {
		return false
	}
	res := slices.IsSorted(nums)

	if !res {
		slices.Reverse(nums)
		res2 := slices.IsSorted(nums)
		if !res2 {
			return false
		}
	}

	for i := range len(nums) - 1 {
		diff := math.Abs(float64(nums[i] - nums[i+1]))

		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func Part1Day2() int {

	data, err := os.ReadFile("./day2/input_day2.txt")

	if err != nil {
		log.Fatal("Unable to read input data")
	}

	d := string(data)

	splitString := strings.Split(d, "\n")

	var ans int

	for _, k := range splitString {
		numArrayString := strings.Split(k, " ")
		nums := make([]int, 0)
		for _, l := range numArrayString {
			if l == "" {
				continue
			}
			n, err := strconv.Atoi(l)
			if err != nil {
				log.Fatal("Unable to convert l ", l, "to integer ", err)
			}
			nums = append(nums, n)
		}

		y := checkFollowsTheRule(nums)

		if y == true {
			ans = ans + 1
		}
	}

	return ans
}

func isDesc(nums []int) bool {
	return sort.SliceIsSorted(nums, func(i, j int) bool { return nums[i] > nums[j] })
}
func isAsc(nums []int) bool {
	return sort.SliceIsSorted(nums, func(i, j int) bool { return nums[i] < nums[j] })
}

func checkFollowsTheRuleWithProblemDampner(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	agood := false
	for i := 0; i < len(nums); i++ {
    onums := append([]int{}, nums...)
		onums = append(onums[:i], onums[i+1:]...)

		if !isAsc(onums) && !isDesc(onums) {
			continue
		}

		slices.Sort(onums)

		good := true

		for k := 0; k < len(onums)-1; k++ {
			diff := math.Abs(float64(onums[k+1] - onums[k]))
			if diff <= 0 || diff > 3 {
				good = false
				break
			}
		}

		if good {
			agood = true
      return agood
		}
	}

	return agood
}

func Part2Day2() int {

	data, err := os.ReadFile("./day2/input_day2.txt")

	if err != nil {
		log.Fatal("Unable to read input data")
	}

	d := string(data)

	splitString := strings.Split(d, "\n")

	var ans int

	for _, k := range splitString {
		numArrayString := strings.Split(k, " ")
		nums := make([]int, 0)
		for _, l := range numArrayString {
			if l == "" {
				continue
			}
			n, err := strconv.Atoi(l)
			if err != nil {
				log.Fatal("Unable to convert l ", l, "to integer ", err)
			}
			nums = append(nums, n)
		}
		y := checkFollowsTheRuleWithProblemDampner(nums)

		if y == true {
			ans = ans + 1
		}
	}

	return ans
}
