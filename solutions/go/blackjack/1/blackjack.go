package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	//panic("Please implement the ParseCard function")
	switch card {
        case "ace":
        	return 11
        case "ten", "jack", "queen", "king":
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
	//panic("Please implement the FirstTurn function")
    card1_val := ParseCard(card1)
    card2_val := ParseCard(card2)
    sumCards := card1_val + card2_val
    dealerCard_val := ParseCard(dealerCard)
    switch {
        case card1 == "ace" && card2 == "ace":
        	return "P"
        case sumCards == 21:
        	if dealerCard_val >= 10 {
        		return "S"
            }
        	return "W"
        case sumCards >= 17 && sumCards <= 20:
        	return "S"
        case sumCards >= 12 && sumCards <= 16:
        	if dealerCard_val < 7 {
        		return "S"
            }
        	return "H"
        case sumCards <= 11:
        	return "H"
        default:
        	return "Cannot find any matching conditions"
    }
}
