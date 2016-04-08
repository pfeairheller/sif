package sif


type Bolt interface {
	Prepare(conf map[string]string, context *TopologyContext, emitter chan Values)
	Execute(tuple Values)
	DeclareOutputFields() *Fields
}

type Spout interface {
	Open(conf map[string]string, context *TopologyContext, emitter chan Values)
	NextTuple()
	DeclareOutputFields() *Fields
}

type Grouping interface {
	Prepare(conf map[string]string, context *TopologyContext, dest []chan Values)
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



