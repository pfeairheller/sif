package sif


type Cluster interface {

	SubmitTopology(topologyId string, conf *Configuration, topology Topology)

}
