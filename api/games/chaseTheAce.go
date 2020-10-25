package games

import (
	"errors"
	"fmt"
)

var _ Game = (*ChaseTheAce)(nil)

// ChaseTheAce ...
type ChaseTheAce struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Deck              *Deck    `json:"deck"`
	Started           bool     `json:"started"`
	Players           []Player `json:"players"`
	LastPlayerName    string   `json:"last_player_name"`
	NextPlayerName    string   `json:"next_player_name"`
	NextPlayerHasKing bool     `json:"next_player_has_king"`
	PreviousCard      Card     `json:"previous_card"`
	GameFinished      bool     `json:"game_finished"`
	RoundOver         bool     `json:"round_over"`
	listeners         [](chan (struct{}))
}

// AddListener ...
func (ch *ChaseTheAce) AddListener(listener chan (struct{})) {
	ch.listeners = append(ch.listeners, listener)
}

// AddPlayer ...
func (ch *ChaseTheAce) AddPlayer(name string) error {
	for _, player := range ch.Players {
		if player.Name == name {
			return errors.New("Sorry a player already has this name")
		}
	}

	ch.Players = append(ch.Players, Player{
		Name:  name,
		Lives: 3,
	})

	for _, listener := range ch.listeners {
		listener <- struct{}{}
	}

	return nil
}

// GetPlayers ...
func (ch *ChaseTheAce) GetPlayers() []Player {
	return ch.Players
}

// Start ...
func (ch *ChaseTheAce) Start() error {
	ch.Started = true
	ch.GameFinished = false
	if len(ch.Players) < 2 {
		return errors.New("not enough players")
	}

	ch.Deck = NewDeck()
	ch.Deck.Shuffle()

	cardNumber := 0
	ch.LastPlayerName = ""
	ch.PreviousCard = Card{}

	for j := range ch.Players {
		ch.Players[j].Cards = append(ch.Players[j].Cards, ch.Deck[cardNumber])

		if ch.Deck[cardNumber].Value == King {
			ch.Players[j].LastCardPlayed = &ch.Deck[cardNumber]
		} else {
			ch.Players[j].LastCardPlayed = new(Card)
			ch.Players[j].LastCardPlayed.Description = "back"
		}
		ch.Deck[cardNumber].Used = true
		cardNumber++
		ch.Players[j].IsTurn = false
	}

	ch.Players[0].Dealer = true
	ch.Players[1].IsTurn = true

	nextPlayerPosition := ch.GetNextPlayerPosition()
	fmt.Printf("next player pos: [%d]\n", nextPlayerPosition)
	ch.NextPlayerName = ch.Players[nextPlayerPosition].Name
	hasKing := ch.nextPlayerHasKing()

	if hasKing && !ch.Players[nextPlayerPosition].Dealer {
		ch.NextPlayerHasKing = true
	}

	for _, listener := range ch.listeners {
		listener <- struct{}{}
	}
	return nil
}

// Turn ...
func (ch *ChaseTheAce) Turn(card Card, stick bool) {
	lastGo := false

	for i, player := range ch.Players {
		if player.IsTurn && player.Lives > 0 {
			if !player.Dealer {
				nextPlayer := ch.GetNextPlayerPosition()
				ch.Players[i].IsTurn = false
				ch.Players[nextPlayer].IsTurn = true

				if !stick {
					currentCards := ch.Players[i].Cards
					nextPlayerCards := ch.Players[nextPlayer].Cards
					ch.Players[nextPlayer].Cards = currentCards
					ch.Players[i].Cards = nextPlayerCards
				}

				nextPlayerPosition := ch.GetNextPlayerPosition()
				hasKing := ch.nextPlayerHasKing()

				if hasKing && !ch.Players[nextPlayerPosition].Dealer {
					ch.NextPlayerHasKing = true
				}

				ch.NextPlayerName = ch.Players[nextPlayerPosition].Name

			} else {
				if !stick {
					for _, card := range ch.Deck {
						if !card.Used {
							ch.Players[i].Cards[0] = card
							break
						}
					}
				}

				lastGo = true
			}
			break
		}
	}

	if lastGo {
		ch.RoundOver = true
		var playersToLoseLives []int

		lowest := King

		for i, player := range ch.Players {
			if player.Cards[0].Value < lowest {
				lowest = player.Cards[0].Value
				playersToLoseLives = []int{i}
			} else if player.Cards[0].Value == lowest {
				playersToLoseLives = append(playersToLoseLives, i)
			}

			ch.Players[i].LastCardPlayed = &player.Cards[0]
		}

		for _, playerPosition := range playersToLoseLives {
			ch.Players[playerPosition].Lives = ch.Players[playerPosition].Lives - 1
		}
	}

	ch.checkGameOver()

	for _, listener := range ch.listeners {
		listener <- struct{}{}
	}
}

