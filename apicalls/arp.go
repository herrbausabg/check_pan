package apicalls

import (
        "encoding/xml"
        "strings"
)

func GetArps(host string, token string) string {

// definition of the xml struct

type Entry struct {
        Status string `xml:"status"`
        Ip string `xml:"ip"`
        Mac string `xml:"mac"`
        Ttl string `xml:"ttl"`
        Interface string `xml:"interface"`
        Port string `xml:"port"`
}

type Entries struct {
    XMLName xml.Name `xml:"result"`
    Entrylist []Entry `xml:"entries>entry"`
}

type Response struct {
    XMLName xml.Name `xml:"response"`
    Entries Entries `xml:"result"`
}

// definition of the cast uri

const uri = "/api/?type=op&cmd=<show><arp><entry+name+%3D+%27all%27%2F><%2Farp><%2Fshow>"


var result = ""
var url string
resp := new(Response)
   
    
s :=  []string{"https://",host,uri,"&key=",token}
url = strings.Join(s,"")
    
         
CallRestapi(url, resp)

for _, v := range resp.Entries.Entrylist {
            result = strings.Join([]string{result,"Status:",v.Status," IP:",v.Ip," MAC:",v.Mac," at:",v.Interface," TTL:",v.Ttl,".\n"},"")
            }

return result    
}


