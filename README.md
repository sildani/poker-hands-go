# Poker Hands in Go
A solution to the Poker Hands problem from cyber-dojo.org in Go.

## Problem

A poker deck contains 52 cards - each card has a suit which
is one of clubs, diamonds, hearts, or spades 
(denoted C, D, H, and S in the input data). 

Each card also has a value which is one of 
2, 3, 4, 5, 6, 7, 8, 9, 10, jack, queen, king, ace 
(denoted 2, 3, 4, 5, 6, 7, 8, 9, T, J, Q, K, A). 

For scoring purposes, the suits are unordered while the
values are ordered as given above, with 2 being the lowest
and ace the highest value.

A poker hand consists of 5 cards dealt from the deck. Poker
hands are ranked by the following partial order from lowest
to highest.

High Card: Hands which do not fit any higher category are
ranked by the value of their highest card. If the highest
cards have the same value, the hands are ranked by the next
highest, and so on.

Pair: 2 of the 5 cards in the hand have the same value. 
Hands which both contain a pair are ranked by the value of
the cards forming the pair. If these values are the same, 
the hands are ranked by the values of the cards not 
forming the pair, in decreasing order.

Two Pairs: The hand contains 2 different pairs. Hands 
which both contain 2 pairs are ranked by the value of 
their highest pair. Hands with the same highest pair 
are ranked by the value of their other pair. If these 
values are the same the hands are ranked by the value 
of the remaining card.

Three of a Kind: Three of the cards in the hand have the 
same value. Hands which both contain three of a kind are 
ranked by the value of the 3 cards.

Straight: Hand contains 5 cards with consecutive values. 
Hands which both contain a straight are ranked by their 
highest card.

Flush: Hand contains 5 cards of the same suit. Hands which 
are both flushes are ranked using the rules for High Card.

Full House: 3 cards of the same value, with the remaining 2
cards forming a pair. Ranked by the value of the 3 cards.

Four of a kind: 4 cards with the same value. Ranked by the
value of the 4 cards.

Straight flush: 5 cards of the same suit with consecutive
values. Ranked by the highest card in the hand.

Your job is to rank pairs of poker hands and to indicate
which, if either, has a higher rank.

Examples:
Input: Black: 2H 3D 5S 9C KD White: 2C 3H 4S 8C AH
Output: White wins - high card: Ace 

Input: Black: 2H 4S 4C 2D 4H White: 2S 8S AS QS 3S
Output: Black wins - full house

Input: Black: 2H 3D 5S 9C KD White: 2C 3H 4S 8C KH
Output: Black wins - high card: 9

Input: Black: 2H 3D 5S 9C KD White: 2D 3H 5C 9S KH
Output: Tie

## Approach

I'll keep things simple to start. I'll have a packages for:

- Parsing a hand
- Evaluating a parsed hand to assign a score
- Ranking evaluated parsed hands

Very basic error handling to start.

A main program will have a simple workflow using the packages provided:

1. Parse hands given
1. Evaluate parsed hands
1. Rank parsed hands
1. Output the rank

## Implementation notes

### I'm new

I've been coding in Go for about five days now, so it's still very early for me. My initial implementations will likely be terrible. I just want to have a fun problem to solve to further push my learning. I will tag "releases" as I evolve the program to tie back revisions to my evolution within the language.

### The "calc stats then derrive meaning" tact

As I thought about the implementation, I wanted it to be extremely light. Could I make one pass over the cards and glean enough information via stats to derrive the nature of the hand? Here are some ideas:

- Split the card into two parts - its value, and its suit
- Count the values, count the suits
- Straight Flush will have a delta of five between high and low card, and one single suit count of five (this can only be true for consecutive values)
- Four of a kind will have a single value count of four
- Full House will have one value count of three, and another value count of two
- Flush will have a delta greater than five between high and low card, and one single suit count of five
- Straight will have a delta of five between high and low card, but have at least two distinct suit counts
- Three of a kind will have a single value count of three
- Two Pairs will have two distinct value counts of two each
- One Pair will have one value count of two
- High Card when none of the other rules apply - highest card value applies

For scoring, I will keep the rule set in the evaluator, and each eval will have five "scores" sorted from highest value to lowest value, so that starting at index 0, the highest score wins (in case of tie, go to next score). This will be stored in the Evaluation package, as it will be evaluating hands and assigning scores.