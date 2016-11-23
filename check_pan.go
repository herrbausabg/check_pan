package main

import (
        "flag"
        "fmt"
        "net"
        "os"
        "./apicalls"
)

var hostFlag = flag.String("h", "", "Host to check. (required)")
var tokenFlag = flag.String("t", "", "Authorization Token to use. (required)")
var commandFlag = flag.String("c", "", "Command to check. (required). Implemented:\n\tadmin - show admins\n\tarp - show arp all\n\tcert - show certificates\n\tlicense - show license info")
var outputFlag = flag.Int("o", 0, "Outputformat. (optional) Implemented:\n\t0 - human readable(default)\n\t1 - icinga/nagios plugin")







func main() {
    required := []string{"h", "t", "c"}
    host := ""
    flag.Parse()
    seen := make(map[string]bool)
    flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
    for _, req := range required {
        if !seen[req] {
            fmt.Fprintf(os.Stderr, "Error: missing required -%s argument/flag, use -help for more info\n", req)
            os.Exit(1) 
        }
    }

// Errorhandling Hostname / IP Address
    ipaddr := net.ParseIP(*hostFlag)
    if ipaddr == nil {
        addr, err := net.ResolveIPAddr("ip", *hostFlag)
        if err != nil {
            fmt.Println("Resolution error or invalid address", err.Error())
            os.Exit(2)
        }
//        fmt.Println("Hostmame:", *hostFlag, "Address:", addr.String())
        host = addr.String()
    } else {
//        fmt.Println("Host Address:", ipaddr.String())
        host = ipaddr.String()

    }

    

fmt.Println()

switch *commandFlag {
    case "admin": fmt.Println(apicalls.GetAdmins(host, *tokenFlag, *outputFlag))
    case "arp": fmt.Println(apicalls.GetArps(host, *tokenFlag, *outputFlag)) 
    case "cert": fmt.Println(apicalls.GetCerts(host, *tokenFlag, *outputFlag))
    case "license": fmt.Println(apicalls.GetLics(host, *tokenFlag, *outputFlag))
    default: { fmt.Fprintf(os.Stderr,"Error: Unrecognized command with flag -c, use -help for more info\n")
               os.Exit(1)
             }  
}


   

}


    