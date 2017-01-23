package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Deck ...
type Deck struct {
	Cards []Card
}

// Card ...
type Card struct {
	Suit    string
	Value   interface{}
	IsRoyal bool
	IsAce   bool
}

// Player ...
type Player struct {
	Name  string
	Total int
	Wins  int
	Cards []Card
}

func (d *Deck) createDeck() {
	suit := []string{"hearts", "spades", "clubs", "diamonds"}
	royal := []string{"queen", "jack", "king"}

	// add reg cards
	for i := 1; i < 10; i++ {
		for _, s := range suit {
			d.Cards = append(d.Cards, Card{
				Suit:    s,
				Value:   i,
				IsRoyal: false,
				IsAce:   false,
			})
		}
	}

	for _, r := range royal {
		for _, s := range suit {
			d.Cards = append(d.Cards, Card{
				Suit:    s,
				Value:   r,
				IsRoyal: true,
				IsAce:   false,
			})
		}
	}

	for _, s := range suit {
		d.Cards = append(d.Cards, Card{
			Suit:    s,
			Value:   "ace",
			IsRoyal: false,
			IsAce:   true,
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

func (p *Player) totalCards() {
	var aces int
	p.Total = 0
	for _, c := range p.Cards {
		if c.IsAce {
			aces++
		} else {
			var value int
			if c.IsRoyal {
				value = 10
			} else {
				value = c.Value.(int)
			}
			p.Total = p.Total + value
		}
	}
	if aces > 0 {
		for i := 0; i < aces; i++ {
			if p.Total+11 <= 21 {
				p.Total = p.Total + 11
			} else {
				p.Total = p.Total + 1
			}
		}
	}
}

func (p *Player) hasBlackJack() bool {
	p.totalCards()
	if p.Total == 21 {
		return true
	}
	return false
}

func main() {
	p1Color := color.New(color.FgHiBlue, color.Bold).SprintFunc()
	p2Color := color.New(color.FgHiMagenta, color.Bold).SprintFunc()

	win := color.New(color.FgGreen, color.Bold)
	cardColor := color.New(color.FgYellow, color.Bold).SprintFunc()
	lose := color.New(color.FgRed, color.Bold)

	deck := Deck{}
	deck.createDeck()
	deck.shuffle()

	var p1 Player
	var p2 Player
	p1.Name = "Player 1"
	p2.Name = "Computer"

	for {
		p1.Total = 0
		p1.Cards = []Card{}
		p2.Total = 0
		p2.Cards = []Card{}

		p1.Cards = append(p1.Cards, deck.draw())
		p2.Cards = append(p2.Cards, deck.draw())
		p1.Cards = append(p1.Cards, deck.draw())
		p2.Cards = append(p2.Cards, deck.draw())

		if p2.hasBlackJack() {
			lose.Printf("You lost, dealer has blackjack\n\n")
			p2.Wins++
			continue
		}

		fmt.Printf("%s your cards are: ", p1Color(p1.Name))
		for i := 0; i < len(p1.Cards); i++ {
			fmt.Printf("%v of %s, ", cardColor(p1.Cards[i].Value), cardColor(p1.Cards[i].Suit))
		}
		fmt.Printf("\n")

		p1.totalCards()
		fmt.Printf("Total: %d\n\n", p1.Total)
		if p1.hasBlackJack() {
			win.Printf("You won, you have blackjack\n\n")
			p1.Wins++
			continue
		}

		fmt.Printf("%s shows 1 card: %v of %s\n", p2Color(p2.Name), cardColor(p2.Cards[0].Value), cardColor(p2.Cards[0].Suit))

		var action string
		var p1Stay = false
		var p2Stay = false
		for {
			if !p1Stay {
				fmt.Printf("What will you do, hit or stay [h, s]\n")
				fmt.Scan(&action)
				action = strings.ToLower(action)
				if strings.Contains(action, "h") || strings.Contains(action, "s") {
					if strings.Contains(action, "h") {
						card := deck.draw()
						p1.Cards = append(p1.Cards, card)
						p1.totalCards()
						fmt.Printf("You drew the card: %v of %s\n", cardColor(card.Value), cardColor(card.Suit))
						if p1.Total > 21 {
							lose.Printf("You lost total: %d\n\n", p1.Total)
							p2.Wins++
							break
						}
						fmt.Printf("Your total: %d\n\n", p1.Total)
						if p1.hasBlackJack() {
							win.Printf("You won, you have blackjack\n\n")
							p1.Wins++
							break
						}
					} else {
						p1Stay = true
					}
				}
			}

			p2.totalCards()
			if p2.Total < 16 {
				card := deck.draw()
				p2.Cards = append(p2.Cards, card)
				fmt.Printf("%s drew the card: %v of %s\n\n", p2Color(p2.Name), cardColor(card.Value), cardColor(card.Suit))
				p2.totalCards()
				if p2.Total > 21 {
					win.Printf("You win, %s when bust with %d\n\n", p2Color(p2.Name), p2.Total)
					p1.Wins++
					break
				}
				if p2.hasBlackJack() {
					lose.Printf("You lost, %s has blackjack\n\n", p2Color(p2.Name))
					p2.Wins++
					break
				}
			} else {
				p2Stay = true
			}

			p2.totalCards()
			if p2.Total > 21 {
				win.Printf("You win, dealer when bust with %d\n\n", p2.Total)
				p1.Wins++
				break
			}
			if p2.Total == 21 {
				lose.Printf("You lost, dealer has blackjack\n\n")
				p2.Wins++
				break
			}

			if p1Stay && p2Stay {
				p1.totalCards()
				p2.totalCards()
				if p1.Total > p2.Total {
					win.Printf("You win, dealer had: %d and you had: %d \n\n", p2.Total, p1.Total)
					p1.Wins++
				} else {
					lose.Printf("You lost, dealer had: %d and you had: %d \n\n", p2.Total, p1.Total)
					p2.Wins++
				}
				break
			}
		}

		fmt.Println("play again? [y, n]")
		for {
			fmt.Scan(&action)
			if strings.Contains(strings.ToLower(action), "y") {
				fmt.Printf("\n\n")
				break
			}
			win.Printf("You won %d games.  The %s won %d games\n", p1.Wins, p2Color(p2.Name), p2.Wins)
			return
		}
	}
}
