package distributed

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/client"
	"github.com/sif"
	"golang.org/x/net/context"
)

type DistributedCluster struct {
	Monitor                 chan string
	ID, ClusterName, Leader string
	EtcdClient              client.Client
	LeaderIndex, ClusterIndex uint64
}

func NewDistributedCluster(clusterName string) *DistributedCluster {
	out := new(DistributedCluster)
	uuid, _ := NewV4()
	out.ID = uuid.String()
	out.ClusterName = clusterName
	out.Monitor = make(chan string)
	out.joinCluster()
	return out
}

func (dc *DistributedCluster) SubmitTopology(topologyId string, conf *sif.Configuration, topology *sif.Topology) {

}

func (dc *DistributedCluster) joinCluster() {

	cfg := client.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
		Transport: client.DefaultTransport,
	}

	var err error

	dc.EtcdClient, err = client.New(cfg)
	if err != nil {
		log.Fatalln("Bad config", err)
	}

	kapi := client.NewKeysAPI(dc.EtcdClient)
	get, err := kapi.Get(context.Background(), dc.leaderKey(), &client.GetOptions{})

	if err != nil || get.Node.Value == "" {
		go dc.leadCluster("")
	} else {
		dc.Leader = get.Node.Value
		dc.LeaderIndex = get.Node.ModifiedIndex
		dc.ClusterIndex = get.Index
		go dc.watchCluster()
	}

	ch := make(chan bool)
	<-ch

}

func (dc *DistributedCluster) leadCluster(currentLeader string) {

	leaderSetOptions := &client.SetOptions{
		TTL: 5 * time.Second,
	}

	if currentLeader == "" {
		leaderSetOptions.PrevExist = client.PrevNoExist
	} else {
		leaderSetOptions.PrevValue = currentLeader
	}

	kapi := client.NewKeysAPI(dc.EtcdClient)
	resp, err := kapi.Set(context.Background(), dc.leaderKey(), dc.ID, leaderSetOptions)
	if err != nil {
		dc.LeaderIndex = 0
		dc.ClusterIndex = 0
		dc.Leader = ""
		go dc.watchCluster()
		return
	}

	dc.Leader = resp.Node.Value
	dc.LeaderIndex = resp.Node.ModifiedIndex
	dc.ClusterIndex = resp.Index

	fmt.Printf("I (%s) am leader\n", dc.ID)

	timer := time.Tick(4 * time.Second)

	leaderRefreshOptions := &client.SetOptions{
		TTL:     5 * time.Second,
		Refresh: true,
	}

	for {
		<-timer
		_, err := kapi.Set(context.Background(), dc.leaderKey(), "", leaderRefreshOptions)
		if err != nil {
			log.Fatalln("Can't continue to be leader", err)
		}

		fmt.Printf("I'm (%s) still leader\n", dc.ID)
	}

}

func (dc *DistributedCluster) watchCluster() {
	fmt.Println("Identified", dc.Leader, " as the leader")

	kapi := client.NewKeysAPI(dc.EtcdClient)
	watcher := kapi.Watcher(dc.leaderKey(), &client.WatcherOptions{})

	var err error
	err = nil
	var resp *client.Response
	for err == nil {
		resp, err = watcher.Next(context.Background())
		if resp.Node.Value == "" {
			go dc.leadCluster("")
			return
		} else {
			dc.Leader = resp.Node.Value
			fmt.Println(dc.Leader, "is the new leader")
		}
	}

	if err != nil {
		log.Fatalln("Error watching", err)
	}

}

func (dc *DistributedCluster) leaderKey() string {
	return fmt.Sprintf("/sif/clusters/%s", dc.ClusterName)
}

func (dc *DistributedCluster) getCurrentLeader() error {
	kapi := client.NewKeysAPI(dc.EtcdClient)
	get, err := kapi.Get(context.Background(), dc.leaderKey(), &client.GetOptions{})

	if err != nil || get.Node.Value == "" {
		go dc.leadCluster("")
	} else {
		dc.Leader = get.Node.Value
		dc.LeaderIndex = get.Node.ModifiedIndex
		dc.ClusterIndex = get.Index
		go dc.watchCluster()
	}

}
