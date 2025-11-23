package main

import "testing"

func TestRoyalFlush(t *testing.T) {

	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "C", Rank: 10},
		{Suit: "C", Rank: 11},
		{Suit: "C", Rank: 12},
		{Suit: "C", Rank: 13},
		{Suit: "C", Rank: 14},
	}
	// UpdateHandArrays(myHand) // Running prep functions
	organizeHand(myHand)
	got := isRoyalFlush(myHand)
	want := true
	handRankingString := handRanking(myHand)

	// Assert
	if got != want {
		t.Errorf("isRoyalFlush() = %v, want %v and handranking is: %s", got, want, handRankingString)
	}

}

func TestStraightFlush1(t *testing.T) {

	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "D", Rank: 8},
		{Suit: "D", Rank: 12},
		{Suit: "D", Rank: 11},
		{Suit: "D", Rank: 10},
		{Suit: "D", Rank: 9},
	}
	// UpdateHandArrays(myHand) // Running prep functions
	organizeHand(myHand)
	got := isStraightFlush(myHand)
	want := true
	handRankingString := handRanking(myHand)

	// Assert
	if got != want {
		t.Errorf("isStraightFlush() = %v, want %v and handranking is: %s", got, want, handRankingString)
	}

}

func TestStraightFlush2(t *testing.T) {

	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "H", Rank: 8},
		{Suit: "H", Rank: 10},
		{Suit: "H", Rank: 11},
		{Suit: "H", Rank: 7},
		{Suit: "H", Rank: 9},
	}
	// UpdateHandArrays(myHand) // Running prep functions
	organizeHand(myHand)
	got := isStraightFlush(myHand)
	want := true
	handRankingString := handRanking(myHand)

	// Assert
	if got != want {
		t.Errorf("isStraightFlush() = %v, want %v and handranking is: %s", got, want, handRankingString)
	}

}

func TestPoker1(t *testing.T) {

	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "H", Rank: 10},
		{Suit: "S", Rank: 12},
		{Suit: "S", Rank: 10},
		{Suit: "D", Rank: 10},
		{Suit: "C", Rank: 10},
	}
	// UpdateHandArrays(myHand) // Running prep functions
	organizeHand(myHand)
	got := isFourOfAKind(myHand)
	want := true
	handRankingString := handRanking(myHand)

	// Assert
	if got != want {
		t.Errorf("isFourOfAKind() = %v, want %v and handranking is: %s", got, want, handRankingString)
	}

}
func TestPoker2(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "H", Rank: 10},
		{Suit: "S", Rank: 13},
		{Suit: "S", Rank: 10},
		{Suit: "D", Rank: 10},
		{Suit: "C", Rank: 10},
	}
	// UpdateHandArrays(myHand) // Running prep functions
	organizeHand(myHand)
	got := isFourOfAKind(myHand)
	want := true
	handRankingString := handRanking(myHand)

	// Assert
	if got != want {
		t.Errorf("isFourOfAKind() = %v, want %v and handranking is: %s", got, want, handRankingString)
	}

}
func TestPoker3(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "H", Rank: 8},
		{Suit: "S", Rank: 12},
		{Suit: "S", Rank: 8},
		{Suit: "D", Rank: 8},
		{Suit: "C", Rank: 8},
	}
	// UpdateHandArrays(myHand) // Running prep functions
	organizeHand(myHand)
	got := isFourOfAKind(myHand)
	want := true
	handRankingString := handRanking(myHand)

	// Assert
	if got != want {
		t.Errorf("isFourOfAKind() = %v, want %v and handranking is: %s", got, want, handRankingString)
	}

}
func TestPoker4(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "H", Rank: 7},
		{Suit: "S", Rank: 13},
		{Suit: "S", Rank: 7},
		{Suit: "D", Rank: 7},
		{Suit: "C", Rank: 7},
	}
	// UpdateHandArrays(myHand) // Running prep functions
	organizeHand(myHand)
	got := isFourOfAKind(myHand)
	want := true
	handRankingString := handRanking(myHand)

	// Assert
	if got != want {
		t.Errorf("isFourOfAKind() = %v, want %v and handranking is: %s", got, want, handRankingString)
	}

}

