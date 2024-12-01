package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func processString(s string) int64 {
  var numbers []int 

  for _, e := range s {
    if e >= '0' && e<='9' {
      n, err := strconv.Atoi(string(e))
      if err != nil {
        log.Fatal("cannot convert to int")
      }
      numbers = append(numbers,n) 
    }
  }
  
  var ans int64
  
  if len(numbers) == 0 {
    return 0
  }

  ans+=int64(numbers[0])
  ans*=10
  ans+=int64(numbers[len(numbers)-1])
  return ans 
}


func main() {
  data, err := os.ReadFile("day_01_input.txt")  
  if err != nil {
    log.Fatal("Unable to read input file")
  }

  var ans int64
  
  d := string(data)
  
  k := strings.Split(d, "\n")
  for _, each := range k {
    ans+=processString(each)
  }
  fmt.Print(ans)
}
