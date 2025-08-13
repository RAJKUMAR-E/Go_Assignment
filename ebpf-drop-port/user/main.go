package main

import (
    "log"

    "github.com/cilium/ebpf"
    "github.com/cilium/ebpf/link"
)

const defaultPort uint16 = 4040

func main() {
    spec, err := ebpf.LoadCollectionSpec("user/dropobj.bpf.o")
    if err != nil {
        log.Fatalf("failed to load spec: %v", err)
    }

    coll, err := ebpf.NewCollection(spec)
    if err != nil {
        log.Fatalf("failed to create collection: %v", err)
    }

    prog := coll.Programs["drop_tcp_port"]
    if prog == nil {
        log.Fatalf("program not found")
    }

    // Attach to XDP on interface (e.g., eth0)
    link, err := link.AttachXDP(link.XDPOptions{
        Program:   prog,
        Interface: 2, // Replace with your interface index
    })
    if err != nil {
        log.Fatalf("failed to attach XDP: %v", err)
    }
    defer link.Close()

    // Set port in map
    portMap := coll.Maps["blocked_port"]
    key := uint32(0)
    port := defaultPort
    if err := portMap.Put(key, port); err != nil {
        log.Fatalf("failed to set port: %v", err)
    }

    log.Printf("Blocking TCP packets on port %d", port)
    select {}
}
