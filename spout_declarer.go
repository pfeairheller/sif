package gomethius

type SpoutDeclarer struct {
	id string
	Spout Spout
	parallelism int
}

func NewSpoutDeclarer(id string, spout Spout, parallelism int) (*SpoutDeclarer) {
	out := new(SpoutDeclarer)
	out.id = id
	out.Spout = spout
	out.parallelism = parallelism

	return out
}


