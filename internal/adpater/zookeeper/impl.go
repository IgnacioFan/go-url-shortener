package zookeeper

import (
	"fmt"
	"strconv"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

const (
  ROOT_PATH = "/url"
  ZK_DOMAIN = "localhost" // os.Getenv("ZOOKER_URL")
	TOKEN_RANGE = 100
)

type Zookeeper interface {
	SetNewRange() (int, int, error)
}

type Impl struct {
	Client *zk.Conn
}

func InitZooKeeper() (*Impl, error) {
	conn, _, err := zk.Connect([]string{ZK_DOMAIN}, time.Second)
	if err != nil {
		return nil, err
	}

	ok, _, err := conn.Exists(ROOT_PATH)
	if err != nil {
		return nil, err
	}
	fmt.Println("Does root path existed?", ok)
	if !ok {
		res, err := conn.Create(ROOT_PATH, []byte("0"), 0, zk.WorldACL(zk.PermAll))
		fmt.Println("create a root path", res)
		if err != nil {
			return nil, err
		}
	} 
	return &Impl{conn}, nil
}

func (i *Impl) SetNewRange() (int, int, error) {
	var start, end int
	val, _, _ := i.Client.Get(ROOT_PATH)
	start, _ = strconv.Atoi(string(val))
	end = start + TOKEN_RANGE

	queue := make([]string, 0)
	queue = append(queue, rangeStr(start, end))
	for len(queue) > 0 {
		currRange := queue[0]
		queue = queue[1:]

		str, err := i.Client.Create(ROOT_PATH + "/" + currRange, []byte(currRange), 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			start = end
			end = start + TOKEN_RANGE
			queue = append(queue, rangeStr(start, end))
		} else {
			fmt.Println(str)
			break
		}
	}
	data := fmt.Sprintf("%v", end)
	if _, err := i.Client.Set(ROOT_PATH, []byte(data), -1); err != nil {
		return -1, -1, err
	}
	return start + 1, end, nil
}

func rangeStr(start, end int) string {
	return fmt.Sprintf("%v-%v", start + 1, end)
}



