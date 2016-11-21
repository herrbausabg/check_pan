package main

import (
        "crypto/tls"
        "encoding/xml"
        "flag"
        "fmt"
//      "io"
        "net"
        "net/http"
        "os"
)

var hostFlag = flag.String("h", "", "Host to check. (required)")
var tokenFlag = flag.String("t", "", "Authorization Token to use. (required)")
var host = ""


type Item struct {
        Name string `xml:"admin"`
        From string `xml:"from"`
        Type string `xml:"type"`
        Start string `xml:"start-from"`
        Idle string `xml:"idle-for"`
}

type Items struct {
    XMLName xml.Name `xml:"result"`
    Itemlist []Item `xml:"admins>entry"`
}

type Response struct {
    XMLName xml.Name `xml:"response"`
    Items Items `xml:"result"`
}






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

//  fmt.Println(*hostFlag)
// Errorhandling Hostname / IP Address
    ipaddr := net.ParseIP(*hostFlag)
    if ipaddr == nil {
        addr, err := net.ResolveIPAddr("ip", *hostFlag)
        if err != nil {
            fmt.Println("Resolution error or invalid address", err.Error())
            os.Exit(1)
        }
        fmt.Println("Hostmame:", *hostFlag, "Address:", addr.String())
        host = addr.String()
    } else {
        fmt.Println("Host Address:", ipaddr.String())
        host = ipaddr.String()

    }

    fmt.Println(*tokenFlag)
    
    // Setup insecure transport (we know what we are connecting to)

    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    // show admins
    const uri = "/api/?type=op&cmd=<show><admins><%2Fadmins><%2Fshow>"
    response, err := client.Get("https://"+host+uri+"&key="+*tokenFlag)
    if err != nil {
        fmt.Println(err)
    }   else {
        defer response.Body.Close()
//      _, err := io.Copy(os.Stdout, response.Body)
//      if err != nil {
//        fmt.Println(err) }
        d := xml.NewDecoder(response.Body)
        var resp Response
        if err := d.Decode(&resp); err != nil {
        // Handle error
        fmt.Println("Could not get result", err.Error())
            os.Exit(1)    
        } else {
        // No error
        fmt.Printf("%+v", resp)
        }
    
//    fmt.Println(response)

}   

}


    