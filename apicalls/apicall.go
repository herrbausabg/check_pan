package apicalls

import (
        "crypto/tls"
        "encoding/xml"
        "net/http"
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

	return xml.NewDecoder(r.Body).Decode(target)
} 