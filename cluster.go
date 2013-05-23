package gomethius


type LocalCluster struct {

}

func NewLocalCluster() (*LocalCluster) {
	out := new(LocalCluster)
	return out
}

func (lc *LocalCluster)SubmitTopology(topologyId string, conf map[string]string, topology *Topology) {

}


type Cluster interface {

	SubmitTopology(topologyId string, conf map[string]string, topology Topology)

}
