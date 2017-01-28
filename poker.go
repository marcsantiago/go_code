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

		"pair of a kind ace":   110,
		"pair of a kind king":  105,
		"pair of a kind queen": 100,
		"pair of a kind jack":  95,
		"pair of a kind 9":     89,
		"pair of a kind 8":     88,
		"pair of a kind 7":     87,
		"pair of a kind 6":     86,
		"pair of a kind 5":     85,
		"pair of a kind 4":     84,
		"pair of a kind 3":     83,
		"pair of a kind 2":     82,
		"pair of a kind 1":     81,
	}
)

// Deck ...
type Deck struct {
	Cards []Card
}

// Hand ...
type Hand []Card

func (slice Hand) Len() int {
	return len(slice)
}

func (slice Hand) Less(i, j int) bool {
	return slice[i].Value < slice[j].Value
}

func (slice Hand) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Card ...
type Card struct {
	Suit      string
	Value     int
	IsRoyal   bool
	RoyalType string
	IsAce     bool
}

// Player ...
type Player struct {
	Name  string
	Total int
	Wins  int
	Cards Hand
}

func (d *Deck) createDeck() {

	suit := []string{fmt.Sprintf("%s", heart("hearts")), fmt.Sprintf("%s", spade("spades")), fmt.Sprintf("%s", clubs("clubs")), fmt.Sprintf("%s", diamond("diamonds"))}
	royal := []string{fmt.Sprintf("%s", royal("queen")), fmt.Sprintf("%s", royal("jack")), fmt.Sprintf("%s", royal("king"))}

	// add reg cards
	for i := 1; i < 10; i++ {
		for _, s := range suit {
			d.Cards = append(d.Cards, Card{
				Suit:      s,
				Value:     i,
				IsRoyal:   false,
				RoyalType: "",
				IsAce:     false,
			})
		}
	}

	for _, r := range royal {
		for _, s := range suit {
			d.Cards = append(d.Cards, Card{
				Suit:      s,
				Value:     10,
				IsRoyal:   true,
				RoyalType: r,
				IsAce:     false,
			})
		}
	}

	for _, s := range suit {
		d.Cards = append(d.Cards, Card{
			Suit:      s,
			Value:     1,
			IsRoyal:   false,
			RoyalType: "",
			IsAce:     true,
		})
	}
}

func (d *Deck) shuffle() {
	t := time.Now()
	source := rand.NewSource(t.Unix())
	r := rand.New(source)
	arr := r.Perm(52)
	for i, a := range arr {
		d.Cards[i] = d.Cards[a]
	}
}

func (d *Deck) draw() Card {
	if len(d.Cards) != 0 {
		card := d.Cards[0]
		d.Cards = d.Cards[1:]
		return card
	}
	d.createDeck()
	d.shuffle()
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

func initHand(deck *Deck, p *Player, c *Player) {
	p.Cards = Hand{}
	c.Cards = Hand{}
	for i := 0; i < 10; i++ {
		card := deck.draw()
		if i%2 == 0 {
			p.Cards = append(p.Cards, card)
		} else {
			c.Cards = append(c.Cards, card)
		}
	}
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
			if hand[i].Value+1 != hand[future].Value {
				straight = false
				break
			}
		}
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
	for _, card := range hand {
		cards := findCardsByValue(hand, card.Value, card.IsRoyal, card.RoyalType)
		// fmt.Printf("card: %v  cards: %v  len: %d\n", c, cards, len(cards))
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
				return "pair of a kind ace", 0
			} else if cards[0].IsRoyal {
				return fmt.Sprintf("pair of a kind %s", cards[0].RoyalType), 0
			} else {
				return fmt.Sprintf("pair of a kind %d", cards[0].Value), 0
			}
		}
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
func showHand(p *Player) {
	sort.Sort(p.Cards)
	for n, c := range p.Cards {
		if c.IsRoyal {
			fmt.Printf("%d) %v of %v\n", n+1, c.RoyalType, c.Suit)
		} else if c.IsAce {
			fmt.Printf("%d) %v of %v\n", n+1, "ace", c.Suit)
		} else {
			fmt.Printf("%d) %v of %v\n", n+1, c.Value, c.Suit)
		}

	}
}

func (p *Player) removeCard(selected []int) {
	cardsToRemove := []Card{}
	for _, s := range selected {
		cardsToRemove = append(cardsToRemove, p.Cards[s-1])
	}
	for _, card := range cardsToRemove {
		deleted := 0
		for i := range p.Cards {
			j := i - deleted
			if p.Cards[j] == card {
				p.Cards = p.Cards[:j+copy(p.Cards[j:], p.Cards[j+1:])]
				fmt.Printf("here: %v   %d\n\n", p.Cards, len(p.Cards))
				deleted++
			}
		}
	}
}

func main() {
	deck := Deck{}
	deck.createDeck()
	deck.shuffle()
	// game loop
	for {
		initHand(&deck, &player, &computer) // alternating 5 card draw
		// main game
		for {
			fmt.Printf("Player draw 5 cards\n")
			showHand(&player)
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
				showHand(&player)
				fmt.Printf("Please choose which cards to discard (0-5)\n")
				selected := []int{}
				var input int
				for {

					if len(selected) == discard {
						player.removeCard(selected)
						break
					}
					_, err := fmt.Scan(&input)
					if err != nil {
						fmt.Printf("Please only enter a digit between 1 and 5 inclusive\n")
						showHand(&player)
						continue
					}
					selected = append(selected, input)
				}
				fmt.Printf("Picking up %d more cards\n", discard)
				for i := 0; i < discard; i++ {
					player.Cards = append(player.Cards, deck.draw())
				}
			}
			fmt.Printf("New hand\n")
			showHand(&player)
			break
		}

		// another game?
		for {
			fmt.Println("play again? [y, n]")
			var action string
			fmt.Scan(&action)
			if strings.Contains(strings.ToLower(action), "y") {
				fmt.Printf("\n\n")
				break
			}
			return
		}
	}

	// ph, _ := checkHand(player.Cards)
	// ch, _ := checkHand(computer.Cards)
	//
	// for _, c := range player.Cards {
	// 	fmt.Println(c)
	// }
	// fmt.Printf("hand: %s score: %d\n\n", ph, score[ph])
	// for _, c := range computer.Cards {
	// 	fmt.Println(c)
	// }
	// fmt.Printf("hand: %s score: %d\n\n", ch, score[ch])

}
