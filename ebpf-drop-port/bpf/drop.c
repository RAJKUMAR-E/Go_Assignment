#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_endian.h>

#define TCP_PROTOCOL 6

struct {
    __uint(type, BPF_MAP_TYPE_ARRAY);
    __uint(max_entries, 1);
    __type(key, __u32);
    __type(value, __u16); // Port number
} blocked_port SEC(".maps");

SEC("xdp")
int drop_tcp_port(struct xdp_md *ctx) {
    void *data_end = (void *)(long)ctx->data_end;
    void *data = (void *)(long)ctx->data;

    struct ethhdr *eth = data;
    if ((void *)(eth + 1) > data_end) return XDP_PASS;

    if (bpf_ntohs(eth->h_proto) != ETH_P_IP) return XDP_PASS;

    struct iphdr *ip = data + sizeof(*eth);
    if ((void *)(ip + 1) > data_end) return XDP_PASS;

    if (ip->protocol != TCP_PROTOCOL) return XDP_PASS;

    int ip_hdr_len = ip->ihl * 4;
    struct tcphdr *tcp = data + sizeof(*eth) + ip_hdr_len;
    if ((void *)(tcp + 1) > data_end) return XDP_PASS;

    __u32 key = 0;
    __u16 *target_port = bpf_map_lookup_elem(&blocked_port, &key);
    if (!target_port) return XDP_PASS;

    if (bpf_ntohs(tcp->dest) == *target_port) {
        bpf_printk("Dropping TCP packet on port %d\n", *target_port);
        return XDP_DROP;
    }

    return XDP_PASS;
}

char LICENSE[] SEC("license") = "GPL";
