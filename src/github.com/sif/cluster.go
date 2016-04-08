package sif

import "fmt"


type LocalCluster struct {
	Monitor chan string
}

func NewLocalCluster() (*LocalCluster) {
	out := new(LocalCluster)
	out.Monitor = make(chan string)
	return out
}

func (lc *LocalCluster)SubmitTopology(topologyId string, conf *Configuration, topology *Topology) {
	for _, bridges := range topology.Bolts {
		for _, boltBridge := range bridges {
			boltBridge.Launch()
		}
	}

	for _, groupSlices := range topology.Groupings {
		for _, grouping := range groupSlices{
			grouping.Launch()
		}
	}

	for _, bridges := range topology.Spouts {
		for _, spoutBridge := range bridges {
			spoutBridge.Launch()
		}
	}

	lc.monitor()
}

func (lc *LocalCluster) monitor() {
	for {
		message := <- lc.Monitor
		fmt.Println(message)
	}
}


type Cluster interface {

	SubmitTopology(topologyId string, conf *Configuration, topology Topology)

}
