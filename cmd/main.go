package main

import (
	"flag"
	"fmt"
	"net/netip"
	"os"

	"github.com/ttl256/dygn/internal/hwaddr"
)

func NewPrefix(prefix netip.Prefix, mac hwaddr.MAC) netip.Addr {
	eui64 := hwaddr.EUIFrom6(mac)
	addr := eui64.AppendToPrefixAddr(prefix.Addr().As16())
	return netip.AddrFrom16(addr)
}

func main() {
	var macAddr *hwaddr.MAC
	var prefix *netip.Prefix

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Func(
		"mac", "MAC address\nAny format will do: xx:xx:xx:xx:xx:xx, xx-xx-xx-xx-xx-xx, xxxx.xxxx.xxxx, xxxxxxxxxxxx",
		func(s string) error {
			addr, err := hwaddr.ParseAddr(s)
			if err != nil {
				return err
			}
			macAddr = &addr
			return nil
		})
	flag.Func("prefix", "IPv6 prefix of prefix-length 64", func(s string) error {
		p, err := netip.ParsePrefix(s)
		if err != nil {
			return err
		}
		prefix = &p
		return nil
	})
	flag.Parse()

	if macAddr == nil || prefix == nil {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println(NewPrefix(*prefix, *macAddr))
}
