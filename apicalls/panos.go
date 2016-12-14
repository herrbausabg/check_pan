package apicalls

import (
        "encoding/xml"
        "fmt"
        "strings"
//        "time"
)

func GetPanos(host string, token string, output int) int{

// definition of the xml struct

type Entry struct {
        Version string `xml:"version"`
        Filename string `xml:"filename"`
        Released string `xml:"released-on"`
        Downloaded string `xml:"downloaded"`
        Uploaded string `xml:"uploaded"`
        Latest string `xml:"latest"`
        Current string `xml:"current"`
        Size string `xml:"size"`
}

type Entries struct {
    XMLName xml.Name `xml:"result"`
    Entrylist []Entry `xml:"sw-updates>versions>entry"`
}

type Response struct {
    XMLName xml.Name `xml:"response"`
    Entries Entries `xml:"result"`
}

// definition of the cast uri

const uri = "/api/?type=op&cmd=<request><system><software><check><%2Fcheck><%2Fsoftware><%2Fsystem><%2Frequest>"


var result = ""
var url string
var CurrentVersion string
var LatestVersion string
var LatestVersionIsKnown = false
var critical = false
var exitCode = 0
resp := new(Response)



    s :=  []string{"https://",host,uri,"&key=",token}
    url = strings.Join(s,"")

    CallRestapi(url, resp)

    switch output {
    case 0: {            
 
       for _, v := range resp.Entries.Entrylist {
            result = strings.Join([]string{result,"Version: ",v.Version," Filename: ",v.Filename," released-on: ",v.Released," downloaded: ",v.Downloaded," current: ",v.Current," latest: ",v.Latest,".\n"},"")
                }
            fmt.Println(result)   
            exitCode = 0    
            }
    case 1: {   

// Setup Icinga Output




            for _, v := range resp.Entries.Entrylist {
                if (v.Latest == "yes") && (v.Current == "no"){
                LatestVersion = v.Version
                LatestVersionIsKnown = true
                critical = true
                } 
                if (v.Latest == "yes") && (v.Current == "yes"){
                LatestVersion = v.Version
                LatestVersionIsKnown = true
                }
                if (v.Latest == "no") && (v.Current == "yes"){
                CurrentVersion = v.Version    
                }
            }    
                


                if (critical && LatestVersionIsKnown) {fmt.Println("PANOS CRITICAL || Installed Version: ", CurrentVersion, " Latest Version: ", LatestVersion)
                            exitCode = 2
                    } else { 
                            if !LatestVersionIsKnown {fmt.Println("PANOS WARNING || No Latest Version - Support License Problem? Installed Version: ", CurrentVersion)
                                        exitCode = 1
                           } else { fmt.Println("PANOS OK || Latest Version ", LatestVersion, " installed.")
                                        exitCode = 0}


                } 

            }    
    default: {
            panic("No valid outputformat given")
            }        

    }        

return exitCode   
}




