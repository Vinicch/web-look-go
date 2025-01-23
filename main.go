package main

import (
	"fmt"
	"net"
	"os"
)

func help() {
	fmt.Println("Usage: web-look-go [domain] [command]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("cn: returns the canonical name for the given host")
	fmt.Println("ip: returns the host's IPv4 and IPv6 addresses")
	fmt.Println("mx: returns the DNS MX records for the given domain name sorted by preference")
	fmt.Println("ns: returns the DNS NS records for the given domain name")
}

func main() {
	if len(os.Args) != 3 {
		help()
		return
	}

	domain, command := os.Args[1], os.Args[2]
	switch command {
	case "cn":
		cname, err := net.LookupCNAME(domain)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Println(cname)
	case "ip":
		ipAddrs, err := net.LookupIP(domain)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		for _, addr := range ipAddrs {
			fmt.Println(addr)
		}
	case "mx":
		mxRecords, err := net.LookupMX(domain)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		for _, record := range mxRecords {
			fmt.Println(record.Host, record.Pref)
		}
	case "ns":
		nsRecords, err := net.LookupNS(domain)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		for _, record := range nsRecords {
			fmt.Println(record.Host)
		}
	default:
		help()
	}
}
