package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ynqa/gomcts"
	"github.com/ynqa/gomcts/example/nim"
)

var (
	chips   int
	maxPick int
)

func main() {
	flag.IntVar(&chips, "chips", 10, "chips for nim")
	flag.IntVar(&maxPick, "maxPick", 3, "max number of chips to pick")
	flag.Parse()

	state := nim.NewNimState(chips, maxPick)

	for len(state.GetMoves()) > 0 {
		var move gomcts.Move
		if state.GetPlayerJustMoved() == 1 {
			// p2 play out
			move, _ = gomcts.UCT(state, 100)
		} else if state.GetPlayerJustMoved() == 2 {
			// p1 play out
			move, _ = gomcts.UCT(state, 10000)
		} else {
			log.Fatal("Who are you?")
		}
		state.DoMove(move)
		fmt.Println(state.String())
	}
	result, _ := state.GetResult(1)
	if result == 1. {
		fmt.Printf("Player#1 wins!")
	} else if result == 0. {
		fmt.Printf("Player#2 wins!")
	} else {
		log.Fatal("Who wins?")
	}
}
