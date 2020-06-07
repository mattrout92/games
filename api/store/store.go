package store

import "github.com/mattrout92/games/api/games"

// Store ...
type Store interface {
	Create(string) string
	Get(id string) games.Game
	Save(string, games.Game)
}
