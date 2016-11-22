package apicalls

import (
        "encoding/xml"
        "strings"
)

func GetLics(host string, token string) string {

// definition of the xml struct

type Entry struct {
        Feature string `xml:"feature"`
        Desc string `xml:"description"`
        Issued string `xml:"issued"`
        Expires string `xml:"expires"`
        Expired string `xml:"expired"`
}

type Entries struct {
    XMLName xml.Name `xml:"result"`
    Entrylist []Entry `xml:"licenses>entry"`
}

type Response struct {
    XMLName xml.Name `xml:"response"`
    Entries Entries `xml:"result"`
}

// definition of the cast uri

const uri = "/api/?type=op&cmd=<request><license><info><%2Finfo><%2Flicense><%2Frequest>"


var result = ""
var url string
resp := new(Response)



    s :=  []string{"https://",host,uri,"&key=",token}
    url = strings.Join(s,"")

    CallRestapi(url, resp)
        
        for _, v := range resp.Entries.Entrylist {
            result = strings.Join([]string{result,"Feature: ",v.Feature," Description: ",v.Desc," issued: ",v.Issued, " expires: ",v.Expires," and is expired: ",v.Expired,".\n"},"")
            }

return result    
}




