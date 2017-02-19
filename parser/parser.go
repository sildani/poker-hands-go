package parser

import (
	"errors"
	"fmt"
	"strings"
)

var cards [52]string = [52]string{
	"2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "TH", "JH", "QH", "KH", "AH", "2S",
	"3S", "4S", "5S", "6S", "7S", "8S", "9S", "TS", "JS", "QS", "KS", "AS", "2C", "3C",
	"4C", "5C", "6C", "7C", "8C", "9C", "TC", "JC", "QC", "KC", "AC", "2D", "3D", "4D",
	"5D", "6D", "7D", "8D", "9D", "TD", "JD", "QD", "KD", "AD",
}

var valueConversion map[string]int = map[string]int{
	"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8,
	"9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14,
}

func ParseCardValue(s string) (int, error) {
	value := valueConversion[s]

	if value == 0 {
		return value, fmt.Errorf("Invalid card value: Must be single digit 0-9 or one of T (10), J (Jack), Q (Queen), K (King), or A (Ace)")
	} else {
		return value, nil
	}
}

func ParseHand(hand string) ([]string, error) {
	parsedHand := strings.Split(hand, " ")

	if len(parsedHand) != 5 {
		return []string{""}, errors.New("Invalid hand: must have five cards")
	}

	for _, card := range parsedHand {
		isCardValid := false
		for _, validCard := range cards {
			if !isCardValid {
				if card == validCard {
					isCardValid = true
				}
			}
		}
		if !isCardValid {
			return []string{""}, fmt.Errorf("Invalid hand: contains invalid card")
		}
	}

	return parsedHand, nil
}
