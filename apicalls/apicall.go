package apicalls

import (
        "crypto/tls"
        "encoding/xml"
        "fmt"
        "net/http"
        "os"
)



func CallRestapi(url string, target interface{}) error {
	
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
	
	r, err := client.Get(url)


	if err != nil {
		return err
	}
	defer r.Body.Close()

	if v, ok :=  r.Header["Status"]; ok {
		fmt.Println("Targetsystem returns error in HTTP Response:", v)
		os.Exit(2)
	}


	return xml.NewDecoder(r.Body).Decode(target)
} 