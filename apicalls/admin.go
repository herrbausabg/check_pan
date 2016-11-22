package apicalls

import (
        "encoding/xml"
        "strings"
)

func GetAdmins(host string, token string) string {

// definition of the xml struct

type Admin struct {
        Admin string `xml:"admin"`
        From string `xml:"from"`
        Type string `xml:"type"`
        Start string `xml:"start-from"`
        Idle string `xml:"idle-for"`
}

type Entries struct {
    XMLName xml.Name `xml:"result"`
    Adminlist []Admin `xml:"admins>entry"`
}

type Response struct {
    XMLName xml.Name `xml:"response"`
    Entries Entries `xml:"result"`
}

// definition of the cast uri

const uri = "/api/?type=op&cmd=<show><admins><%2Fadmins><%2Fshow>"


var result = ""
var url string
resp := new(Response)



    s :=  []string{"https://",host,uri,"&key=",token}
    url = strings.Join(s,"")

    CallRestapi(url, resp)
        
        for _, v := range resp.Entries.Adminlist {
            result = strings.Join([]string{result,"Admin:",v.Admin," logged in from the:",v.Type," with IP:",v.From," at:",v.Start," and idles for:",v.Idle,".\n"},"")
            }

return result    
}




