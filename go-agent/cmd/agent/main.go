package main

//http.Get(creds.Protocol + "://" + "127.0.0.1:" + creds.Port)
import "github.com/tonhowtf/lol-agent/internal/lcu"

func main() {
	
	println(lcu.LOLREQ("/riotclient/region-locale"))

}
