package gomethius


type Bolt interface {
	Execute(tuple Values)
	DeclareOutputFields() [] string
}

type Spout interface {
	NextTuple(emitter chan Values)
	DeclareOutputFields() [] string
}

type Grouping interface {
	Run()
}
