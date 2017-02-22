package evaluator // github.com/sildani/poker-hands-go/evaluator

import (
	"fmt"
	"github.com/sildani/poker-hands-go/parser"
	"sort"
	"strconv"
	"strings"
)

const straightFlushBaseScore = 900
const fourOfAKindBaseScore = 800
const highCardBaseScore = 100

type Stats struct {
	suits  map[string]int
	values map[int]int
}

type Evaluation struct {
	hand   string
	result [5]struct {
		score       int
		description string
	}
}

func EvaluateParsedHand(parsedHand []string) Evaluation {
	stats, _ := gatherStats(parsedHand)

	result := [5]struct {
		score       int
		description string
	}{
		{score: 0, description: ""},
		{score: 0, description: ""},
		{score: 0, description: ""},
		{score: 0, description: ""},
		{score: 0, description: ""},
	}

	if isStraightFlush(stats) {
		sortedValues := []int{0, 0, 0, 0, 0}
		i := 0
		for k, _ := range stats.values {
			sortedValues[i] = k
			i++
		}
		sort.Sort(sort.Reverse(sort.IntSlice(sortedValues)))

		for j, value := range sortedValues {
			result[j] = struct {
				score       int
				description string
			}{
				score:       straightFlushBaseScore + value,
				description: "Straight Flush, High Card: " + strconv.Itoa(value),
			}
		}
	} else if isFourOfAKind(stats) {
		for value, count := range stats.values {
			if count == 4 {
				result[0] = struct {
					score       int
					description string
				}{
					score:       fourOfAKindBaseScore + value,
					description: "Four of a kind, High Card: " + strconv.Itoa(value),
				}
			} else {
				for i := 1; i < 5; i++ {
					result[i] = struct {
						score       int
						description string
					}{
						score:       highCardBaseScore + value,
						description: "High Card: " + strconv.Itoa(value),
					}
				}
			}
		}
	}

	return Evaluation{
		hand:   strings.Join(parsedHand, " "),
		result: result,
	}
}

func gatherStats(parsedHand []string) (Stats, error) {
	if len(parsedHand) != 5 {
		return Stats{suits: map[string]int{"": 0}, values: map[int]int{0: 0}},
			fmt.Errorf("Parsed hand must contain five cards. Did you use parser package to parse hand from user input?")
	}

	suits := make(map[string]int)
	values := make(map[int]int)

	for _, card := range parsedHand {
		parsedCard := strings.Split(card, "")

		suit := parsedCard[1]

		value, err := parser.ParseCardValue(parsedCard[0])

		if len(parsedCard) != 2 ||
			!parser.IsCardSuitValid(suit) ||
			err != nil {
			return Stats{suits: map[string]int{"": 0}, values: map[int]int{0: 0}},
				fmt.Errorf("Parsed hand contained an invalid card. Did you use parser package to parse hand from user input?")
		}

		suits[suit] += 1
		values[value] += 1
	}

	return Stats{suits, values}, nil
}

func isStraightFlush(stats Stats) bool {
	values := make([]int, len(stats.values))
	i := 0
	for value, _ := range stats.values {
		values[i] = value
		i++
	}
	sort.Ints(values)
	return values[len(values)-1]-values[0] == 4 && len(stats.suits) == 1
}

func isFourOfAKind(stats Stats) bool {
	return len(stats.suits) == 4 && len(stats.values) == 2
}
