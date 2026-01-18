package lcu

import (
	"crypto/tls"
	"io"
	"net/http"
)


func LOLREQ(endpoint string) string {

	lockfile := GetLockFile()
	
	creds := ParseLockFile(lockfile)
	
	url := creds.Protocol + "://" + "127.0.0.1:" + creds.Port + endpoint


	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Transport: transport,
	}

	req, err := http.NewRequest("get", url, nil)
	if err != nil {
		panic(err)
	}
	
	req.SetBasicAuth("riot", creds.Password)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}