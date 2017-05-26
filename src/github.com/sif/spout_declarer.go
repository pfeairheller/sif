package sif

type SpoutDeclarer struct {
	id string
	Spout Spout
	parallelism, numTasks int
}

func NewSpoutDeclarer(id string, spout Spout, parallelism int) (*SpoutDeclarer) {
	out := new(SpoutDeclarer)
	out.id = id
	out.Spout = spout
	out.parallelism = parallelism

	return out
}

func (sd *SpoutDeclarer) SetNumTasks(num int) (*SpoutDeclarer) {
	sd.numTasks = num
	return sd
}

