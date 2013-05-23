
package gomethius

type Value interface {}

type Values struct {
	values []Value
}

func NewValues(values ...Value) *Values {
	out := new(Values)
	out.values = values
	return out
}

func (v *Values) Get(idx int) Value {
	return v.values[idx]
}

