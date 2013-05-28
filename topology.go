package gomethius


type Bolt interface {
	Prepare(conf map[string]string, emitter chan Values)
	Execute(tuple Values)
	DeclareOutputFields() [] string
}

type Spout interface {
	Open(conf map[string]string, emitter chan Values)
	NextTuple()
	DeclareOutputFields() [] string
}

type Grouping interface {
	Run()
	Launch()
	Tuple(tuple Values)
}

type Topology struct {
	Bolts map[string][]*BoltBridge
	Spouts map[string][]*SpoutBridge
	Groupings map[string][]Grouping
}

func NewTopology() (*Topology) {
	topology := new(Topology)
	topology.Bolts     = make(map[string][]*BoltBridge)
	topology.Spouts    = make(map[string][]*SpoutBridge)
	topology.Groupings = make(map[string][]Grouping)
	return topology
}



