package apicalls

import (
        "crypto/tls"
        "encoding/xml"
        "fmt"
        "net/http"
        "os"
        "strings"
)

var result = ""

type Entry struct {
        Admin string `xml:"admin"`
        From string `xml:"from"`
        Type string `xml:"type"`
        Start string `xml:"start-from"`
        Idle string `xml:"idle-for"`
}

type Entries struct {
    XMLName xml.Name `xml:"result"`
    Entrylist []Entry `xml:"admins>entry"`
}

type Response struct {
    XMLName xml.Name `xml:"response"`
    Entries Entries `xml:"result"`
}

func GetAdmins(host string, tokenFlag string) string {

 tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    // show admins
    const uri = "/api/?type=op&cmd=<show><admins><%2Fadmins><%2Fshow>"
    response, err := client.Get("https://"+host+uri+"&key="+tokenFlag)
    if err != nil {
        fmt.Println(err)
    }   else {
        defer response.Body.Close()
        d := xml.NewDecoder(response.Body)
        var resp Response
        if err := d.Decode(&resp); err != nil {
        // Handle error
        fmt.Println("Could not get result", err.Error())
            os.Exit(1)    
        } else {
        // No error
        
        for _, v := range resp.Entries.Entrylist {
            result = strings.Join([]string{result,"Admin:",v.Admin,"logged in from the:",v.Type,"with IP:",v.From,"at:",v.Start," and idles for:",v.Idle,".\n"}," ")
            }





        }

    }
    
return result    
}