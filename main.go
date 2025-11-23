package main

/*
   1. Royal Flush: The best possible hand. It's a sequence of A, K, Q, J, 10, all of the same suit (e.g., all Hearts).
   2. Straight Flush: Five cards in a row (e.g., 5, 6, 7, 8, 9), all of the same suit.
   3. Four of a Kind: Four cards of the same rank (e.g., four Kings).
   4. Full House: Three cards of one rank and two cards of another rank (e.g., three 2s and two Queens). This is the example from your assignment.
   5. Flush: Any five cards of the same suit, but not in a sequence.
   6. Straight: Five cards in a row, but with different suits.
   7. Three of a Kind: Three cards of the same rank (e.g., three Jacks).
   8. Two Pair: Two cards of one rank, and two cards of another rank (e.g., two 8s and two Kings).
   9. One Pair: Two cards of the same rank (e.g., two Aces).
   10. High Card: If a hand doesn't fit any of the above categories, its value is determined by its single highest card.
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your hand: ")

	for scanner.Scan() {
		myHand := NewHand()
		myInput := scanner.Text()
		cardsParts := strings.Fields(myInput)

		// parsedCards := [5]Card

		fmt.Println(cardsParts)
		for _, p := range cardsParts {

			suit := p[:1]
			rank := p[1:]

			switch rank {
			case "A":
				rank = "14"
			case "K":
				rank = "13"
			case "Q":
				rank = "12"
			case "J":
				rank = "11"
			case "T":
				rank = "10"
			}

			i, err := strconv.Atoi(rank)
			if err != nil {
				fmt.Println("error occured converting string to int")
				panic(err)
			}

			card := Card{
				Suit: suit,
				Rank: i,
			}
			myHand.Cards = append(myHand.Cards, card)

			fmt.Println(card)
		}

		organizeHand(myHand)
		handRanking(myHand)

		fmt.Println(myHand)
	}

}

func HandComparison(h1 *Hand, h2 *Hand) string {

	organizeHand(h1)
	organizeHand(h2)

	// response := fmt.Sprintf("Variable string %d content", data)

	if h1.HandValue > h2.HandValue {
		return fmt.Sprintf("Hand number 1 wins with the following hand: %h and and Hand number 2 loses with the following hand %h", h1, h2)
	}

	for i := len(h1.rankArraySorted) - 1; i >= 0; i-- {
		// Get the highest card from each hand
		card1_rank := h1.rankArraySorted[i]
		card2_rank := h2.rankArraySorted[i]

		// What comparison do you do here?
		if card1_rank < card2_rank {
			return "Hand 1 wins"
		}
		if card2_rank > card1_rank {
			return "Hand 2 wins"
		}

	}

	return "It's a tie"

}

