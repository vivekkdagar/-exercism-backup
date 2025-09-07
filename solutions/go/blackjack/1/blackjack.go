package blackjack

// ParseCard returns the integer value of a card following blackjack rules.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "jack", "queen", "king", "ten":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	switch {
	case playerScore == 22:
		return "P" // Player has two aces, always split.
	case playerScore == 21:
		if dealerScore >= 10 {
			return "S" // Stand if dealer has 10, jack, queen, king, or ace.
		}
		return "W" // Win if dealer has any other card.
	case playerScore >= 17 && playerScore <= 20:
		return "S" // Stand on 17-20.
	case playerScore >= 12 && playerScore <= 16:
		if dealerScore >= 7 {
			return "H" // Hit if dealer has 7 or more.
		}
		return "S" // Stand otherwise.
	default: // playerScore <= 11
		return "H" // Hit on 11 or less.
	}
}