package Game

type GameState int8

const (
	Run GameState = iota
	PlayerDead
	Quit
)
