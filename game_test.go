package main

import "testing"

func TestGutterGame(t *testing.T) {
     game := Game{}
     for i := 0; i < 20; i++ {
         game.Roll(0)
     }
     if game.GetScore() != 0 {
        t.Fail()
        t.Logf("Expected %d, got %d", 0, game.GetScore())
     }
}
