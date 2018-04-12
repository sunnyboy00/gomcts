package main

import (
	"fmt"
	"log"
	"flag"

	"github.com/ynqa/gomcts"
	"github.com/ynqa/gomcts/example/nim"
)

var (
	chips int
	maxPick int
)

func main() {
	flag.IntVar(&chips, "chips", 15, "chips for nim")
	flag.IntVar(&maxPick, "maxPick", 3, "max number of chips to pick")
	flag.Parse()

	state := nim.NewNimState(chips, maxPick)
	for len(state.GetMoves()) > 0 {
		move, err := gomcts.UCT(1, state, 10)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(state.String())
		state.DoMove(move)
	}
}