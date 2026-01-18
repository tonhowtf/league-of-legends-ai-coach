package lcu

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	creds *LCUCredentials
	httpClient *http.Client
}


func NewClient() (*Client, error) {
	lockfile := GetLockFile()
	creds := ParseLockFile(lockfile)

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	return &Client{
		creds: creds,
		httpClient: &http.Client{Transport: transport},
	}, nil
}


func (c *Client) MakeRequest(endpoint string) ([]byte, error) {	
	
	
	url := fmt.Sprintf("%s://127.0.0.1:%s%s", c.creds.Protocol, c.creds.Port, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth("riot", c.creds.Password)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}