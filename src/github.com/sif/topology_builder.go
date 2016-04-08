package sif

type TopologyBuilder struct {
	Context *TopologyContext
}


func NewTopologyBuilder() *TopologyBuilder {
	out := new(TopologyBuilder)
	out.Context = NewTopologyContext();
	return out
}

func (tb *TopologyBuilder) SetSpout(id string, spout Spout, parallelism int)  (*SpoutDeclarer) {
	sp := NewSpoutDeclarer(id, spout, parallelism)
	tb.Context.Spouts[id] = sp
	return sp
}

func (tb *TopologyBuilder) SetBolt(id string, bolt Bolt, parallelism int)  (*BoltDeclarer) {
	bd := NewBoltDeclarer(id, bolt, parallelism)
	tb.Context.Bolts[id] = bd
	return bd
}

func (tb *TopologyBuilder) CreateTopology() (*Topology) {
	topology := NewTopology()
	conf := make(map[string]string)

	for boltId, bd := range tb.Context.Bolts {
		var dests [] chan Values

		for i := 0; i < bd.parallelism; i++ {
			bb := NewBoltBridge(bd.Bolt)
			dests = append(dests, bb.Src)
			topology.Bolts[boltId] = append(topology.Bolts[boltId], bb)
		}

		for sourceId,grouping := range bd.groupings {
			grouping.Prepare(conf, tb.Context, dests)
			topology.Groupings[sourceId] = append(topology.Groupings[sourceId], grouping)
		}
		
	}

	for spoutId, sd := range tb.Context.Spouts {
		for i := 0 ; i < sd.parallelism; i++ {
			groupings := topology.Groupings[spoutId]
			spoutBridge := NewSpoutBridge(sd.Spout, groupings)
			spoutBridge.Open(conf, tb.Context)
			topology.Spouts[spoutId] = append(topology.Spouts[spoutId], spoutBridge)
		}
	}

	for boltId, bridges := range topology.Bolts {
		for _, bb := range bridges {
			bb.Groupings = topology.Groupings[boltId]
			bb.Prepare(conf, tb.Context)
		}
	}

	return topology
}

