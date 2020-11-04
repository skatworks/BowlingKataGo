package main

type Game struct {
     score int
}

func (game *Game) Roll(pins int) {
     game.score += pins
}

func (game *Game) GetScore() int {
     return game.score
}