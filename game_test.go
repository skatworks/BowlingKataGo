package main

import "testing"

func TestGutterGame(t *testing.T) {
     game := Game{}
     RollMany(20, 0, &game)
     ResultCheck(game.GetScore(), 0, t)
}

func Test1Pins(t *testing.T) {
     game := Game{}
     RollMany(20, 1, &game)
     ResultCheck(game.GetScore(), 20, t)
}

func TestSpare(t *testing.T) {
     game := Game{}
     game.Roll(5)
     game.Roll(5)
     game.Roll(3)
     game.Roll(4)
     RollMany(16, 0, &game)
     ResultCheck(game.GetScore(), 20, t)
}

func RollMany(n int, pins int, game *Game) {
     for i := 0; i < n; i++ {
         game.Roll(pins)
     }
}

func ResultCheck(actual int, expected int, t *testing.T) {
     if actual != expected {
        t.Fail()
        t.Logf("Expected %d, got %d", expected, actual)
     }
}