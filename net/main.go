package main

import (
	"fmt"
	"github.com/subgraph/go-nfnetlink/nfqueue"
	"os"
)

func main() {
	q := nfqueue.NewNFQueue(1)

	ps, err := q.Open()
	if err != nil {
		fmt.Printf("Error opening NFQueue: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("open queue success...")
	}
	defer q.Close()

	for p := range ps {
		networkLayer := p.Packet.NetworkLayer()
		fmt.Println("networkLayer=", networkLayer)
		ipsrc, ipdst := networkLayer.NetworkFlow().Endpoints()

		transportLayer := p.Packet.TransportLayer()
		fmt.Println("transportLayer=", transportLayer)
		tcpsrc, tcpdst := transportLayer.TransportFlow().Endpoints()

		fmt.Printf("A new tcp connection will be established: %s:%s -> %s:%s\n",
			ipsrc, tcpsrc, ipdst, tcpdst)
		p.Drop()
	}
}
