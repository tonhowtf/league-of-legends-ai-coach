package main

import "github.com/tonhowtf/lol-agent/internal/lcu"

func main() {
	
	client, err := lcu.NewClient()
	if err != nil {
		panic(err)
	}

	data, _ := client.MakeRequest("/riotclient/region-locale")
	println(string(data))

	client.ConnectWebSocket()

}
