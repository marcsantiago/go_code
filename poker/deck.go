package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Card ...
type Card struct {
	Suit      string
	Value     int
	IsRoyal   bool
	RoyalType string
	IsAce     bool
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

// Deck ...
type Deck struct {
	Cards []Card
}

// CreateDeck ...
func (d *Deck) CreateDeck() {

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

// Draw  ...
func (d *Deck) Draw() Card {
	if len(d.Cards) != 0 {
		card := d.Cards[0]
		d.Cards = d.Cards[1:]
		return card
	}
	d.CreateDeck()
	d.Shuffle()
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

// Shuffle ...
func (d *Deck) Shuffle() {
	t := time.Now()
	source := rand.NewSource(t.Unix())
	r := rand.New(source)
	arr := r.Perm(52)
	for i, a := range arr {
		d.Cards[i] = d.Cards[a]
	}
}
