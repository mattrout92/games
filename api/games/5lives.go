package games

import "errors"

var _ Game = (*FiveLives)(nil)

// FiveLives ...
type FiveLives struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Deck         *Deck    `json:"-"`
	Started      bool     `json:"started"`
	Players      []Player `json:"players"`
	PreviousCard Card     `json:"previous_card"`
	LivesToLose  int      `json:"lives_to_lose"`
	GameFinished bool     `json:"game_finished"`
	messages     chan (struct{})
}

// Messages ...
func (fl *FiveLives) Messages() chan (struct{}) {
	if fl.messages == nil {
		fl.messages = make(chan struct{})
	}
	return fl.messages
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

	fl.messages <- struct{}{}

	return nil
}

// GetPlayers ...
func (fl *FiveLives) GetPlayers() []Player {
	return fl.Players
}

// Start ...
func (fl *FiveLives) Start() error {
	if len(fl.Players) < 2 {
		return errors.New("not enough players")
	}

	fl.Deck = NewDeck()
	fl.Deck.Shuffle()

	cardNumber := 0
	fl.LivesToLose = 1

	for i := 0; i < 5; i++ {
		for j := range fl.Players {
			fl.Players[j].Cards = append(fl.Players[j].Cards, fl.Deck[cardNumber])
			fl.Deck[cardNumber].Used = true
			cardNumber++
		}
	}

	fl.Players[0].Dealer = true
	fl.Players[1].IsTurn = true

	fl.messages <- struct{}{}
	return nil
}

// Turn ...
func (fl *FiveLives) Turn(card Card) {
	playerLostLife := false

	for i, player := range fl.Players {
		if player.IsTurn && player.Lives > 0 {
			if fl.PreviousCard.Value == card.Value {
				playerLostLife = true
				player.Lives = player.Lives - fl.LivesToLose
				if (player.Lives) < 0 {
					player.Lives = 0
				}
				fl.LivesToLose++
			}

			for j, playerCard := range player.Cards {
				if card.Suit == playerCard.Suit && card.Value == playerCard.Value {
					player.Cards = remove(player.Cards, j)
					break
				}
			}

			fl.checkGameOver()

			if !fl.GameFinished {
				fl.setNextPlayer(i)
			}

			fl.checkRoundOver()
			fl.Players[i] = player
		}
	}

	if !playerLostLife {
		fl.LivesToLose = 1
	}
	fl.messages <- struct{}{}
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
		if len(player.Cards) > 0 {
			return
		}
	}

	fl.Deck = NewDeck()
	fl.Deck.Shuffle()

	cardNumber := 0
	fl.LivesToLose = 1

	for i := 0; i < 5; i++ {
		for j := range fl.Players {
			fl.Players[j].Cards = append(fl.Players[j].Cards, fl.Deck[cardNumber])
			fl.Deck[cardNumber].Used = true
			cardNumber++
		}
	}

	for i := range fl.Players {
		if fl.Players[i].Dealer {
			dealerPosition := fl.setNextDealer(i)
			fl.setFirstPlayer(dealerPosition)
		}
	}
}

func (fl *FiveLives) setNextDealer(currentDealerPosition int) int {
	fl.Players[currentDealerPosition].Dealer = false

	if currentDealerPosition >= len(fl.Players) {
		fl.Players[0].Dealer = true
		return 0
	}

	if currentDealerPosition == len(fl.Players)-1 {
		for i := range fl.Players {
			if fl.Players[i].Lives > 0 {
				fl.Players[i].Dealer = true
				return i
			}
		}
	}

	for i := currentDealerPosition + 1; i < len(fl.Players); i++ {
		if fl.Players[i].Lives > 0 {
			fl.Players[i].Dealer = true
			return i
		}
	}

	for i := 0; i < len(fl.Players); i++ {
		if fl.Players[i].Lives > 0 {
			fl.Players[i].Dealer = true
			return i
		}
	}

	return 0
}

func (fl *FiveLives) setFirstPlayer(dealerPosition int) {
	for _, player := range fl.Players {
		player.IsTurn = false
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
