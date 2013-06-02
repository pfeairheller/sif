package gomethius


type TopologyContext struct {
	Spouts map[string]*SpoutDeclarer
	Bolts map[string]*BoltDeclarer
}

func NewTopologyContext() *TopologyContext {
	out := new(TopologyContext)
	out.Spouts = make(map[string]*SpoutDeclarer)
	out.Bolts = make(map[string]*BoltDeclarer)
	return out
}

func (tc *TopologyContext) getOutputFields(sourceId string) (*Fields, bool) {
	spoutDec, exists := tc.Spouts[sourceId]
	if exists {
		return spoutDec.Spout.DeclareOutputFields(), true
	}

	boltDec, exists := tc.Bolts[sourceId]
	if exists {
		return boltDec.Bolt.DeclareOutputFields(), true
	}

	return nil, false
}

