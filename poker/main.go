package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	player   Player
	computer Player
	spade    = color.New(color.FgBlue, color.Bold).SprintFunc()
	clubs    = color.New(color.FgHiBlue, color.Bold).SprintFunc()
	heart    = color.New(color.FgHiRed, color.Bold).SprintFunc()
	diamond  = color.New(color.FgRed, color.Bold).SprintFunc()
	royal    = color.New(color.FgHiYellow, color.Bold).SprintFunc()
	score    = map[string]int{
		"royal flush":       200,
		"flush":             190,
		"straight":          180,
		"full house":        170,
		"4 of a kind ace":   160,
		"4 of a kind king":  155,
		"4 of a kind queen": 150,
		"4 of a kind jack":  145,
		"4 of a kind 9":     139,
		"4 of a kind 8":     138,
		"4 of a kind 7":     137,
		"4 of a kind 6":     136,
		"4 of a kind 5":     135,
		"4 of a kind 4":     134,
		"4 of a kind 3":     133,
		"4 of a kind 2":     132,
		"4 of a kind 1":     131,

		"3 of a kind ace":   130,
		"3 of a kind king":  125,
		"3 of a kind queen": 120,
		"3 of a kind jack":  125,
		"3 of a kind 9":     119,
		"3 of a kind 8":     118,
		"3 of a kind 7":     117,
		"3 of a kind 6":     116,
		"3 of a kind 5":     115,
		"3 of a kind 4":     114,
		"3 of a kind 3":     113,
		"3 of a kind 2":     112,
		"3 of a kind 1":     111,

		"high pair of a kind ace":   110,
		"high pair of a kind king":  105,
		"high pair of a kind queen": 100,
		"high pair of a kind jack":  95,
		"high pair of a kind 9":     89,
		"high pair of a kind 8":     88,
		"high pair of a kind 7":     87,
		"high pair of a kind 6":     86,
		"high pair of a kind 5":     85,
		"high pair of a kind 4":     84,
		"high pair of a kind 3":     83,
		"high pair of a kind 2":     82,
		"high pair of a kind 1":     81,

		"pair of a kind ace":   80,
		"pair of a kind king":  79,
		"pair of a kind queen": 78,
		"pair of a kind jack":  77,
		"pair of a kind 9":     76,
		"pair of a kind 8":     75,
		"pair of a kind 7":     74,
		"pair of a kind 6":     73,
		"pair of a kind 5":     72,
		"pair of a kind 4":     71,
		"pair of a kind 3":     70,
		"pair of a kind 2":     69,
		"pair of a kind 1":     68,
	}
)

func checkHand(hand Hand) (string, int) {
	// check for royal flush
	sort.Sort(hand)
	numberOfTens := 0
	numberOfRoyals := 0
	numberOfAces := 0
	allSameSuit := true
	future := 1
	for i, c := range hand {
		if c.IsRoyal {
			numberOfRoyals++
		}
		if c.IsAce {
			numberOfAces++
		}
		if c.Value == 10 && !c.IsRoyal {
			numberOfTens++
		}
		if future < len(hand) {
			if hand[i].Suit != hand[future].Suit {
				allSameSuit = false
			}
		}
		future++
	}

	if allSameSuit {
		if numberOfRoyals == 3 { // jack, queen, king
			if numberOfAces == 1 { // ace
				if numberOfTens == 1 { // 10
					return "royal flush", 0
				}
			}
		}
	}

	// check for flush
	if allSameSuit {
		flush := true
		future = 1
		for i := range hand {
			if future < len(hand) {
				if hand[i].Value+1 != hand[future].Value {
					flush = false
					break
				}
			}
		}
		if flush {
			return "flush", 0
		}
	}
	// check for straight
	straight := true
	future = 1
	for i := range hand {
		if future < len(hand) {
			if hand[i].Value != hand[future].Value {
				straight = false
				break
			}
		}
		future++
	}
	if straight {
		return "straight", 0
	}

	// full house?
	pair := false
	threeOfAKind := false

	for _, c := range hand {
		cards := findCardsByValue(hand, c.Value, c.IsRoyal, c.RoyalType)
		if len(cards) == 3 {
			threeOfAKind = true
		}
		if len(cards) == 2 {
			pair = true
		}
	}
	if pair && threeOfAKind {
		return "full house", 0
	}

	// determine if 4 of a kind, 3 or a kind, or 2 of a kind
	twoPairs := []string{}
	for _, card := range hand {
		cards := findCardsByValue(hand, card.Value, card.IsRoyal, card.RoyalType)

		if len(cards) == 4 {
			if cards[0].IsAce {
				return "4 of a kind ace", 0
			} else if cards[0].IsRoyal {
				return fmt.Sprintf("4 of a kind %s", cards[0].RoyalType), 0
			} else {
				return fmt.Sprintf("4 of a kind %d", cards[0].Value), 0
			}
		}
		if len(cards) == 3 {
			if cards[0].IsAce {
				return "3 of a kind ace", 0
			} else if cards[0].IsRoyal {
				return fmt.Sprintf("3 of a kind %s", cards[0].RoyalType), 0
			} else {
				return fmt.Sprintf("3 of a kind %d", cards[0].Value), 0
			}
		}
		if len(cards) == 2 {
			if cards[0].IsAce {
				if !inSlice(twoPairs, "pair of a kind ace") {
					twoPairs = append(twoPairs, "pair of a kind ace")
				}

			} else if cards[0].IsRoyal {
				if !inSlice(twoPairs, fmt.Sprintf("pair of a kind %s", cards[0].RoyalType)) {
					twoPairs = append(twoPairs, fmt.Sprintf("pair of a kind %s", cards[0].RoyalType))
				}
			} else {
				if !inSlice(twoPairs, fmt.Sprintf("pair of a kind %d", cards[0].Value)) {
					twoPairs = append(twoPairs, fmt.Sprintf("pair of a kind %d", cards[0].Value))
				}
			}
		}
	}
	if len(twoPairs) == 2 {
		if score[twoPairs[0]] > score[twoPairs[1]] {
			return "high " + twoPairs[0], 0
		}
		return "high " + twoPairs[1], 0
	} else if len(twoPairs) == 1 {
		return twoPairs[0], 0
	}

	// logic here is broken
	HighTwoCards := hand[len(hand)-2:]
	tot := 0
	for _, c := range HighTwoCards {
		if c.IsRoyal {
			if c.RoyalType == fmt.Sprintf("%s", royal("king")) {
				tot = tot + 13
			} else if c.RoyalType == fmt.Sprintf("%s", royal("queen")) {
				tot = tot + 12
			} else {
				tot = tot + 11
			}
		} else {
			tot = tot + c.Value
		}
	}
	return "", tot
}

