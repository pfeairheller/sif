package gomethius


type Bolt interface {
	Execute(tuple Values)
	DeclareOutputFields() [] string
}

type Spout interface {
	NextTuple(emitter chan Values)
	DeclareOutputFields() [] string
}

type Grouping interface {
	Run()
}

type Topology struct {
	Bolts [] Bolt
	Spouts [] Spout
}

func NewTopology() (*Topology) {
	topology := new(Topology)

	return topology
}

