package main

import "testing"

func TestGutterGame(t *testing.T) {
     game := Game{}
     RollMany(20, 0, &game)
     ResultCheck(0, game.GetScore(), t)
}

func Test1Pins(t *testing.T) {
     game := Game{}
     RollMany(20, 1, &game)
     ResultCheck(20, game.GetScore(), t)
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