func TestFullHouse1(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "H", Rank: 2},
		{Suit: "S", Rank: 12},
		{Suit: "C", Rank: 2},
		{Suit: "D", Rank: 2},
		{Suit: "C", Rank: 12},
	}
	// UpdateHandArrays(myHand) // Running prep functions
	organizeHand(myHand)
	got := isFullHouse(myHand)
	want := true
	handRankingString := handRanking(myHand)

	// Assert
	if got != want {
		t.Errorf("isFullHouse() = %v, want %v and handranking is: %s", got, want, handRankingString)
	}

}

func TestFullHouse2(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "H", Rank: 2},
		{Suit: "S", Rank: 11},
		{Suit: "C", Rank: 2},
		{Suit: "D", Rank: 2},
		{Suit: "C", Rank: 11},
	}
	// UpdateHandArrays(myHand) // Running prep functions
	organizeHand(myHand)
	got := isFullHouse(myHand)
	want := true
	handRankingString := handRanking(myHand)

	// Assert
	if got != want {
		t.Errorf("isFullHouse() = %v, want %v and handranking is: %s", got, want, handRankingString)
	}

}

func TestFlush1(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "H", Rank: 13},
		{Suit: "H", Rank: 12},
		{Suit: "H", Rank: 2},
		{Suit: "H", Rank: 4},
		{Suit: "H", Rank: 5},
	}
	// UpdateHandArrays(myHand) // Running prep functions
	organizeHand(myHand)
	got := isFlush(myHand)
	want := true
	handRankingString := handRanking(myHand)

	// Assert
	if got != want {
		t.Errorf("isFlush() = %v, want %v and handranking is: %s", got, want, handRankingString)
	}

}

func TestFlush2(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "D", Rank: 5},
		{Suit: "D", Rank: 4},
		{Suit: "D", Rank: 2},
		{Suit: "D", Rank: 12},
		{Suit: "D", Rank: 13},
	}
	// UpdateHandArrays(myHand) // Running prep functions
	organizeHand(myHand)
	got := isFlush(myHand)
	want := true
	handRankingString := handRanking(myHand)

	// Assert
	if got != want {
		t.Errorf("isFlush() = %v, want %v and handranking is: %s", got, want, handRankingString)
	}

}

func TestStraight1(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "H", Rank: 3},
		{Suit: "S", Rank: 7},
		{Suit: "H", Rank: 5},
		{Suit: "D", Rank: 6},
		{Suit: "H", Rank: 4},
	}
	// UpdateHandArrays(myHand)
	organizeHand(myHand)

	got := isStraight(myHand)
	want := true
	handRankingString := handRanking(myHand)

	if got != want {
		t.Errorf("isStraight() = %v, want %v (rank: %s)", got, want, handRankingString)
	}
}

func TestStraight2(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "C", Rank: 9},
		{Suit: "C", Rank: 10},
		{Suit: "S", Rank: 11},
		{Suit: "D", Rank: 7},
		{Suit: "H", Rank: 8},
	}
	// UpdateHandArrays(myHand)
	organizeHand(myHand)

	got := isStraight(myHand)
	want := true
	handRankingString := handRanking(myHand)

	if got != want {
		t.Errorf("isStraight() = %v, want %v (rank: %s)", got, want, handRankingString)
	}
}

func TestStraight3(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "H", Rank: 4},
		{Suit: "S", Rank: 5},
		{Suit: "H", Rank: 14}, // Ace
		{Suit: "D", Rank: 3},
		{Suit: "H", Rank: 2},
	}
	// UpdateHandArrays(myHand)
	organizeHand(myHand)

	got := isStraight(myHand)
	want := true
	handRankingString := handRanking(myHand)

	if got != want {
		t.Errorf("isStraight() = %v, want %v (rank: %s)", got, want, handRankingString)
	}
}

