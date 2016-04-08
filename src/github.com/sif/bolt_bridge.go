package sif


type BoltBridge struct {
	Src chan Values
	Dest chan Values
	Groupings [] Grouping
	bolt Bolt
}

func NewBoltBridge(b Bolt) *BoltBridge {
	out := new(BoltBridge)
	out.Src = make(chan Values, 10)
	out.Dest = make(chan Values, 10)
	out.bolt = b
	return out
}

func (bb *BoltBridge) Prepare(conf map[string]string, context *TopologyContext) {
	bb.bolt.Prepare(conf, context, bb.Dest)
}


func (bb *BoltBridge) Launch() {
	go bb.processTuples()
	go bb.Run()
}

func (bb *BoltBridge) Run() {
	for {
		tuple := <- bb.Src
		bb.bolt.Execute(tuple)
	}
}

func (bb *BoltBridge) processTuples() {
	for {
		tuple := <- bb.Dest
		for _, grouping := range bb.Groupings {
			grouping.Tuple(tuple)
		}
	}
}
