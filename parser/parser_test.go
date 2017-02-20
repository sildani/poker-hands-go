package parser // github.com/sildani/poker-hands-go/parser

import (
	"testing"
)

func TestIsValidCardSuit(t *testing.T) {
	var tests = []struct {
		suit             string
		expectedValidity bool
	}{
		{"", false},
		{"D", true},
		{"H", true},
		{"S", true},
		{"C", true},
		{"A", false},
		{"G", false},
		{"ZZ", false},
		{"Diamonds", false},
		{"Hearts", false},
		{"Spade", false},
		{"Club", false},
	}

	for _, test := range tests {
		validity := IsCardSuitValid(test.suit)
		if validity != test.expectedValidity {
			t.Errorf("IsCardSuitValid(%q) == %t but expected %t", test.suit, validity, test.expectedValidity)
		}
	}
}

func TestParseCardValueInvalidCardValue(t *testing.T) {
	var tests = []struct {
		value               string
		expectedParsedValue int
		expectedErr         string
	}{
		{"0", 0, "Invalid card value: Must be single digit 0-9 or one of T (10), J (Jack), Q (Queen), K (King), or A (Ace)"},
		{"NaN", 0, "Invalid card value: Must be single digit 0-9 or one of T (10), J (Jack), Q (Queen), K (King), or A (Ace)"},
		{"2 3", 0, "Invalid card value: Must be single digit 0-9 or one of T (10), J (Jack), Q (Queen), K (King), or A (Ace)"},
		{"23", 0, "Invalid card value: Must be single digit 0-9 or one of T (10), J (Jack), Q (Queen), K (King), or A (Ace)"},
	}

	for _, test := range tests {
		parsedValue, err := ParseCardValue(test.value)
		if err == nil {
			t.Errorf("ParseCardValue(%q) err == nil but expected %q", test.value, test.expectedErr)
		} else {
			if err.Error() != test.expectedErr {
				t.Errorf("ParseCardValue(%q) err == %q but expected %q", test.value, err, test.expectedErr)
			}
		}
		if parsedValue != test.expectedParsedValue {
			t.Errorf("ParseCardValue(%q) parsedValue == %q but expected %q",
				test.value, parsedValue, test.expectedParsedValue)
		}
	}
}

func TestParseCardValueValidCardValue(t *testing.T) {
	var tests = []struct {
		value               string
		expectedParsedValue int
	}{
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
		{"9", 9},
		{"T", 10},
		{"J", 11},
		{"Q", 12},
		{"K", 13},
		{"A", 14},
	}

	for _, test := range tests {
		parsedValue, err := ParseCardValue(test.value)
		if err != nil {
			t.Errorf("ParseHand(%q) err == %q but expected nil", test.value, err)
		}
		if parsedValue != test.expectedParsedValue {
			t.Errorf("ParseCardValue(%q) == %t but expected %t",
				test.value, parsedValue, test.expectedParsedValue)
		}
	}
}

func TestParseHandInvalidHandMissingCards(t *testing.T) {
	var tests = []struct {
		hand               string
		expectedParsedHand []string
		expectedErr        string
	}{
		{"", []string{""}, "Invalid hand: must have five cards"},
		{" ", []string{""}, "Invalid hand: must have five cards"},
		{"4S", []string{""}, "Invalid hand: must have five cards"},
		{"4S 4C", []string{""}, "Invalid hand: must have five cards"},
		{"4S 4C 2D", []string{""}, "Invalid hand: must have five cards"},
		{"4S 4C 2D 4H", []string{""}, "Invalid hand: must have five cards"},
		{"2H 4S 4C 2D 4H 2S", []string{""}, "Invalid hand: must have five cards"},
		{"not a hand", []string{""}, "Invalid hand: must have five cards"},
		{"not", []string{""}, "Invalid hand: must have five cards"},
	}

	for _, test := range tests {
		parsedHand, err := ParseHand(test.hand)
		if err == nil {
			t.Errorf("ParseCardValue(%q) err == nil but expected %q", test.hand, test.expectedErr)
		} else {
			if err.Error() != test.expectedErr {
				t.Errorf("ParseHand(%q) err == %q but expected %q", test.hand, err, test.expectedErr)
			}
		}
		if parsedHand[0] != test.expectedParsedHand[0] ||
			len(parsedHand) != len(test.expectedParsedHand) {
			t.Errorf("ParseHand(%q) == %q but expected %q",
				test.hand, parsedHand, test.expectedParsedHand)
		}
	}
}

func TestParseHandInvalidHandInvalidCard(t *testing.T) {
	var tests = []struct {
		hand               string
		expectedParsedHand []string
		expectedErr        string
	}{
		{"2H 4S 4C 2D 4P", []string{""}, "Invalid hand: contains invalid card"},
		{"5T 4S 4C 2D 4P", []string{""}, "Invalid hand: contains invalid card"},
		{"2H 4S 4C 2D 10H", []string{""}, "Invalid hand: contains invalid card"},
		{"Every good boy does fine", []string{""}, "Invalid hand: contains invalid card"},
	}

	for _, test := range tests {
		parsedHand, err := ParseHand(test.hand)
		if err == nil {
			t.Errorf("ParseCardValue(%q) err == nil but expected %q", test.hand, test.expectedErr)
		} else {
			if err.Error() != test.expectedErr {
				t.Errorf("ParseHand(%q) err == %q but expected %q", test.hand, err, test.expectedErr)
			}
		}
		if parsedHand[0] != test.expectedParsedHand[0] ||
			len(parsedHand) != len(test.expectedParsedHand) {
			t.Errorf("ParseHand(%q) == %q but expected %q",
				test.hand, parsedHand, test.expectedParsedHand)
		}
	}
}

func TestParseHandInvalidHandDuplicateCard(t *testing.T) {
	var tests = []struct {
		hand               string
		expectedParsedHand []string
		expectedErr        string
	}{
		{"2H 2H 4C 2D 4P", []string{""}, "Invalid hand: contains duplicate card"},
		{"5H 4S 4C 2D 2D", []string{""}, "Invalid hand: contains duplicate card"},
	}

	for _, test := range tests {
		parsedHand, err := ParseHand(test.hand)
		if err == nil {
			t.Errorf("ParseCardValue(%q) err == nil but expected %q", test.hand, test.expectedErr)
		} else {
			if err.Error() != test.expectedErr {
				t.Errorf("ParseHand(%q) err == %q but expected %q", test.hand, err, test.expectedErr)
			}
		}
		if parsedHand[0] != test.expectedParsedHand[0] ||
			len(parsedHand) != len(test.expectedParsedHand) {
			t.Errorf("ParseHand(%q) == %q but expected %q",
				test.hand, parsedHand, test.expectedParsedHand)
		}
	}
}

func TestParseHandValidHand(t *testing.T) {
	hand := "2H 4S 4C 2D 4H"
	parsedHand, err := ParseHand(hand)

	if err != nil {
		t.Errorf("ParseHand(%q) err == %q but expected nil", hand, err)
	}

	handLen := len(parsedHand)
	expectedHandLen := 5
	if handLen != expectedHandLen {
		t.Errorf("len(ParseHand(%q)) == %d but expected %d", hand, handLen, expectedHandLen)
	}

	expectedCards := []string{"2H", "4S", "4C", "2D", "4H"}
	for i, expectedCard := range expectedCards {
		card := parsedHand[i]
		if card != expectedCard {
			t.Errorf("parsedHand[0] == %q but expected %q", card, expectedCard)
		}
	}
}
