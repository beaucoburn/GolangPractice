// Package blackjack implements the BlackJack card game, but is not a real Go package.
package blackjack

func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "two":
		return 2
	case card == "three":
		return 3
	case card == "four":
		return 4
	case card == "five":
		return 5
	case card == "six":
		return 6
	case card == "seven":
		return 7
	case card == "eight":
		return 8
	case card == "nine":
		return 9
	case card == "ten":
		return 10
	case card == "jack":
		return 10
	case card == "queen":
		return 10
	case card == "king":
		return 10
	default:
		return 0
	}
}

func FirstTurn(card1, card2, dealerCard string) string {

	if ParseCard(card1)+ParseCard(card2) == 22 {
		return "P"
	} else if ParseCard(card1)+ParseCard(card2) == 21 && ParseCard(dealerCard) > 10 {
		return "S"
	} else if ParseCard(card1)+ParseCard(card2) == 21 && ParseCard(dealerCard) <= 9 {
		return "W"
	} else if ParseCard(card1)+ParseCard(card2) <= 16 && ParseCard(dealerCard) > 6 {
		return "H"
	} else if ParseCard(card1)+ParseCard(card2) <= 15 && ParseCard(dealerCard) == 6 {
		return "S"
	} else if ParseCard(card1)+ParseCard(card2) >= 16 {
		return "S"
	} else if ParseCard(card1)+ParseCard(card2) <= 10 {
		return "H"
	} else if ParseCard(card1)+ParseCard(card2) < 12 {
		return "H"
	} else {
		return "Error"
	}
}
