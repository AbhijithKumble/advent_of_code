package day1

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part2Day1() int {

	data, err := os.ReadFile("./day1/input_day1.txt")

	if err != nil {
		log.Fatal("Unable to read input data")
	}

	d := string(data)

	var s1 []int
	var s2 []int

	splitString := strings.Split(d, "\n")

	for _, i := range splitString {
		if i == "" {
			break
		}

		splitData := strings.Split(i, "   ")

		n1, err := strconv.Atoi(splitData[0])
		if err != nil {
			log.Fatal("Could not convert n1 to integer", err)
		}

		n2, err := strconv.Atoi(splitData[1])
		if err != nil {
			log.Fatal("Could not convert n2 to integer", err)
		}
		s1 = append(s1, n1)
		s2 = append(s2, n2)
	}

	//sort.Ints(s1)
	//sort.Ints(s2)
  sim1 := make(map[int]int)
  sim2 := make(map[int]int)
  
  for i := range len(s1) {
    sim1[s1[i]]++
  } 

  for i := range len(s2) {
    sim2[s2[i]]++
  } 

  ans := 0
  for v1, v2 := range sim1 {
    _, ok := sim2[v1] 
    
    if ok {
      ans += (v1*v2*sim2[v1])
    }else {
      continue
    }
  }
  
  return ans
}

func Part1Day1() {

	data, err := os.ReadFile("./input_day1.txt")

	if err != nil {
		log.Fatal("Unable to read input data")
	}

	d := string(data)

	var s1 []int
	var s2 []int

	splitString := strings.Split(d, "\n")

	for _, i := range splitString {
		if i == "" {
			break
		}

		splitData := strings.Split(i, "   ")

		n1, err := strconv.Atoi(splitData[0])
		if err != nil {
			log.Fatal("Could not convert n1 to integer", err)
		}

		n2, err := strconv.Atoi(splitData[1])
		if err != nil {
			log.Fatal("Could not convert n2 to integer", err)
		}
		s1 = append(s1, n1)
		s2 = append(s2, n2)
	}

	sort.Ints(s1)
	sort.Ints(s2)

	ans := 0

	for i := range len(s1) {
		ans += int(math.Abs(float64(s1[i] - s2[i])))
	}

	fmt.Print(ans)
	return
}
