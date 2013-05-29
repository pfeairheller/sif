

package gomethius

import (
	"math/rand"
)


type FieldGrouping struct {
	Src chan Values
	Dests [] chan Values
	Selector *Fields
}


func NewFieldGrouping(selector *Fields) *FieldGrouping {
	out := new(FieldGrouping)
	out.Src = make(chan Values, 10)
	out.Selector = selector
	return out
}

func(g *FieldGrouping) Prepare(conf map[string]string, dests []chan Values) {
	g.Dests = dests
}

func(g*FieldGrouping) Launch() {
	go g.Run()
}

func (g *FieldGrouping) Tuple(tuple Values) {
	g.Src <- tuple
}

func (g *FieldGrouping) Run() {
	for {
		tuple := <- g.Src
		
		idx := rand.Int31n(int32(len(g.Dests)))
		g.Dests[idx] <- tuple
	}
}

type Fields struct {
	fields []string
	fieldMap map[string]int
}

func NewFields(fields ...string) (*Fields) {
	out := new(Fields)
	out.fields = fields
	out.fieldMap = make(map[string]int)
	for idx, field :=  range fields {
		out.fieldMap[field] = idx
	}
	return out
}

func(f *Fields) fieldIndex(field string) int {
	return f.fieldMap[field]
}

func(f *Fields) get(idx int) string {
	return f.fields[idx]
}

func(f *Fields) size() int {
	return len(f.fields)
}

func(f *Fields) selects(selector *Fields, tuple Values) [] Value {
	var out []Value
	for _, value := range selector.fields {
		idx := f.fieldIndex(value)
		out = append(out, tuple.Get(idx))
	}
	return out
}

