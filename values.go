
package gomethius

type Values struct {
	values []string
}

func NewValues(values []string) *Values {
	out := new(Values)
	out.values = values
	return out
}

