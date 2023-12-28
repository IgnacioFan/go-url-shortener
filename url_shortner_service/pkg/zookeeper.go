package pkg

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

const (
	zkIdRangePath = "/id_range"
)

type ZooKeeperClient struct {
	Client *zk.Conn
}

func NewZookeeperClient() *ZooKeeperClient {
	servers := []string{"localhost:2181"}

	conn, _, err := zk.Connect(servers, time.Second)
	if err != nil {
		fmt.Println("zookeeper connect errors", err)
	}
	defer conn.Close()

	exist, _, err := conn.Exists(zkIdRangePath)
	if err != nil {
		fmt.Println("parent node err:", err)
	}
	if !exist {
		_, err = conn.Create(zkIdRangePath, []byte("ID range information"), 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			fmt.Println("can't create node, err:", err)
		}
	}

	return &ZooKeeperClient{
		Client: conn,
	}
}

func (z *ZooKeeperClient) GetTokenRange(id int) {
	path := fmt.Sprintf("%s/node_%d", zkIdRangePath, id)
	createdPath, err := z.Client.Create(path, nil, zk.FlagSequence, zk.WorldACL(zk.PermAll))
	if err != nil {
		fmt.Printf("Node %d failed to create a sequential znode: %v \n", id, err)
		return
	}

	// Extract the sequential ID from the created znode path
	sequentialID := extractSequentialID(createdPath)
	fmt.Println("Node", id, " got ID range:", sequentialID)
}

func extractSequentialID(path string) string {
	// Extract the sequential ID from the end of the path
	lastIndex := len(path) - 1
	for i := lastIndex; i >= 0; i-- {
		if path[i] == '-' {
			return path[i+1 : lastIndex]
		}
	}

	// Return the original path if no sequential ID is found (shouldn't happen)
	return path
}