func computerLogic(comp *Player, deck *Deck) (string, int) {
	// first check the hand to see of the score is high enough to do nothing
	// computer will stay with a hand of 3 of a kind or greater
	handName, baseScore := checkHand(comp.Cards)
	compScore := baseScore
	if baseScore > 0 {
		compScore = score[handName]
	}
	if compScore >= 111 {
		return handName, 0
	}
	// determine how many and which cards to RemoveCard
	// later I will go back and add probalities of how many cards to discard to maximize the handName
	// for now just to get the game working I'm going to randomize 0-5
	t := time.Now()
	source := rand.NewSource(t.Unix())
	r := rand.New(source)
	n := r.Perm(5)[0] // number of cards to drop

	selected := []int{}
	for i := 0; i < n; i++ {
		t := time.Now()
		source := rand.NewSource(t.Unix())
		r := rand.New(source)
		n := r.Perm(5)[0] + 1
		contains := func(n int, sel []int) bool {
			for _, s := range selected {
				if n == s {
					return true
				}
			}
			return false
		}(n, selected)
		if !contains {
			selected = append(selected, n)
		}
	}
	comp.RemoveCard(selected)
	for i := 0; i < len(selected); i++ {
		comp.Cards = append(comp.Cards, deck.Draw())
	}
	return checkHand(comp.Cards)
}

func findCardsByValue(hand Hand, value int, isRoyal bool, royalType string) []Card {
	cards := Hand{}
	for _, c := range hand {
		if isRoyal {
			if c.RoyalType == royalType {
				cards = append(cards, c)
			}
		} else {
			if c.Value == value {
				cards = append(cards, c)
			}
		}

	}
	return cards
}

func inSlice(slice []string, tar string) bool {
	for _, s := range slice {
		if s == tar {
			return true
		}
	}
	return false
}

func initHand(deck *Deck, p *Player, c *Player) {
	p.Cards = Hand{}
	c.Cards = Hand{}
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			p.Cards = append(p.Cards, deck.Draw())
		} else {
			c.Cards = append(c.Cards, deck.Draw())
		}
	}
}

func main() {
	deck := Deck{}
	deck.CreateDeck()
	// game loop
	for {
		deck.Shuffle()
		initHand(&deck, &player, &computer) // alternating 5 card Draw
		// main game
		for {
			fmt.Printf("Player Draw 5 cards\n")
			ShowHand(&player)
			var discard int
			for {
				fmt.Printf("How many cards would like you to discard? (0-5)\n")
				_, err := fmt.Scan(&discard)
				if err != nil {
					fmt.Printf("Please only enter a digit between 0 and 5 inclusive\n")
					continue
				}
				break
			}

			if discard > 0 {
				ShowHand(&player)
				fmt.Printf("Please choose which cards to discard (0-5)\n")
				selected := []int{}
				var input int
				for {

					if len(selected) == discard {
						player.RemoveCard(selected)
						break
					}
					_, err := fmt.Scan(&input)
					if err != nil {
						fmt.Printf("Please only enter a digit between 1 and 5 inclusive\n")
						ShowHand(&player)
						continue
					}
					selected = append(selected, input)
				}
				fmt.Printf("Picking up %d more cards\n", discard)
				for i := 0; i < discard; i++ {
					player.Cards = append(player.Cards, deck.Draw())
				}
			}
			fmt.Printf("New hand\n")
			ShowHand(&player)
			fmt.Printf("\n\n")

			player1HandName, baseScore := checkHand(player.Cards)
			player1Score := baseScore
			if player1Score == 0 {
				player1Score = score[player1HandName]
			}
			player.TotalScore = player.TotalScore + player1Score

			computerHandName, baseScore := computerLogic(&computer, &deck)
			computerScore := baseScore
			if computerScore == 0 {
				computerScore = score[computerHandName]
			}
			computer.TotalScore = computer.TotalScore + computerScore

			if computerScore > player1Score {
				fmt.Printf("Computer won with the hand: %s score: %d\n", computerHandName, computerScore)
				computer.HandsWon++
			} else {
				fmt.Printf("Player1 won with the hand: %s score: %d\n", player1HandName, player1Score)
				player.HandsWon++
			}

			// another game?
			for {
				fmt.Printf("play again? [y, n]\n")
				var action string
				fmt.Scan(&action)
				if strings.Contains(strings.ToLower(action), "y") {
					fmt.Printf("\n\n")
					break
				}
				fmt.Printf("Player1 final score: %d\nPlayer1 total hands won %d\n", player.TotalScore, player.HandsWon)
				fmt.Printf("Computer final score: %d\nComputer total hands won %d\n", computer.TotalScore, computer.HandsWon)
				return
			}
		}

	}
}
