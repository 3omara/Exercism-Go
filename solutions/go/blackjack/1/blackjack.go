package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "jack", "queen", "king":
		return 10
	case "ace":
		return 11
	case "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	stand := "S"
	hit := "H"
	split := "P"
	auto_win := "W"

	c1, c2, d := ParseCard(card1), ParseCard(card2), ParseCard(dealerCard)

	switch {
	case c1 == 11 && c2 == 11:
		return split
	case c1+c2 == 21 && d < 10:
		if d < 10 {
			return auto_win
		}
		return stand
	case c1+c2 >= 17 && c1+c2 <= 20:
		return stand
	case c1+c2 >= 12 && c1+c2 <= 16:
		if d >= 7 {
			return hit
		}
		return stand

	case c1+c2 <= 11:
		return hit

	default:
		return stand
	}

}
