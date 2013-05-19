package gomethius


type BoltBridge struct {
	Src chan Values
	bolt Bolt
}

func NewBoltBridge(b Bolt) *BoltBridge {
	out := new(BoltBridge)
	out.Src = make(chan Values, 10)
	out.bolt = b
	return out
}

func (bb *BoltBridge) Run() {
	for {
		tuple := <- bb.Src
		bb.bolt.Execute(tuple)
	}
}
