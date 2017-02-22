package evaluator // github.com/sildani/poker-hands-go/evaluator

import (
	"testing"
)

func TestEvaluateParsedHand(t *testing.T) {
	tests := []struct {
		parsedHand         []string
		expectedEvaluation Evaluation
	}{
		{
			[]string{
				"2D", "3D", "4D", "5D", "6D",
			},
			Evaluation{
				hand: "2D 3D 4D 5D 6D",
				result: [5]struct {
					score       int
					description string
				}{
					{906, "Straight Flush, High Card: 6"},
					{905, "Straight Flush, High Card: 5"},
					{904, "Straight Flush, High Card: 4"},
					{903, "Straight Flush, High Card: 3"},
					{902, "Straight Flush, High Card: 2"},
				},
			},
		},
		{
			[]string{
				"2D", "2C", "2S", "2H", "6D",
			},
			Evaluation{
				hand: "2D 2C 2S 2H 6D",
				result: [5]struct {
					score       int
					description string
				}{
					{802, "Four of a kind, High Card: 2"},
					{106, "High Card: 6"},
					{106, "High Card: 6"},
					{106, "High Card: 6"},
					{106, "High Card: 6"},
				},
			},
		},
	}

	for _, test := range tests {
		evaluation := EvaluateParsedHand(test.parsedHand)
		if evaluation.hand != test.expectedEvaluation.hand {
			t.Errorf("evaluation.hand == %q but expected %q for hand: %v",
				evaluation.hand, test.expectedEvaluation.hand, test.parsedHand)
		}
		if evaluation.result != test.expectedEvaluation.result {
			t.Errorf("evaluation.result == %v but expected %v for hand: %v",
				evaluation.result, test.expectedEvaluation.result, test.parsedHand)
		}
	}
}

func TestGatherStatsInvalidHand(t *testing.T) {
	tests := []struct {
		parsedHand    []string
		expectedStats Stats
		expectedErr   string
	}{
		{
			[]string{"2D", "3D", "PP", "5D", "6D"},
			Stats{suits: map[string]int{"": 0}, values: map[int]int{0: 0}},
			"Parsed hand contained an invalid card. Did you use parser package to parse hand from user input?",
		},
		{
			[]string{"2D", "3D", "Ace of Spades", "5D", "6D"},
			Stats{suits: map[string]int{"": 0}, values: map[int]int{0: 0}},
			"Parsed hand contained an invalid card. Did you use parser package to parse hand from user input?",
		},
		{
			[]string{"2D", "3D", "AZ", "5D", "6D"},
			Stats{suits: map[string]int{"": 0}, values: map[int]int{0: 0}},
			"Parsed hand contained an invalid card. Did you use parser package to parse hand from user input?",
		},
		{
			[]string{"2D", "3D", "5D", "6D"},
			Stats{suits: map[string]int{"": 0}, values: map[int]int{0: 0}},
			"Parsed hand must contain five cards. Did you use parser package to parse hand from user input?",
		},
		{
			[]string{"PP", "3D", "5D", "6D"},
			Stats{suits: map[string]int{"": 0}, values: map[int]int{0: 0}},
			"Parsed hand must contain five cards. Did you use parser package to parse hand from user input?",
		},
		{
			[]string{"2D", "3D", "4D", "5D", "6D", "7D"},
			Stats{suits: map[string]int{"": 0}, values: map[int]int{0: 0}},
			"Parsed hand must contain five cards. Did you use parser package to parse hand from user input?",
		},
		{
			[]string{""},
			Stats{suits: map[string]int{"": 0}, values: map[int]int{0: 0}},
			"Parsed hand must contain five cards. Did you use parser package to parse hand from user input?",
		},
	}

	for _, test := range tests {
		stats, err := gatherStats(test.parsedHand)
		if err == nil {
			t.Errorf("gatherStats(%q) err == nil but expected %v", test.parsedHand, test.expectedErr)
		} else {
			if err.Error() != test.expectedErr {
				t.Errorf("gatherStats(%q) err == %q but expected %q", test.parsedHand, err, test.expectedErr)
			}
		}
		if stats.suits[""] != test.expectedStats.suits[""] ||
			len(stats.suits) != len(test.expectedStats.suits) {
			t.Errorf("gatherStats(%q) stats == %v but expected %v",
				test.parsedHand, stats.suits, test.expectedStats.suits)
		}
		if stats.values[0] != test.expectedStats.values[0] ||
			len(stats.values) != len(test.expectedStats.values) {
			t.Errorf("gatherStats(%q) stats == %v but expected %v",
				test.parsedHand, stats.values, test.expectedStats.values)
		}
	}
}

