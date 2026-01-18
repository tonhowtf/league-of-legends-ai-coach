package lcu

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func GetLockFile() {
	cmd := exec.Command("powershell", "-NoProfile", "-Command", "(Get-Process LeagueClientUx).Path")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))

	path := GetDir(string(output))

	lockfile, err := ReadLockFile(path)

	if err != nil {
		panic(err)
	}

	println(string(lockfile))
}

func GetDir(leagueUX string) string {
	return filepath.Dir(leagueUX)
}

func ReadLockFile(path string) ([]byte, error) {
	filePath := filepath.Join(path, "lockfile")
	return os.ReadFile(filePath)
}