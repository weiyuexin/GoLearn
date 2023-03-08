package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
)

var (
	pcapFile string = "file.cap"
	handle   *pcap.Handle
	err      error
)

func main() {
	// Open file instead of device
	handle, err = pcap.OpenOffline(pcapFile)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Loop through packets in file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		//fmt.Println(packet.Data())
		//fmt.Println("-----------------------------")
		fmt.Println(packet)
	}
}
