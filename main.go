package main

import (
"fmt"
 "github.com/sildani/poker-hands-go/parser"
)

func main() {
  // hand := "AD AH QS JS TC"
  hand := "AD AD QS JS TC"
  parsedHand, _ := parser.ParseHand(hand)
  fmt.Printf("parsedHand: %v\n", parsedHand)
}