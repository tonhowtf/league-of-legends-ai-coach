package main

import (
	"fmt"

	"github.com/tonhowtf/lol-agent/internal/lcu"
)
func main() {
	lockfile := lcu.GetLockFile()
	
	creds := lcu.ParseLockFile(lockfile)

	fmt.Println(*creds)


}
