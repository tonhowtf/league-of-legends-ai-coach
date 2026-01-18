package lcu

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func (c *Client) ConnectWebSocket() {

	dialer := websocket.Dialer{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}


	url := fmt.Sprintf("wss://127.0.0.1:%s", c.creds.Port)

	header := http.Header{}
	header.Set(
		"authorization",
		"Basic "+base64.StdEncoding.EncodeToString(
			[]byte("riot:"+c.creds.Password),
		),
	)

	conn, _, err := dialer.Dial(url, header)
	if err != nil {
		panic(err)
	}
	defer conn.Close()


	err = conn.WriteJSON([]interface{}{5, "OnJsonApiEvent"})
	if err != nil {
		return
	}

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			return 
		}

		if messageType == websocket.TextMessage {
			fmt.Println(string(message))
		}
	}
	
}