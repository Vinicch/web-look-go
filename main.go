package main

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
)

var address string

func main() {
	rootCmd := &cobra.Command{
		Use:     "wlc",
		Short:   "Let's you query IPs, CNAMEs, MX records and Name Servers!",
		Version: "1.1.1",
	}

	cn(rootCmd)
	ip(rootCmd)
	mx(rootCmd)
	ns(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func cn(rootCmd *cobra.Command) {
	cnCmd := &cobra.Command{
		Use:   "cn",
		Short: "Looks up the Common Name for a particular host",
		RunE: func(cmd *cobra.Command, args []string) error {
			commonName, err := net.LookupCNAME(address)
			if err != nil {
				return err
			}
			fmt.Println(commonName)
			return nil
		},
	}

	cnCmd.Flags().StringVarP(&address, "address", "a", "google.com", "Sets the host to be looked up")
	rootCmd.AddCommand(cnCmd)
}

func ip(rootCmd *cobra.Command) {
	ipCmd := &cobra.Command{
		Use:   "ip",
		Short: "Looks up the IP addresses for a particular host",
		RunE: func(cmd *cobra.Command, args []string) error {
			ips, err := net.LookupIP(address)
			if err != nil {
				return err
			}
			for _, ip := range ips {
				fmt.Println(ip)
			}
			return nil
		},
	}

	ipCmd.Flags().StringVarP(&address, "address", "a", "google.com", "Sets the host to be looked up")
	rootCmd.AddCommand(ipCmd)
}

func mx(rootCmd *cobra.Command) {
	mxCmd := &cobra.Command{
		Use:   "mx",
		Short: "Looks up the Mail eXchange DNS for a particular host",
		RunE: func(cmd *cobra.Command, args []string) error {
			mxRecords, err := net.LookupMX(address)
			if err != nil {
				return err
			}
			for _, record := range mxRecords {
				fmt.Println(record.Host, record.Pref)
			}
			return nil
		},
	}

	mxCmd.Flags().StringVarP(&address, "address", "a", "google.com", "Sets the host to be looked up")
	rootCmd.AddCommand(mxCmd)
}

func ns(rootCmd *cobra.Command) {
	nsCmd := &cobra.Command{
		Use:   "ns",
		Short: "Looks up the Name Servers for a particular host",
		RunE: func(cmd *cobra.Command, args []string) error {
			nameServers, err := net.LookupNS(address)
			if err != nil {
				return err
			}
			for _, nameServer := range nameServers {
				fmt.Println(nameServer.Host)
			}
			return nil
		},
	}

	nsCmd.Flags().StringVarP(&address, "address", "a", "google.com", "Sets the host to be looked up")
	rootCmd.AddCommand(nsCmd)
}
