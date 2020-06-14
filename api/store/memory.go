package store

import (
	"math/rand"
	"strings"
	"time"

	"github.com/mattrout92/games/api/games"
)

var _ Store = (*Memory)(nil)

// Memory ...
type Memory struct {
	games map[string]games.Game
}

// Create ...
func (m *Memory) Create(gameName string) string {
	id := newID()
	if m.games == nil {
		m.games = make(map[string]games.Game)
	}

	if gameName == "5Lives" {
		m.games[id] = &games.FiveLives{
			ID:   id,
			Name: gameName,
		}
	}

	return id
}

// Get ...
func (m *Memory) Get(id string) games.Game {
	game, ok := m.games[id]
	if ok {
		return game
	}

	return nil
}

// Save ...
func (m *Memory) Save(id string, game games.Game) {
	if m.games == nil {
		m.games = make(map[string]games.Game)
	}
	m.games[id] = game
}

func newID() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune(
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	length := 4
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return strings.ToLower(str)
}
