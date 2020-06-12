package games

// Game ...
type Game interface {
	AddPlayer(name string) error
	GetPlayers() []Player
	Start() error
	Turn(Card)
	AddListener(chan (struct{}))
}

// Player ...
type Player struct {
	Name           string `json:"name"`
	IsTurn         bool   `json:"is_turn"`
	Lives          int    `json:"lives"`
	Cards          []Card `json:"cards"`
	LastCardPlayed *Card  `json:"last_card_played,omitempty"`
	Dealer         bool   `json:"dealer"`
}