// GetNextPlayerPosition ...
func (ch *ChaseTheAce) GetNextPlayerPosition() int {
	for i, player := range ch.Players {
		if player.Lives > 0 && player.IsTurn {
			if i == len(ch.Players)-1 {
				for j, pl := range ch.Players {
					if pl.Lives > 0 {
						return j
					}
				}
			}

			for j := i + 1; j <= len(ch.Players)-1; j++ {
				if ch.Players[j].Lives > 0 {
					return j
				}
			}

			for j, pl := range ch.Players {
				if pl.Lives > 0 {
					return j
				}
			}
		}
	}

	return 0
}

func (ch *ChaseTheAce) nextPlayerHasKing() bool {
	for i, player := range ch.Players {
		if player.Lives > 0 && player.IsTurn {
			if i == len(ch.Players)-1 {
				for _, pl := range ch.Players {
					if pl.Lives > 0 && pl.Cards[0].Value == King {
						return true
					}
				}
			}

			for j := i + 1; j <= len(ch.Players)-1; j++ {
				if ch.Players[j].Lives > 0 && ch.Players[j].Cards[0].Value == King {
					return true
				}
			}

			for _, pl := range ch.Players {
				if pl.Lives > 0 && pl.Cards[0].Value == King {
					return true
				}
			}
		}
	}

	return false
}

// NextRound ...
func (ch *ChaseTheAce) NextRound() error {
	ch.Started = true
	ch.GameFinished = false
	ch.RoundOver = false
	if len(ch.Players) < 2 {
		return errors.New("not enough players")
	}

	ch.Deck = NewDeck()
	ch.Deck.Shuffle()

	cardNumber := 0
	ch.LastPlayerName = ""
	ch.PreviousCard = Card{}

	currDealer := 0

	for j := range ch.Players {
		if ch.Players[j].Dealer {
			currDealer = j
		}

		if ch.Players[j].Lives > 0 {
			ch.Players[j].Cards = []Card{}
			ch.Players[j].Cards = append(ch.Players[j].Cards, ch.Deck[cardNumber])
			if ch.Deck[cardNumber].Value == King {
				ch.Players[j].LastCardPlayed = &ch.Deck[cardNumber]
			} else {
				ch.Players[j].LastCardPlayed = new(Card)
				ch.Players[j].LastCardPlayed.Description = "back"
			}
			ch.Deck[cardNumber].Used = true
			cardNumber++
			ch.Players[j].IsTurn = false
		}
	}

	dealerPosition := ch.setNextDealer(currDealer)
	ch.setFirstPlayer(dealerPosition)

	for _, listener := range ch.listeners {
		listener <- struct{}{}
	}
	return nil
}

func (ch *ChaseTheAce) setNextDealer(currentDealerPosition int) int {
	ch.Players[currentDealerPosition].Dealer = false

	if currentDealerPosition >= len(ch.Players) {
		ch.Players[0].Dealer = true
		return 0
	}

	if currentDealerPosition == len(ch.Players)-1 {
		for i := range ch.Players {
			if ch.Players[i].Lives > 0 && i != currentDealerPosition {
				ch.Players[i].Dealer = true
				return i
			}
		}
	}

	for i := currentDealerPosition + 1; i < len(ch.Players); i++ {
		if ch.Players[i].Lives > 0 && i != currentDealerPosition {
			ch.Players[i].Dealer = true
			return i
		}
	}

	for i := 0; i < len(ch.Players); i++ {
		if ch.Players[i].Lives > 0 && i != currentDealerPosition {
			ch.Players[i].Dealer = true
			return i
		}
	}

	return 0
}

func (ch *ChaseTheAce) setFirstPlayer(dealerPosition int) {
	for i, player := range ch.Players {
		player.IsTurn = false
		ch.Players[i] = player
	}

	if dealerPosition >= len(ch.Players) {
		for i, player := range ch.Players {
			if player.Lives > 0 {
				ch.Players[i].IsTurn = true
				return
			}
		}
	}

	for i := dealerPosition + 1; i < len(ch.Players); i++ {
		if ch.Players[i].Lives > 0 {
			ch.Players[i].IsTurn = true
			return
		}
	}

	for i, player := range ch.Players {
		if player.Lives > 0 {
			ch.Players[i].IsTurn = true
			return
		}
	}
}

func (ch *ChaseTheAce) checkGameOver() {
	playersStillIn := 0

	for _, player := range ch.Players {
		if player.Lives > 0 {
			playersStillIn++
		}
	}

	if playersStillIn < 2 {
		ch.GameFinished = true
	}
}
