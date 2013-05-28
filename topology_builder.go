package gomethius

type BoltDeclarer struct {
	id string
	bolt Bolt
	parallelism int
	groupings map[string]string
}

func NewBoltDeclarer(id string, bolt Bolt, parallelism int) (*BoltDeclarer) {
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
	spout Spout
	parallelism int
}

func NewSpoutDeclarer(id string, spout Spout, parallelism int) (*SpoutDeclarer) {
	out := new(SpoutDeclarer)
	out.id = id
	out.spout = spout
	out.parallelism = parallelism

	return out
}
type TopologyBuilder struct {
	spouts map[string]*SpoutDeclarer
	bolts map[string]*BoltDeclarer
}


func NewTopologyBuilder() *TopologyBuilder {
	out := new(TopologyBuilder)
	out.spouts = make(map[string]*SpoutDeclarer)
	out.bolts = make(map[string]*BoltDeclarer)
	return out
}

func (tb *TopologyBuilder) SetSpout(id string, spout Spout, parallelism int)  (*SpoutDeclarer) {
	sp := NewSpoutDeclarer(id, spout, parallelism)
	tb.spouts[id] = sp
	return sp
}

func (tb *TopologyBuilder) SetBolt(id string, bolt Bolt, parallelism int)  (*BoltDeclarer) {
	bd := NewBoltDeclarer(id, bolt, parallelism)
	tb.bolts[id] = bd
	return bd
}

func (tb *TopologyBuilder) CreateTopology() (*Topology) {
	topology := NewTopology()
	conf := make(map[string]string)

	for boltId, bd := range tb.bolts {
		var dests [] chan Values

		for i := 0; i < bd.parallelism; i++ {
			bb := NewBoltBridge(bd.bolt)
			dests = append(dests, bb.Src)
			topology.Bolts[boltId] = append(topology.Bolts[boltId], bb)
		}

		for sourceId,groupingType := range bd.groupings {
			switch groupingType {
			case "shuffle":
				grouping := NewShuffleGrouping(dests)
				topology.Groupings[sourceId] = append(topology.Groupings[sourceId], grouping)
			case "field":
			case "all":
			case "none":
			}
		}

	}

	for spoutId, sd := range tb.spouts {
		for i := 0 ; i < sd.parallelism; i++ {
			groupings := topology.Groupings[spoutId]
			spoutBridge := NewSpoutBridge(sd.spout, groupings)
			spoutBridge.Open(conf)
			topology.Spouts[spoutId] = append(topology.Spouts[spoutId], spoutBridge)
		}
	}

	for boltId, bridges := range topology.Bolts {
		for _, bb := range bridges {
			bb.Groupings = topology.Groupings[boltId]
			bb.Prepare(conf)
		}
	}

	return topology
}

