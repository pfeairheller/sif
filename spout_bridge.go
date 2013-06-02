package gomethius


type SpoutBridge struct {
	Src chan Values
	Groupings [] Grouping
	Spout Spout
}

func NewSpoutBridge(s Spout, groupings []Grouping) *SpoutBridge {
	out := new(SpoutBridge)
	out.Src = make(chan Values, 10)
	out.Spout= s
	out.Groupings = groupings
	return out
}

func (sb *SpoutBridge) Open(conf map[string]string, context *TopologyContext) {
	sb.Spout.Open(conf, context, sb.Src)
}

func (sb *SpoutBridge) Launch() {
	go sb.processTuples()
	go sb.Run()
}

func (sb *SpoutBridge) Run() {
	for {
		sb.Spout.NextTuple()
	}
}

func (sb *SpoutBridge) processTuples() {
	for {
		tuple := <- sb.Src
		for _, grouping := range sb.Groupings {
			grouping.Tuple(tuple)
		}
	}
}
