package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

var (
	device           string = ""
	snapshot_len     int32  = 1024
	promiscuous      bool   = false
	err              error
	timeout          time.Duration = 30 * time.Second
	handle           *pcap.Handle
	snapshotLen      uint32 = 1024
	file_writer      *pcapgo.Writer
	should_save_pcap bool = false
)

func init() {
	flag.BoolVar(&should_save_pcap, "pcap", false, "Wrong")
	flag.BoolVar(&should_save_pcap, "p", false, "Wrong")
	flag.StringVar(&device, "interface", "eth0", "Wrong")
}

func main() {
	flag.Parse()

	if should_save_pcap {
		f, _ := os.Create("captured.pcap")
		file_writer = pcapgo.NewWriter(f)
		file_writer.WriteFileHeader(snapshotLen, layers.LinkTypeEthernet)
		defer f.Close()
	}
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
		if should_save_pcap {
			file_writer.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		}
	}
}
