package gomethius

import (
	// "fmt"
	"hash/fnv"
)

type FieldGrouping struct {
	Src      chan Values
	Dests    []chan Values
	SourceFields *Fields
	Selector *Fields
}

func NewFieldGrouping(selector *Fields) *FieldGrouping {
	out := new(FieldGrouping)
	out.Src = make(chan Values, 10)
	out.Selector = selector
	return out
}

func (g *FieldGrouping) Prepare(conf map[string]string, dests []chan Values) {
	g.Dests = dests
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
		h.Write([]byte(value.([]byte)))
	}
	h = fnv.New32()
	return h.Sum32() % uint32(len(g.Dests))
}

type Fields struct {
	fields   []string
	fieldMap map[string]int
}

func NewFields(fields ...string) *Fields {
	out := new(Fields)
	out.fields = fields
	out.fieldMap = make(map[string]int)
	for idx, field := range fields {
		out.fieldMap[field] = idx
	}
	return out
}

func (f *Fields) FieldIndex(field string) int {
	return f.fieldMap[field]
}

func (f *Fields) Get(idx int) string {
	return f.fields[idx]
}

func (f *Fields) Size() int {
	return len(f.fields)
}

func (f *Fields) Select(selector *Fields, tuple Values) []Value {
	var out []Value
	for _, value := range selector.fields {
		idx := f.FieldIndex(value)
		out = append(out, tuple.Get(idx))
	}
	return out
}
