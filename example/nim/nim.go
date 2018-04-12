package nim

import (
	"fmt"
	"math"

	"github.com/ynqa/gomcts"
)

type NimState struct {
	playerJustMoved int
	chips           int
	maxPick         int
}

func NewNimState(chips, maxPick int) *NimState {
	return &NimState{
		// There are 2 players, p1 moves first.
		playerJustMoved: 2,
		chips:           chips,
		maxPick:         maxPick,
	}
}

func (s *NimState) String() string {
	return fmt.Sprintf("PlayerJustMoved:%d, chips:%d",
		s.playerJustMoved, s.chips)
}

func (s *NimState) Clone() gomcts.State {
	return NewNimState(s.chips, s.maxPick)
}

func (s *NimState) GetPlayerJustMoved() int {
	return s.playerJustMoved
}

func (s *NimState) GetMoves() []gomcts.Move {
	moves := make([]gomcts.Move, 0)

	pickable := s.maxPick
	if s.chips < s.maxPick {
		pickable = s.chips
	}

	for i := 1; i <= pickable; i++ {
		moves = append(moves, newNimMove(i))
	}

	return moves
}

func (s *NimState) DoMove(move gomcts.Move) {
	nimMove := move.(*NimMove)
	s.chips -= nimMove.pickChips
	s.playerJustMoved = 3 - s.playerJustMoved
}

func (s *NimState) GetResult(id int) (float64, error) {
	if s.chips != 0 {
		return math.Inf(-1), fmt.Errorf("Chips size is not 0:%d", s.chips)
	}
	if id == s.playerJustMoved {
		return 1.0, nil
	}
	return 0.0, nil
}

type NimMove struct {
	pickChips int
}

func newNimMove(pickChips int) *NimMove {
	return &NimMove{
		pickChips: pickChips,
	}
}
