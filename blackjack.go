package main

import (
	"fmt"
	"math/rand"
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
	suit := []string{"heart", "spade", "club", "diamond"}
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
	source := rand.NewSource(rand.Int63())
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

func (p *Player) isBlackJack() bool {
	p.totalCards()
	if p.Total == 21 {
		return true
	}
	return false
}

func main() {
	deck := Deck{}
	deck.createDeck()
	deck.shuffle()

	fmt.Println("1 player or 2 players: [1, 2]")
	var mode int
	for {
		fmt.Scan(&mode)
		if mode == 1 || mode == 2 {
			break
		}
		fmt.Println("1 player or 2 players: [1, 2]")
	}

	if mode == 1 {
		var p1 Player
		p1.Name = "Player 1"
		var p2 Player
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

			if p2.isBlackJack() {
				fmt.Println("You lost, dealer has blackjack")
				p2.Wins++
				continue
			}

			fmt.Printf("Player 1 your cards are: ")
			for i := 0; i < len(p1.Cards); i++ {
				fmt.Printf("%v of %s\n", p1.Cards[i].Value, p1.Cards[i].Suit)
			}
			p1.totalCards()
			fmt.Printf("Total: %d\n\n", p1.Total)
			if p1.isBlackJack() {
				fmt.Println("You won, you have blackjack")
				p1.Wins++
				continue
			}

			fmt.Printf("Dealer shows 1 card: %v of %s\n", p2.Cards[0].Value, p2.Cards[0].Suit)

			var action string
			var p1Stay = false
			var p2Stay = false
			for {
				fmt.Printf("What will you do, hit or stay [h, s]\n ")
				fmt.Scan(&action)
				if action == "h" || action == "s" {
					if action == "h" {
						card := deck.draw()
						p1.Cards = append(p1.Cards, card)

						p1.totalCards()
						fmt.Printf("You drew the card: %v of %s\n", card.Value, card.Suit)
						if p1.Total > 21 {
							fmt.Printf("You lost total: %d\n", p1.Total)
							p2.Wins++
							break
						}
						fmt.Printf("Your total: %d\n", p1.Total)
						if p1.isBlackJack() {
							fmt.Println("You won, you have blackjack")
							p1.Wins++
							break
						}

						if !p1Stay {
							p2.totalCards()
							if p2.Total < 16 {
								card := deck.draw()
								p2.Cards = append(p2.Cards, card)
								fmt.Printf("Dealer drew the card: %v of %s\n", card.Value, card.Suit)
								p2.totalCards()
								if p2.Total > 21 {
									fmt.Printf("You win, dealer when bust with %d\n", p2.Total)
									p1.Wins++
									break
								}
								if p2.isBlackJack() {
									fmt.Println("You lost, dealer has blackjack")
									p2.Wins++
									break
								}
							} else {
								p2Stay = true
							}
						}

					} else {
						p1Stay = true
					}

					p2.totalCards()
					if p2.Total > 21 {
						fmt.Printf("You win, dealer when bust with %d\n", p2.Total)
						p1.Wins++
						break
					}
					if p2.Total == 21 {
						fmt.Println("You lost, dealer has blackjack")
						p2.Wins++
						break
					}

					if p1Stay {
						if p2.Total < 16 {
							card := deck.draw()
							p2.Cards = append(p2.Cards, card)
							fmt.Printf("Dealer drew the card: %v of %s\n", card.Value, card.Suit)
							p2.totalCards()
							if p2.Total > 21 {
								fmt.Printf("You win, dealer when bust with %d\n", p2.Total)
								p1.Wins++
								break
							}
							if p2.isBlackJack() {
								fmt.Println("You lost, dealer has blackjack")
								p2.Wins++
								break
							}
						} else {
							p2Stay = true
						}
					}

					if p1Stay && p2Stay {
						p1.totalCards()
						p2.totalCards()
						if p1.Total > p2.Total {
							fmt.Printf("You win, dealer had: %d and you had: %d \n", p2.Total, p1.Total)
							p1.Wins++
						} else {
							fmt.Printf("You loss, dealer had: %d and you had: %d \n", p2.Total, p1.Total)
							p2.Wins++
						}
						break
					}
				} else {
					fmt.Println("hit or stay [h, s]")
					continue
				}

			}
			fmt.Println("play again?")
			for {
				fmt.Scan(&action)
				if action == "y" {
					println("")
					break
				}
				fmt.Printf("You won %d games.  The dealer won %d games\n", p1.Wins, p2.Wins)
				return
			}
		}
	} else {
		// TODO
	}

}
