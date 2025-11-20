package main

import "sort"

type Hand struct {
	Cards           []Card
	rankMap         map[int]int
	suitMap         map[string]int
	heartsArray     []int
	diamondsArray   []int
	clubsArray      []int
	spadesArray     []int
	rankArraySorted []int
}

func UpdateRankMap(h *Hand) {
	for i := 0; i < len(h.Cards); i++ {
		currentRank := h.Cards[i].Rank
		value, exists := h.rankMap[currentRank]
		if exists {
			h.rankMap[currentRank] = value + 1
		} else {

		}
		// h.rankMap[currentRank] = h.rankMap[currentRank] + 1 // possible alternative to above code
	}
}

func UpdateSuitMap(h *Hand) {
	for i := 0; i < len(h.Cards); i++ {
		currentSuit := h.Cards[i].Suit
		value, exists := h.suitMap[currentSuit]
		if exists {
			h.suitMap[currentSuit] = value + 1
		} else {

		}
		// h.suitMap[currentSuit] = h.suitMap[currentSuit] + 1 // possible alternative to above code
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
