package apicalls

import (
        "encoding/xml"
        "fmt"
        "strings"
        "time"
)

func GetLics(host string, token string, output int) string {

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

    switch output {
    case 0: {            
       for _, v := range resp.Entries.Entrylist {
            result = strings.Join([]string{result,"Feature: ",v.Feature," Description: ",v.Desc," issued: ",v.Issued, " expires: ",v.Expires," and is expired: ",v.Expired,".\n"},"")
                }
            }
    case 1: {                
// Setup Icinga Output
            short := "January 2, 2006"

            

            for _, v := range resp.Entries.Entrylist {
                t, _ := time.Parse(short, v.Expires)
                if (time.Now().Sub(t).Hours() >= 0){
                fmt.Printf("%s license expired %.1f hours ago\n", v.Feature, time.Now().Sub(t).Hours())
                        } else {
                fmt.Printf("%s license expires in %.1f hours\n", v.Feature, time.Now().Sub(t).Hours()*(-1))    
                    }        

                }
            }    
    default: {
            panic("No valid outputformat given")
            }        

    }        

return result    
}




