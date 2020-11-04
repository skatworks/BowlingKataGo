package main

type Game struct {
     rolls [21]int
     currentIndex int
}

func (game *Game) Roll(pins int) {
     game.rolls[game.currentIndex] = pins
     game.currentIndex++
}

func (game *Game) GetScore() int {
     score := 0
     frameIndex := 0

     for frame := 0; frame < 10; frame++ {
         if game.rolls[frameIndex] == 10 {
            score += 10 + game.rolls[frameIndex + 1] + game.rolls[frameIndex + 2]
            frameIndex += 1
         } else if game.rolls[frameIndex] + game.rolls[frameIndex + 1] == 10 {
            score += 10 + game.rolls[frameIndex + 2]
            frameIndex += 2
         } else {
            score += game.rolls[frameIndex] + game.rolls[frameIndex + 1]
            frameIndex += 2
         }
     }
     return score
}