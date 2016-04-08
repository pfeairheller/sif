package sif

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
