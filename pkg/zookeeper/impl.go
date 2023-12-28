package zookeeper

import (
	"fmt"
	"strconv"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

const (
	TOKEN_PATH = "/token"
)

var (
	URL = "localhost" // os.Getenv("ZOOKER_URL")
	zkTokenRange = tokenRange{
		start: 0,
		end: 0,
		curr: 0,
	}
)

type tokenRange struct {
	start, end, curr int
}

type Impl struct {
	Client *zk.Conn
}

func InitZooKeeper() (*Impl, error) {
	conn, _, err := zk.Connect([]string{URL}, time.Second)
	if err != nil {
		return nil, err
	}

	ok, _, err := conn.Exists(TOKEN_PATH)
	if err != nil {
		return nil, err
	}
	if !ok {
		// when a node starts, register itself in Zookeeper
		data := []byte("0")
		res, err := conn.Create(TOKEN_PATH, data, 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			return nil, err
		}
		fmt.Println("token path registered:", res)
	}
	return &Impl{conn}, nil
}

func (i *Impl) GetCounter() int {
	counter := zkTokenRange.curr
	// increase
	if zkTokenRange.curr < zkTokenRange.end - 1 && zkTokenRange.curr != 0 {
		zkTokenRange.curr++
	} else {
		i.NextRange()
		zkTokenRange.curr++
	}
	return counter
}

func (i *Impl) NextRange() {
	data, _, error := i.Client.Get(TOKEN_PATH)
	if error != nil {
		panic(error)
	}
	fmt.Println(zkTokenRange)
	val, _ := strconv.Atoi(string(data))
	zkTokenRange.start = val + 1000000
	zkTokenRange.curr = val + 1000000
	zkTokenRange.end = val + 2000000

	i.setRange(zkTokenRange.start)
}

func (i *Impl) setRange(start int) {
	data := []byte(strconv.Itoa(start))
	if _, err := i.Client.Set(TOKEN_PATH, data, -1); err != nil {
		panic(err)
	}
} 
