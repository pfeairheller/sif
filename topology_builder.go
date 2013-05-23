package gomethius

type BoltDeclarer struct {
	id string
	bolt *Bolt
	parallelism uint32
	groupings map[string]string
}

func NewBoltDeclarer(id string, bolt *Bolt, parallelism uint32) (*BoltDeclarer) {
	out := new(BoltDeclarer)
	out.id = id
	out.bolt = bolt
	out.parallelism = parallelism
	out.groupings = make(map[string]string)

	return out
}

func (bd *BoltDeclarer) ShuffleGrouping(sourceId string) (*BoltDeclarer) {
	bd.groupings[sourceId] = "shuffle"
	return bd
}

type SpoutDeclarer struct {
	id string
	spout *Spout
	parallelism uint32
}

func NewSpoutDeclarer(id string, spout *Spout, parallelism uint32) (*SpoutDeclarer) {
	out := new(SpoutDeclarer)
	out.id = id
	out.spout = spout
	out.parallelism = parallelism

	return out
}
type TopologyBuilder struct {
	spouts map[string]SpoutDeclarer
	bolts map[string]BoltDeclarer
}


func NewTopologyBuilder() *TopologyBuilder {
	out := new(TopologyBuilder)
	out.spouts = make(map[string]SpoutDeclarer)
	out.bolts = make(map[string]BoltDeclarer)
	return out
}

func (tb *TopologyBuilder) SetSpout(id string, spout Spout, parallelism uint32)  (*SpoutDeclarer) {
	sp := NewSpoutDeclarer(id, &spout, parallelism)
	return sp
}

func (tb *TopologyBuilder) SetBolt(id string, bolt Bolt, parallelism uint32)  (*BoltDeclarer) {
	bd := NewBoltDeclarer(id, &bolt, parallelism)
	return bd
}

func (tb *TopologyBuilder) CreateTopology() (*Topology) {
	topology := NewTopology()

	//For Spouts
//	dests := make([] chan Values, 1)

	//For Bolts
//	bb := NewBoltBridge(bolt)
//	dests[0] = bb.Src

	//For Grouping tie-ins
//	grouping := NewShuffleGrouping(dests)


	return topology
}

