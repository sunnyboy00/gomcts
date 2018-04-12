package gomcts

import (
	"math"
	"math/rand"
	"sort"
)

type State interface {
	GetPlayerJustMoved() int
	Clone() State
	GetMoves() []Move
	DoMove(move Move)
	GetResult(id int) (float64, error)
}

type Move interface {
}

type Node struct {
	id          int
	parent      *Node
	state       State
	move        Move
	children    []*Node
	untriedMove []Move
	visits      float64
	wins        float64
}

func newNode(parent *Node, state State, move Move) *Node {
	return &Node{
		id:          state.GetPlayerJustMoved(),
		parent:      parent,
		state:       state,
		move:        move,
		untriedMove: state.GetMoves(),
	}
}

func (n *Node) UCB1() *Node {
	maxScore := math.Inf(-1)
	var maxChild *Node
	for _, child := range n.children {
		ucb := child.wins/child.visits +
			math.Sqrt(2.0*math.Log(n.visits)/child.visits)
		if ucb >= maxScore {
			maxScore = ucb
			maxChild = child
		}
	}
	return maxChild
}

func UCT(rootState State, iteration int) (Move, error) {
	rootNode := newNode(nil, rootState, nil)
	for i := 0; i < iteration; i++ {
		node := rootNode
		state := rootState.Clone()

		// Select
		for len(node.untriedMove) == 0 && len(node.children) > 0 {
			node = node.UCB1()
			state.DoMove(node.move)
		}

		// Expand
		if len(node.untriedMove) > 0 {
			n := rand.Intn(len(node.untriedMove))
			move := node.untriedMove[n]
			state.DoMove(move)

			// Delete untried
			node.untriedMove =
				append(node.untriedMove[:n],
					node.untriedMove[n+1:]...)

			// Append child
			child := newNode(node, state, move)
			node.children = append(node.children, child)
			node = child
		}

		// Rollout
		for len(state.GetMoves()) > 0 {
			availableMoves := state.GetMoves()
			n := rand.Intn(len(availableMoves))
			state.DoMove(availableMoves[n])
		}

		// Backpropagate
		for node != nil {
			node.visits++
			win, err := state.GetResult(node.id)
			if err != nil {
				return nil, err
			}
			node.wins += win
			node = node.parent
		}
	}

	// The best move
	sort.Sort(visitsDesc(rootNode.children))
	return rootNode.children[0].move, nil
}

// Descending sort by visits
type visitsDesc []*Node

func (n visitsDesc) Len() int           { return len(n) }
func (n visitsDesc) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n visitsDesc) Less(i, j int) bool { return n[i].visits > n[j].visits }
