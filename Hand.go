package main

import (
	"sort"
	// "golang.org/x/tools/go/analysis/passes/bools"
)

type Hand struct {
	Cards           []Card
	rankMap         map[int]int
	suitMap         map[string]int
	heartsArray     []int
	diamondsArray   []int
	clubsArray      []int
	spadesArray     []int
	rankArraySorted []int
	// isWhatRank      []bool
	handRanking string
	HandValue int
}

func NewHand() *Hand {
	return &Hand{
		rankMap: make(map[int]int),
		suitMap: make(map[string]int),
	}
}

func UpdateRankMap(h *Hand) {
	for i := 0; i < len(h.Cards); i++ {
		currentRank := h.Cards[i].Rank
		h.rankMap[currentRank] = h.rankMap[currentRank] + 1 // possible alternative to above code
	}
}

func UpdateSuitMap(h *Hand) {
	for i := 0; i < len(h.Cards); i++ {
		currentSuit := h.Cards[i].Suit
		h.suitMap[currentSuit] = h.suitMap[currentSuit] + 1 // possible alternative to above code
	}
}

func UpdateHandArrays(h *Hand) {
	for i := 0; i < len(h.Cards); i++ {
		currentCard := h.Cards[i]
		if currentCard.Suit == "H" {
			h.heartsArray = append(h.heartsArray, currentCard.Rank)
		} else if currentCard.Suit == "D" {
			h.diamondsArray = append(h.diamondsArray, currentCard.Rank)
		} else if currentCard.Suit == "C" {
			h.clubsArray = append(h.clubsArray, currentCard.Rank)
		} else if currentCard.Suit == "S" {
			h.spadesArray = append(h.spadesArray, currentCard.Rank)
		}
		h.rankArraySorted = append(h.rankArraySorted, currentCard.Rank)
	}
	sort.Ints(h.rankArraySorted)

}

func organizeHand(h *Hand) {
	UpdateHandArrays(h)
	UpdateRankMap(h)
	UpdateSuitMap(h)

}

func handRanking(h *Hand) string {
	if isRoyalFlush(h) {
		h.handRanking = "Royal Flush"
		h.HandValue = 1
		return h.handRanking
	} else if isStraightFlush(h) {
		h.HandValue = 2
		h.handRanking = "Straight Flush"
		return h.handRanking
	} else if isFourOfAKind(h) {
		h.HandValue = 3
		h.handRanking = "Four of a Kind"
		return h.handRanking
	} else if isFullHouse(h) {
		h.HandValue = 4
		h.handRanking = "Full House"
		return h.handRanking
	} else if isFlush(h) {
		h.HandValue = 5
		h.handRanking = "Flush"
		return h.handRanking
	} else if isStraight(h) {
		h.HandValue = 6
		h.handRanking = "Straight"
		return h.handRanking
	} else if isThreeOfAKind(h) {
		h.HandValue = 7
		h.handRanking = "Three of a Kind"
		return h.handRanking
	} else if isTwoPair(h) {
		h.HandValue = 8
		h.handRanking = "Two Pair"
		return h.handRanking
	} else if isOnePair(h) {
		h.HandValue = 9
		h.handRanking = "One Pair"
		return h.handRanking
	} else {
		h.HandValue = 10
		h.handRanking = "High Card"
		return h.handRanking
	}

}

func isRoyalFlush(h *Hand) bool {

	if isFlush(h) && isStraight(h) {
		if h.rankArraySorted[0] == 10 {
			return true
		}
	}

	// potential other solution
	/* 	if h.rankMap[10] != 1 || h.rankMap[11] != 1 || h.rankMap[12] != 1 || h.rankMap[13] != 1 || h.rankMap[14] != 1 {
	   		return false
	   	}
	   	if h.suitMap["H"] != 5 || h.suitMap["D"] != 5 || h.suitMap["C"] != 5 || h.suitMap["S"] != 5 {
	   		return false
	   	} */
	return false
}

func isStraightFlush(h *Hand) bool { // there is an error here

	if isStraight(h) && isFlush(h) {
		return true
	} else {
		return false
	}

	/*
		 	if !(len(h.clubsArray) == 5 || len(h.diamondsArray) == 5 || len(h.clubsArray) == 5 || len(h.spadesArray) == 5) {
				return false
			}

			for key,value := range h.rankMap {

				if value > 1 {
					return false
				}

				if value > 0 {
					if h.rankMap[key+i] != 1 {
					return false
		    		}
				}

				}
				return true
	*/
}

func isFourOfAKind(h *Hand) bool {
	for _, value := range h.rankMap {
		if value >= 4 {
			return true
		}
	}
	return false
}

func isFullHouse(h *Hand) bool {
	var hasTrio bool
	var hasPair bool

	for _, value := range h.rankMap {
		if value >= 4 {
			return false
		}
		if value == 2 {
			hasPair = true
		}

		if value == 3 {
			hasTrio = true
		}
	}
	if hasTrio && hasPair {
		return true
	}

	return false
}

func isFlush(h *Hand) bool {
	if len(h.heartsArray) == 5 || len(h.diamondsArray) == 5 || len(h.clubsArray) == 5 || len(h.spadesArray) == 5 {
		return true
	}
	return false
}

func isStraight(h *Hand) bool {
	if h.rankArraySorted[0] == 2 && h.rankArraySorted[1] == 3 && h.rankArraySorted[2] == 4 && h.rankArraySorted[3] == 5 && h.rankArraySorted[4] == 14 {
		return true
	}
	for i := 0; i < (len(h.rankArraySorted) - 1); i++ {
		if !(h.rankArraySorted[i] == h.rankArraySorted[i+1]-1) {
			return false
		}
	}
	return true
}

func isThreeOfAKind(h *Hand) bool {

	for _, value := range h.rankMap {
		if value == 3 {
			return true
		}
	}
	return false
}

func isTwoPair(h *Hand) bool {
	var firstPair bool
	var secondPair bool
	for _, value := range h.rankMap {
		if value >= 3 {
			return false
		}

		if value == 2 {
			if firstPair {
				secondPair = true
				continue
			}
			firstPair = true
		}
	}
	if firstPair && secondPair {
		return true
	}
	return false

}

func isOnePair(h *Hand) bool {
	for _, value := range h.rankMap {
		if value == 2 {
			return true
		}
	}
	return false
}
