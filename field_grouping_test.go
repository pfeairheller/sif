package gomethius

import (
	"testing"
	"github.com/bmizerany/assert"
)


func TestModHashingOfSingleStringValue(t *testing.T) {
	fields1 := NewFields("Symbol", "Bid", "Ask")
	fields2 := NewFields("Symbol")
	values := NewValues(NewStringValue("USD/EUR"), NewFloat32Value(45), NewFloat32Value(46))


	fg := NewFieldGrouping("test", fields2)
	fg.SourceFields = fields1
	fg.Dests = make([] chan Values, 3)

	assert.Equal(t, fg.ModHash(*values), uint32(1))
}
