package parser // github.com/sildani/poker-hands-go/parser

import (
	"testing"
)

func TestParseInvalidHandMissingCards(t *testing.T) {
	var tests = []struct {
		hand, expectedErr string
	}{
		{"", "Invalid hand: must have five cards"},
		{" ", "Invalid hand: must have five cards"},
		{"4S", "Invalid hand: must have five cards"},
		{"4S 4C", "Invalid hand: must have five cards"},
		{"4S 4C 2D", "Invalid hand: must have five cards"},
		{"4S 4C 2D 4H", "Invalid hand: must have five cards"},
		{"2H 4S 4C 2D 4H 2S", "Invalid hand: must have five cards"},
		{"not a hand", "Invalid hand: must have five cards"},
		{"not", "Invalid hand: must have five cards"},
	}

	for _, test := range tests {
		_, err := ParseHand(test.hand)
		if err == nil {
			t.Errorf("ParseHand(%q) didn't return a error where it was expected", test.hand)
		} else {
			if err.Error() != test.expectedErr {
				t.Errorf("ParseHand(%q) err == %q but expected %q", test.hand, err, test.expectedErr)
			}
		}
	}
}

func TestParseInvalidHandInvalidCards(t *testing.T) {
	var tests = []struct {
		hand, expectedErr string
	}{
		{"2H 4S 4C 2D 4P", "Invalid hand: contains invalid card (4P)"},
		{"5T 4S 4C 2D 4P", "Invalid hand: contains invalid card (5T)"},
		{"2H 4S 4C 2D 10H", "Invalid hand: contains invalid card (10H)"},
		{"Every good boy does fine", "Invalid hand: contains invalid card (Every)"},
	}

	for _, test := range tests {
		_, err := ParseHand(test.hand)
		if err == nil {
			t.Errorf("ParseHand(%q) didn't return a error where it was expected", test.hand)
		} else {
			if err.Error() != test.expectedErr {
				t.Errorf("ParseHand(%q) err == %q but expected %q", test.hand, err, test.expectedErr)
			}
		}
	}
}

func TestParseValidHand(t *testing.T) {
	input := "2H 4S 4C 2D 4H"
	parsedHand, err := ParseHand(input)

	if err != nil {
		t.Errorf("ParseHand(%q) err == %q but expected nil", input, err)
	}

	handLen := len(parsedHand)
	expectedHandLen := 5
	if handLen != expectedHandLen {
		t.Errorf("len(ParseHand(%q)) == %d but expected %d", input, handLen, expectedHandLen)
	}

	expectedCards := []string{"2H", "4S", "4C", "2D", "4H"}
	for i, expectedCard := range expectedCards {
		card := parsedHand[i]
		if card != expectedCard {
			t.Errorf("parsedHand[0] == %q but expected %q", card, expectedCard)
		}
	}
}
