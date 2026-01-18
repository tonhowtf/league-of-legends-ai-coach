package main

//http.Get(creds.Protocol + "://" + "127.0.0.1:" + creds.Port)
import "github.com/tonhowtf/lol-agent/internal/lcu"

func main() {
	
	client, err := lcu.NewClient()
	if err != nil {
		panic(err)
	}

	data, _ := client.MakeRequest("/riotclient/region-locale")
	println(string(data))

}