/* func TestStraight(t *testing.T) {
	testCases := []struct {
		name    string
		hand    []Card
		expect  bool
	}{
		{
			name: "Normal Straight",
			hand: []Card{
				{Suit: "H", Rank: 6},
				{Suit: "D", Rank: 7},
				{Suit: "C", Rank: 8},
				{Suit: "S", Rank: 9},
				{Suit: "H", Rank: 10},
			},
			expect: true,
		},
		{
			name: "Wheel Straight (A-5)",
			hand: []Card{
				{Suit: "H", Rank: 14}, // Ace
				{Suit: "D", Rank: 2},
				{Suit: "C", Rank: 3},
				{Suit: "S", Rank: 4},
				{Suit: "H", Rank: 5},
			},
			expect: true,
		},
		{
			name: "Not a Straight",
			hand: []Card{
				{Suit: "H", Rank: 2},
				{Suit: "D", Rank: 4},
				{Suit: "C", Rank: 6},
				{Suit: "S", Rank: 8},
				{Suit: "H", Rank: 10},
			},
			expect: false,
		},
		{
			name: "Hand with a Pair",
			hand: []Card{
				{Suit: "H", Rank: 7},
				{Suit: "D", Rank: 7}, // The pair
				{Suit: "C", Rank: 8},
				{Suit: "S", Rank: 9},
				{Suit: "H", Rank: 10},
			},
			expect: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isStraight(tc.hand)
			if got != tc.expect {
				t.Errorf("isStraight(%v): expected %v, got %v",
					tc.hand, tc.expect, got)
			}
		})
	}
} */

func TestThreeOfAKind1(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "H", Rank: 2},
		{Suit: "S", Rank: 12},
		{Suit: "S", Rank: 2},
		{Suit: "D", Rank: 2},
		{Suit: "C", Rank: 13},
	}
	// UpdateHandArrays(myHand)
	organizeHand(myHand)

	got := isThreeOfAKind(myHand)
	want := true
	handRankingString := handRanking(myHand)

	if got != want {
		t.Errorf("isThreeOfAKind() = %v, want %v (rank: %s)", got, want, handRankingString)
	}
}

func TestThreeOfAKind2(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "H", Rank: 2},
		{Suit: "S", Rank: 7},
		{Suit: "S", Rank: 2},
		{Suit: "D", Rank: 2},
		{Suit: "C", Rank: 9},
	}
	// UpdateHandArrays(myHand)
	organizeHand(myHand)

	got := isThreeOfAKind(myHand)
	want := true
	handRankingString := handRanking(myHand)

	if got != want {
		t.Errorf("isThreeOfAKind() = %v, want %v (rank: %s)", got, want, handRankingString)
	}
}

func TestThreeOfAKind3(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "H", Rank: 2},
		{Suit: "S", Rank: 8},
		{Suit: "S", Rank: 2},
		{Suit: "D", Rank: 2},
		{Suit: "C", Rank: 9},
	}
	// UpdateHandArrays(myHand)
	organizeHand(myHand)

	got := isThreeOfAKind(myHand)
	want := true
	handRankingString := handRanking(myHand)

	if got != want {
		t.Errorf("isThreeOfAKind() = %v, want %v (rank: %s)", got, want, handRankingString)
	}
}

func TestTwoPair1(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "H", Rank: 5},
		{Suit: "S", Rank: 12},
		{Suit: "C", Rank: 5},
		{Suit: "D", Rank: 10},
		{Suit: "C", Rank: 10},
	}
	// UpdateHandArrays(myHand)
	organizeHand(myHand)

	got := isTwoPair(myHand)
	want := true
	handRankingString := handRanking(myHand)

	if got != want {
		t.Errorf("isTwoPair() = %v, want %v (rank: %s)", got, want, handRankingString)
	}
}

