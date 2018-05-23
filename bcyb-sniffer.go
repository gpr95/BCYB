package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	device       string = ""
	snapshot_len int32  = 1024
	promiscuous  bool   = false
	err          error
	timeout      time.Duration = 30 * time.Second
	handle       *pcap.Handle

	should_save_pcap bool = false
)

func init() {
	flag.BoolVar(&should_save_pcap, "pcap", false, "Wrong")
	flag.BoolVar(&should_save_pcap, "p", false, "Wrong")
	flag.StringVar(&device, "interface", "eth0", "Wrong")
}

func main() {
	flag.Parse()
	// Open device
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)
	}
}
