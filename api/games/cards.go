package games

import (
	"math/rand"
	"time"
)

// Card ...
type Card struct {
	Suit  Suit
	Value CardValue
	Used  bool
}

// Deck ...
type Deck [52]Card

// Suit is a suit
type Suit int

const (
	// Club represents a club
	Club Suit = iota
	// Spade represents a spade
	Spade
	// Heart represents a heart
	Heart
	// Diamond represents a diamon
	Diamond
)

var suits = [...]string{
	"c",
	"s",
	"h",
	"d",
}

func (s Suit) String() string {
	return suits[s]
}

// CardValue ...
type CardValue int

const (
	// Ace ...
	Ace CardValue = iota
	// Two ...
	Two
	// Three ...
	Three
	// Four ...
	Four
	// Five ...
	Five
	// Six ...
	Six
	// Seven ...
	Seven
	// Eight ...
	Eight
	// Nine ...
	Nine
	// Ten ...
	Ten
	// Jack ...
	Jack
	// Queen ...
	Queen
	// King ...
	King
)

var cardValues = [...]string{
	"a",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"10",
	"j",
	"q",
	"k",
}

func (c CardValue) String() string {
	return cardValues[c]
}

// Descriptor ...
func (c Card) Descriptor() string {
	return c.Value.String() + c.Suit.String()
}

// NewDeck ...
func NewDeck() *Deck {
	var deck Deck

	i := 0

	for value := range cardValues {
		for suit := range suits {
			deck[i] = Card{
				Suit:  Suit(suit),
				Value: CardValue(value),
			}

			i++
		}
	}

	return &deck
}

// GetCardFromDescriptor ...
func GetCardFromDescriptor(descriptor string) Card {
	if len(descriptor) < 2 {
		return Card{}
	}

	suitStr := descriptor[len(descriptor)-1:]
	valueStr := descriptor[0 : len(descriptor)-1]

	card := Card{}

	for i, val := range suits {
		if val == suitStr {
			card.Value = CardValue(i)
		}
	}

	for i, val := range cardValues {
		if val == valueStr {
			card.Suit = Suit(i)
		}
	}

	return card
}

// Shuffle ...
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}
