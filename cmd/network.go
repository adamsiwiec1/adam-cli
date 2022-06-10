package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
)

type Address struct {
	City       string
	RegionName string
	Zip        string
	Isp        string
	Query      string
}

var whoIsMe = &cobra.Command{
	Use:   "whoisme",
	Short: "Returns location, ISP, and public ip address.",
	Run: func(cmd *cobra.Command, args []string) {
		req, err := http.Get("http://ip-api.com/json/")
		if err != nil {
			fmt.Printf("Error:", err)
		}
		defer req.Body.Close()

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Printf("Error:", err)
		}

		var addr Address
		json.Unmarshal(body, &addr)

		fmt.Printf("Location: %s, %s %s\n", addr.City, addr.RegionName, addr.Zip)
		fmt.Printf("ISP: %s\n", addr.Isp)
		fmt.Printf("Public IP: %s\n", addr.Query)
	},
}

var publicIp = &cobra.Command{
	Use:   "publicip",
	Short: "returns public ip address",
	Run: func(cmd *cobra.Command, args []string) {
		req, err := http.Get("http://ip-api.com/json/")
		if err != nil {
			fmt.Printf("Error:", err)
		}
		defer req.Body.Close()

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Printf("error:", err)
		}

		var addr Address
		json.Unmarshal(body, &addr)

		fmt.Println(addr.Query)
	},
}

var privateIp = &cobra.Command{
	Use:   "privateip",
	Short: "returns list of private ip addresses on machine",
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := os.Hostname()
		ips, _ := net.LookupIP(host)
		var count = 0 // need 2 use count bc index is not correct
		if len(ips) > 1 {
			for _, addr := range ips {
				if ipv4 := addr.To4(); ipv4 != nil {
					fmt.Printf("ipv4 #%s: %s\n", strconv.Itoa(count), ipv4)
					count++
				}
			}
		} else {
			fmt.Println("error pulling private ip addresses..")
		}
	},
}

var devicesOnNetwork = &cobra.Command{
	Use:   "netdev",
	Short: "finds all device ips on network",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(whoIsMe)
	rootCmd.AddCommand(publicIp)
	rootCmd.AddCommand(privateIp)
}
