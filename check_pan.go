package main

import (
		"flag"
		"fmt"
		"net"
		"os"
)

var hostFlag = flag.String("h", "", "Host to check. (required)")
var tokenFlag = flag.String("t", "", "Authorization Token to use. (required)")



func main() {
	required := []string{"h", "t"}
	flag.Parse()
	// Host and Token are required, build a map for all flags and test against required
	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
    for _, req := range required {
        if !seen[req] {
            fmt.Fprintf(os.Stderr, "missing required -%s argument/flag, use -help for more info\n", req)
            os.Exit(2) 
        }
    }

//	fmt.Println(*hostFlag)
// Errorhandling Hostname / IP Address
    ipaddr := net.ParseIP(*hostFlag)
    if ipaddr == nil {
    	addr, err := net.ResolveIPAddr("ip", *hostFlag)
    	if err != nil {
    		fmt.Println("Resolution error or invalid address", err.Error())
        	os.Exit(1)
    	}
    	fmt.Println("Hostmame:", *hostFlag, "Address:", addr.String())
    } else {
    	fmt.Println("Host Address:", ipaddr.String())
    }

	fmt.Println(*tokenFlag)
		
}	




	