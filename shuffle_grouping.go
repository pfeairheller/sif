

package gomethius

import (
	"math/rand"
)


type ShuffleGrouping struct {
	Src chan Values
	Dests [] chan Values
}


func NewShuffleGrouping(dest [] chan Values) *ShuffleGrouping {
	out := new(ShuffleGrouping)
	out.Src = make(chan Values, 10)
	out.Dests = dest
	return out
}

func(g*ShuffleGrouping) Launch() {
	go g.Run()
}

func (g *ShuffleGrouping) Tuple(tuple Values) {
	g.Src <- tuple
}

func (g *ShuffleGrouping) Run() {
	for {
		tuple := <- g.Src
		idx := rand.Int31n(int32(len(g.Dests)))
		g.Dests[idx] <- tuple
	}
}
