package games

// Game ...
type Game interface {
	AddPlayer(name string) error
	GetPlayers() []Player
	Start() error
	Turn(Card)
	Messages() chan (struct{})
}

// Player ...
type Player struct {
	Name   string `json:"name"`
	IsTurn bool   `json:"is_turn"`
	Lives  int    `json:"lives"`
	Cards  []Card `json:"cards"`
	Dealer bool   `json:"dealer"`
}
