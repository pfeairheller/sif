package sif

type BoltDeclarer struct {
	id string
	Bolt Bolt
	parallelism int
	groupings map[string]Grouping
}

func NewBoltDeclarer(id string, bolt Bolt, parallelism int) (*BoltDeclarer) {
	out := new(BoltDeclarer)
	out.id = id
	out.Bolt = bolt
	out.parallelism = parallelism
	out.groupings = make(map[string]Grouping)

	return out
}

func (bd *BoltDeclarer) ShuffleGrouping(sourceId string) (*BoltDeclarer) {
	bd.groupings[sourceId] = NewShuffleGrouping()
	return bd
}

func (bd *BoltDeclarer) FieldGrouping(sourceId string, fields *Fields) (*BoltDeclarer) {
	bd.groupings[sourceId] = NewFieldGrouping(sourceId, fields)
	return bd
}

