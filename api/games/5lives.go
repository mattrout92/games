package games

import (
	"errors"
	"fmt"
)

var _ Game = (*FiveLives)(nil)

// FiveLives ...
type FiveLives struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Deck           *Deck    `json:"-"`
	Started        bool     `json:"started"`
	Players        []Player `json:"players"`
	LastPlayerName string   `json:"last_player_name"`
	PreviousCard   Card     `json:"previous_card"`
	LivesToLose    int      `json:"lives_to_lose"`
	GameFinished   bool     `json:"game_finished"`
	listeners      [](chan (struct{}))
}

// AddListener ...
func (fl *FiveLives) AddListener(listener chan (struct{})) {
	fl.listeners = append(fl.listeners, listener)
}

// AddPlayer ...
func (fl *FiveLives) AddPlayer(name string) error {
	for _, player := range fl.Players {
		if player.Name == name {
			return errors.New("Sorry a player already has this name")
		}
	}

	fl.Players = append(fl.Players, Player{
		Name:  name,
		Lives: 5,
	})

	for _, listener := range fl.listeners {
		listener <- struct{}{}
	}

	return nil
}

// GetPlayers ...
func (fl *FiveLives) GetPlayers() []Player {
	return fl.Players
}

// Start ...
func (fl *FiveLives) Start() error {
	fl.Started = true
	fl.GameFinished = false
	if len(fl.Players) < 2 {
		return errors.New("not enough players")
	}

	fl.Deck = NewDeck()
	fl.Deck.Shuffle()

	cardNumber := 0
	fl.LivesToLose = 1
	fl.LastPlayerName = ""
	fl.PreviousCard = Card{}
	fl.LivesToLose = 1

	for j := range fl.Players {
		fl.Players[j].Cards = []Card{}
	}

	for i := 0; i < 5; i++ {
		for j := range fl.Players {
			fl.Players[j].Cards = append(fl.Players[j].Cards, fl.Deck[cardNumber])
			fl.Players[j].Lives = 5
			fl.Deck[cardNumber].Used = true
			cardNumber++
			fl.Players[j].IsTurn = false
		}
	}

	fl.Players[0].Dealer = true
	fl.Players[1].IsTurn = true

	for _, listener := range fl.listeners {
		listener <- struct{}{}
	}
	return nil
}

// Turn ...
func (fl *FiveLives) Turn(card Card) {
	playerLostLife := false

	for i, player := range fl.Players {
		if player.IsTurn && player.Lives > 0 {
			player.LastCardPlayed = &card
			if len(fl.PreviousCard.Description) > 0 {
				if fl.PreviousCard.Description[0] == card.Description[0] {
					playerLostLife = true
					for j, lastPlayer := range fl.Players {
						if fl.LastPlayerName == lastPlayer.Name {
							lastPlayer.Lives = lastPlayer.Lives - fl.LivesToLose
							if (lastPlayer.Lives) < 0 {
								lastPlayer.Lives = 0
							}
							fl.Players[j] = lastPlayer
						}
					}
					fl.LivesToLose++
				}
			}

			for j, playerCard := range player.Cards {
				if card.Description == playerCard.Description {
					player.Cards = remove(player.Cards, j)
					break
				}
			}

			player.IsTurn = false
			fl.Players[i] = player
			if !fl.GameFinished {
				fl.setNextPlayer(i)
				fl.LastPlayerName = player.Name
			}
			break
		}
	}

	if !playerLostLife {
		fl.LivesToLose = 1
	}
	fl.PreviousCard = card
	fl.checkGameOver()

	fl.checkRoundOver()
	for _, listener := range fl.listeners {
		listener <- struct{}{}
	}
}

func (fl *FiveLives) setNextPlayer(currentPlayerPosition int) {
	fl.Players[currentPlayerPosition].IsTurn = false

	if currentPlayerPosition >= len(fl.Players) {
		fl.Players[0].IsTurn = true
		return
	}

	if currentPlayerPosition == len(fl.Players)-1 {
		for i := range fl.Players {
			if fl.Players[i].Lives > 0 {
				fl.Players[i].IsTurn = true
				return
			}
		}
	}

	for i := currentPlayerPosition + 1; i < len(fl.Players); i++ {
		if fl.Players[i].Lives > 0 {
			fl.Players[i].IsTurn = true
			return
		}
	}

	for i := 0; i < len(fl.Players); i++ {
		if fl.Players[i].Lives > 0 {
			fl.Players[i].IsTurn = true
			return
		}
	}
}

func (fl *FiveLives) checkGameOver() {
	playersStillIn := 0

	for _, player := range fl.Players {
		if player.Lives > 0 {
			playersStillIn++
		}
	}

	if playersStillIn < 2 {
		fl.GameFinished = true
	}
}

func (fl *FiveLives) checkRoundOver() {
	for _, player := range fl.Players {
		if player.Lives > 0 && len(player.Cards) > 0 {
			return
		}
	}

	fl.LastPlayerName = ""
	fl.Deck = NewDeck()
	fl.Deck.Shuffle()

	cardNumber := 0
	fl.LivesToLose = 1

	for i := 0; i < 5; i++ {
		for j := range fl.Players {
			if fl.Players[j].Lives > 0 {
				fl.Players[j].Cards = append(fl.Players[j].Cards, fl.Deck[cardNumber])
				fl.Deck[cardNumber].Used = true
				cardNumber++
			}
		}
	}

	for i := range fl.Players {
		if fl.Players[i].Dealer {
			dealerPosition := fl.setNextDealer(i)
			fl.setFirstPlayer(dealerPosition)
			break
		}
	}
}

func (fl *FiveLives) setNextDealer(currentDealerPosition int) int {
	fl.Players[currentDealerPosition].Dealer = false

	if currentDealerPosition >= len(fl.Players) {
		fmt.Println("setting here")
		fmt.Println(len(fl.Players))
		fl.Players[0].Dealer = true
		return 0
	}

	if currentDealerPosition == len(fl.Players)-1 {
		for i := range fl.Players {
			if fl.Players[i].Lives > 0 && i != currentDealerPosition {
				fmt.Println("setting here (235)")
				fmt.Println(i)
				fl.Players[i].Dealer = true
				return i
			}
		}
	}

	for i := currentDealerPosition + 1; i < len(fl.Players); i++ {
		if fl.Players[i].Lives > 0 && i != currentDealerPosition {
			fmt.Println("245")
			fmt.Println(i)
			fl.Players[i].Dealer = true
			return i
		}
	}

	for i := 0; i < len(fl.Players); i++ {
		if fl.Players[i].Lives > 0 && i != currentDealerPosition {
			fmt.Println(254)
			fmt.Println(i)
			fl.Players[i].Dealer = true
			return i
		}
	}

	return 0
}

func (fl *FiveLives) setFirstPlayer(dealerPosition int) {
	for i, player := range fl.Players {
		player.IsTurn = false
		fl.Players[i] = player
	}

	if dealerPosition >= len(fl.Players) {
		for i, player := range fl.Players {
			if player.Lives > 0 {
				fl.Players[i].IsTurn = true
				return
			}
		}
	}

	for i := dealerPosition + 1; i < len(fl.Players); i++ {
		if fl.Players[i].Lives > 0 {
			fl.Players[i].IsTurn = true
			return
		}
	}

	for i, player := range fl.Players {
		if player.Lives > 0 {
			fl.Players[i].IsTurn = true
			return
		}
	}
}

func remove(slice []Card, s int) []Card {
	return append(slice[:s], slice[s+1:]...)
}