func TestGatherStats(t *testing.T) {
	tests := []struct {
		parsedHand    []string
		expectedStats Stats
	}{
		{
			[]string{"2D", "3D", "4D", "5D", "6D"},
			Stats{
				suits:  map[string]int{"D": 5},
				values: map[int]int{2: 1, 3: 1, 4: 1, 5: 1, 6: 1},
			},
		},
		{
			[]string{"TD", "JD", "QD", "KD", "AD"},
			Stats{
				suits:  map[string]int{"D": 5},
				values: map[int]int{10: 1, 11: 1, 12: 1, 13: 1, 14: 1},
			},
		},
		{
			[]string{"2S", "2C", "QD", "KD", "AD"},
			Stats{
				suits:  map[string]int{"D": 3, "S": 1, "C": 1},
				values: map[int]int{2: 2, 12: 1, 13: 1, 14: 1},
			},
		},
	}

	for _, test := range tests {
		stats, err := gatherStats(test.parsedHand)
		if err != nil {
			t.Errorf("gatherStats(%q) err == %q but expected nil", test.parsedHand, err)
		}
		for k, v := range test.expectedStats.suits {
			if v != stats.suits[k] {
				t.Errorf("test.expectedStats.suits[%q] == %v but expected %v for hand %v",
					k, stats.suits[k], v, test.parsedHand)
			}
		}
		for k, v := range test.expectedStats.values {
			if v != stats.values[k] {
				t.Errorf("test.expectedStats.values[%v] == %v but expected %v for hand %v",
					k, stats.values[k], v, test.parsedHand)
			}
		}
	}
}

func TestIsStraightFlush(t *testing.T) {
	tests := []struct {
		stats          Stats
		expectedResult bool
	}{
		{
			// 2D 3D 4D 5D 6D
			Stats{
				suits:  map[string]int{"D": 5},
				values: map[int]int{2: 1, 3: 1, 4: 1, 5: 1, 6: 1},
			},
			true,
		},
		{
			// TH JH QH KH AH
			Stats{
				suits:  map[string]int{"H": 5},
				values: map[int]int{14: 1, 13: 1, 12: 1, 11: 1, 10: 1},
			},
			true,
		},
		{
			// TH JH QH KH AH
			// Stats ordering shouldn't matter
			Stats{
				suits:  map[string]int{"H": 5},
				values: map[int]int{12: 1, 11: 1, 10: 1, 14: 1, 13: 1},
			},
			true,
		},
		{
			// TH JH QH KD AH
			Stats{
				suits:  map[string]int{"H": 4, "D": 1},
				values: map[int]int{14: 1, 13: 1, 12: 1, 11: 1, 10: 1},
			},
			false,
		},
		{
			// AD AH QS JS TC
			Stats{
				suits:  map[string]int{"H": 1, "D": 1, "C": 1, "S": 2},
				values: map[int]int{14: 2, 12: 1, 11: 1, 10: 1},
			},
			false,
		},
	}

	for _, test := range tests {
		result := isStraightFlush(test.stats)
		if result != test.expectedResult {
			t.Errorf("isStraightFlush(%v) == %t but expected %t", test.stats, result, test.expectedResult)
		}
	}
}

func TestIsFourOfAKind(t *testing.T) {
	tests := []struct {
		stats          Stats
		expectedResult bool
	}{
		{
			// 2C 2D 2H 2S 3S
			Stats{
				suits:  map[string]int{"C": 1, "D": 1, "H": 1, "S": 2},
				values: map[int]int{2: 4, 3: 1},
			},
			true,
		},
		{
			// TC TD KD TH TS
			Stats{
				suits:  map[string]int{"C": 1, "D": 2, "H": 1, "S": 1},
				values: map[int]int{10: 4, 13: 1},
			},
			true,
		},
		{
			// TC TD KD TH TS
			// Stats ordering shouldn't matter
			Stats{
				suits:  map[string]int{"D": 2, "C": 1, "H": 1, "S": 1},
				values: map[int]int{13: 1, 10: 4},
			},
			true,
		},
		{
			// TH JH QH KD AH
			Stats{
				suits:  map[string]int{"H": 4, "D": 1},
				values: map[int]int{14: 1, 13: 1, 12: 1, 11: 1, 10: 1},
			},
			false,
		},
		{
			// AD AH QS JS TC
			Stats{
				suits:  map[string]int{"H": 1, "D": 1, "C": 1, "S": 2},
				values: map[int]int{14: 2, 12: 1, 11: 1, 10: 1},
			},
			false,
		},
	}

	for _, test := range tests {
		result := isFourOfAKind(test.stats)
		if result != test.expectedResult {
			t.Errorf("isFourOfAKind(%v) == %t but expected %t", test.stats, result, test.expectedResult)
		}
	}
}

// TODO: Full House: 3 cards of the same value, with the remaining 2
// cards forming a pair. Ranked by the value of the 3 cards.

// TODO: Flush: Hand contains 5 cards of the same suit. Hands which
// are both flushes are ranked using the rules for High Card.

// TODO: Straight: Hand contains 5 cards with consecutive values.
// Hands which both contain a straight are ranked by their
// highest card.

// TODO: Three of a Kind: Three of the cards in the hand have the
// same value. Hands which both contain three of a kind are
// ranked by the value of the 3 cards.

// TODO: Two Pairs: The hand contains 2 different pairs. Hands
// which both contain 2 pairs are ranked by the value of
// their highest pair. Hands with the same highest pair
// are ranked by the value of their other pair. If these
// values are the same the hands are ranked by the value
// of the remaining card.

// TODO: Pair: 2 of the 5 cards in the hand have the same value.
// Hands which both contain a pair are ranked by the value of
// the cards forming the pair. If these values are the same,
// the hands are ranked by the values of the cards not
// forming the pair, in decreasing order.

// TODO: High Card: Hands which do not fit any higher category are
// ranked by the value of their highest card. If the highest
// cards have the same value, the hands are ranked by the next
// highest, and so on.
