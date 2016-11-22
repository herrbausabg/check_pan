package apicalls

import (
        "encoding/xml"
        "strings"
)

func GetCerts(host string, token string) string {

// definition of the xml struct

type Entry struct {
        Issuer string `xml:"issuer"`
        ValidTo string `xml:"not-valid-after"`
        Cn string `xml:"common-name"`
        Priv string `xml:"private-key"`
        Ca string `xml:"ca"`
}

type Entries struct {
    XMLName xml.Name `xml:"result"`
    Certlist []Entry `xml:"certificate>entry"`
}

type Response struct {
    XMLName xml.Name `xml:"response"`
    Entries Entries `xml:"result"`
}

// definition of the cast uri

const uri = "/api/?type=config&action=get&xpath=%2Fconfig%2Fshared%2Fcertificate"


var result = "Certificates with private Keys on Target: "+host+"\n"
var url string
resp := new(Response)
   
    
s :=  []string{"https://",host,uri,"&key=",token}
url = strings.Join(s,"")
    
         
CallRestapi(url, resp)


for _, v := range resp.Entries.Certlist {
	if v.Priv == "********" {
			result = strings.Join([]string{result,"CN:",v.Cn," Issuer:",v.Issuer," valid-to:",v.ValidTo," is Ca:",v.Ca, "Privkey:",v.Priv,"\n"},"")
	          }
            }

return result    
}