func TestTwoPair2(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "H", Rank: 9},
		{Suit: "S", Rank: 12},
		{Suit: "C", Rank: 9},
		{Suit: "D", Rank: 10},
		{Suit: "C", Rank: 10},
	}
	// UpdateHandArrays(myHand)
	organizeHand(myHand)

	got := isTwoPair(myHand)
	want := true
	handRankingString := handRanking(myHand)

	if got != want {
		t.Errorf("isTwoPair() = %v, want %v (rank: %s)", got, want, handRankingString)
	}
}

func TestOnePair1(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "H", Rank: 3},
		{Suit: "S", Rank: 8},
		{Suit: "H", Rank: 5},
		{Suit: "D", Rank: 8},
		{Suit: "C", Rank: 14},
	}
	// UpdateHandArrays(myHand)
	organizeHand(myHand)

	got := isOnePair(myHand)
	want := true
	handRankingString := handRanking(myHand)

	if got != want {
		t.Errorf("isOnePair() = %v, want %v (rank: %s)", got, want, handRankingString)
	}
}

func TestOnePair2(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "S", Rank: 4},
		{Suit: "D", Rank: 14},
		{Suit: "H", Rank: 3},
		{Suit: "C", Rank: 14},
		{Suit: "H", Rank: 10},
	}
	// UpdateHandArrays(myHand)
	organizeHand(myHand)

	got := isOnePair(myHand)
	want := true
	handRankingString := handRanking(myHand)

	if got != want {
		t.Errorf("isOnePair() = %v, want %v (rank: %s)", got, want, handRankingString)
	}
}

func TestHighCard1(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "H", Rank: 3},
		{Suit: "S", Rank: 8},
		{Suit: "H", Rank: 5},
		{Suit: "D", Rank: 13},
		{Suit: "C", Rank: 14},
	}
	// UpdateHandArrays(myHand)
	organizeHand(myHand)

	got := handRanking(myHand)
	want := "High Card"

	if got != want {
		t.Errorf("handRanking() = %v, want %v", got, want)
	}
}

func TestHighCard2(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "H", Rank: 3},
		{Suit: "S", Rank: 8},
		{Suit: "H", Rank: 5},
		{Suit: "D", Rank: 13},
		{Suit: "C", Rank: 10},
	}
	// UpdateHandArrays(myHand)
	organizeHand(myHand)

	got := handRanking(myHand)
	want := "High Card"

	if got != want {
		t.Errorf("handRanking() = %v, want %v", got, want)
	}
}

func TestHighCard3(t *testing.T) {
	myHand := NewHand()
	myHand.Cards = []Card{
		{Suit: "H", Rank: 3},
		{Suit: "S", Rank: 8},
		{Suit: "H", Rank: 5},
		{Suit: "D", Rank: 13},
		{Suit: "C", Rank: 2},
	}
	// UpdateHandArrays(myHand)
	organizeHand(myHand)

	got := handRanking(myHand)
	want := "High Card"

	if got != want {
		t.Errorf("handRanking() = %v, want %v", got, want)
	}
}

/* // royal flush
"CT CJ CQ CK CA"

// straight flush
"D8 DQ DJ DT D9"
"H8 HT HJ H7 H9"

// poker
"HT SQ ST DT CT"
"HT SK ST DT CT"
"H8 SQ S8 D8 C8"
"H7 SK S7 D7 C7"

// full house
"H2 SQ C2 D2 CQ"
"H2 SJ C2 D2 CJ"

// flush
"HK HQ H2 H4 H5"
"D5 D4 D2 DQ DK"

// straight
"H3 S7 H5 D6 H4"
"C9 CT SJ D7 H8"
"H4 S5 HA D3 H2"


// three of a kind
"H2 SQ S2 D2 CK"
"H2 S7 S2 D2 C9",
"H2 S8 S2 D2 C9"

// two pairs
"H5 SQ C5 DT CT"
"H9 SQ C9 DT CT"

// one pair
"H3 S8 H5 D8 CA"
"S4 DA H3 CA HT"

// high card
"H3 S8 H5 DK CA"
"H3 S8 H5 DK CT"
"H3 S8 H5 DK C2"
*/
