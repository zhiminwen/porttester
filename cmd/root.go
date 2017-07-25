package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "porttester",
	Short: "test if the tcp port is open",
	Long:  "A cross platform TCP port availability tester. Useful for firewall testing both host and external.",
	Run:   port_tester,
}

var host string
var port int

func init() {
	RootCmd.PersistentFlags().StringVarP(&host, "hostname", "o", "", "hostname or ip to connect to")
	RootCmd.PersistentFlags().IntVarP(&port, "port", "p", -1, "TCP port to test")
}

func port_tester(cmd *cobra.Command, args []string) {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		log.Fatalf("Connection Failed: %v", err)
	} else {
		conn.Close()
	}

	log.Printf("Port %d is opened", port)
}
