package main

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
	peerstore "github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	p2p.connectNode()
	node, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/46269"),
		libp2p.Ping(false),
	)
	if err != nil {
		panic(err)
	}

	// configure our own ping protocol
	pingService := &ping.PingService{Host: node}
	node.SetStreamHandler(ping.ID, pingService.PingHandler)

	// print the node's listening addresses
	fmt.Println("Listen addresses:", node.Addrs())

    	// print the node's PeerInfo in multiaddr format
	peerInfo := peerstore.AddrInfo{
		ID:    node.ID(),
		Addrs: node.Addrs(),
	}
	addrs, err := peerstore.AddrInfoToP2pAddrs(&peerInfo)
	fmt.Println("libp2p node address:", addrs[0])
	if err != nil {
	    fmt.Println("libp2p node address: Not Found!!!")
	    return	
	}

    go func ()  { 
        // shut the node down
        if err := node.Close(); err != nil {
            panic(err)
        }
    }()

	// Connect to Local Database
	db, err := leveldb.OpenFile("cahin-storage/db", nil)

	if err != nil {
		fmt.Println("LevelDB Not Found")
		return 
	} else {
	    fmt.Println("DB connected")
	}

	err = db.Put([]byte("userID"), []byte(node.ID()), nil)
	if err != nil {
		fmt.Println("Node.ID() not found")
		return  
	} 

	defer db.Close()	
}