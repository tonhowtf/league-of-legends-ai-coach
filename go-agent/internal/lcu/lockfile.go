package lcu

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type LCUCredentials struct {
	ProcessName string
	PID string
	Port string
	Password string
	Protocol string
}

func ParseLockFile(items string) *LCUCredentials{
	lockFileParts := strings.Split(string(items), ":")

	return &LCUCredentials{
		ProcessName: lockFileParts[0],
		PID: lockFileParts[1],
		Port: lockFileParts[2],
		Password: lockFileParts[3],
		Protocol: lockFileParts[4],
	}
}


func GetLockFile() string {
	cmd := exec.Command("powershell", "-NoProfile", "-Command", "(Get-Process LeagueClientUx).Path")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	path := GetDir(string(output))

	lockfile, err := ReadLockFile(path)

	if err != nil {
		panic(err)
	}

	return string(lockfile)
}

func GetDir(leagueUX string) string {
	return filepath.Dir(leagueUX)
}

func ReadLockFile(path string) ([]byte, error) {
	filePath := filepath.Join(path, "lockfile")
	return os.ReadFile(filePath)
}