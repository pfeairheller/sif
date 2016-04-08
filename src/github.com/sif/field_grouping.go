package sif

import (
	"hash/fnv"
)

type FieldGrouping struct {
	Src      chan Values
	Dests    []chan Values
	SourceId string
	SourceFields *Fields
	Selector *Fields
}

func NewFieldGrouping(sourceId string, selector *Fields) *FieldGrouping {
	out := new(FieldGrouping)
	out.Src = make(chan Values, 10)
	out.SourceId = sourceId
	out.Selector = selector
	return out
}

func (g *FieldGrouping) Prepare(conf map[string]string, context *TopologyContext,  dests []chan Values) {
	g.Dests = dests
	sourceFields, exists := context.getOutputFields(g.SourceId)
	if exists {
		g.SourceFields = sourceFields
	}
}

func (g *FieldGrouping) Launch() {
	go g.Run()
}

func (g *FieldGrouping) Tuple(tuple Values) {
	g.Src <- tuple
}

func (g *FieldGrouping) Run() {
	for {
		tuple := <-g.Src
		idx := g.ModHash(tuple)
		g.Dests[idx] <- tuple
	}
}

func (g *FieldGrouping) ModHash(tuple Values) uint32 {
	h := fnv.New32()
	values := g.SourceFields.Select(g.Selector, tuple)
	for _, value := range values {
		h.Write(value.Encode())
	}
	h = fnv.New32()
	return h.Sum32() % uint32(len(g.Dests))
}



