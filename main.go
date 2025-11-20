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
		var myHand Hand
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
		UpdateHandArrays(&myHand)
		UpdateRankMap(&myHand)
		UpdateSuitMap(&myHand)

		fmt.Println(myHand)
	}

}

// func handRanking(h Hand) string {

// }
