package main

import (
	"bufio"
	"fmt"
	"github.com/Ullaakut/nmap/v2"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
)

func contains(l []string, x string) bool {
	for _, y := range l {
		if x == y {
			return true
		}
	}
	return false
}

type Device struct {
	name string
	ip   string
}

var reader = bufio.NewReader(os.Stdin)

//func scan() []Device {
//	var deviceList []Device;
//
//
//
//	append(deviceList, )
//
//	return
//}

func grabPublicIp() {
	host, _ := os.Hostname()
	ips, _ := net.LookupIP(host)
	var desiredCidr string = ""
	if len(ips) > 1 {
		var cidrRanges []string
		for _, addr := range ips {
			if ipv4 := addr.To4(); ipv4 != nil {
				re := regexp.MustCompile("\\d{1,3}")
				var x = re.FindString(ipv4.String())
				if !contains(cidrRanges, x) {
					cidrRanges = append(cidrRanges, x)
				}
			}

		}
		if len(cidrRanges) > 1 {
			fmt.Println("we found multiple cidr ranges on your network.. choose which one to scan:")

			// print cidrs
			for i, x := range cidrRanges {
				fmt.Printf("[%s] %s\n", strconv.Itoa(i), x)
			}

			// ask for input
			fmt.Printf("number (0,1,etc..):")
			text, _ := reader.ReadString('\n')
			j := 0
			j, _ = strconv.Atoi(text)
			desiredCidr = cidrRanges[j]

			// else, if there is only 1 cidr range
		} else {
			var r string = cidrRanges[0]
			desiredCidr = r
		}

		scan, err := nmap.NewScanner(
			nmap.WithTargets(fmt.Sprintf("%s.0.0.0/24", desiredCidr)),
		)
		if err != nil {
			log.Fatalf("unable to create nmap scanner: %v", err)
		}
		fmt.Println("scanning..")
		result, _, err := scan.Run()
		if err != nil {
			log.Fatalf("nmap scan failed: %v", err)
		} else {
			fmt.Println("scan succeded.")
		}

		for _, host := range result.Hosts {
			fmt.Printf("%s %s\n", host.Addresses[0], host.Hostnames[0])
		}
	} else {
		fmt.Printf("error pulling private ip addresses.. %v", ips)
	}

}

func main() {
	grabPublicIp()
}
