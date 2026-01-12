package main

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type Process struct {
		PID  uint32
		Name string
	}

func main() {
	
}


func FindProcessByName(name string) (*Process, error) {
	processes, err := ListProcesses()
	if err != nil {
		return nil, err
	}

	for _, p := range processes {
		if p.Name == name {
			return &p, nil
		}
	}

	return nil, nil
}

func ListProcesses() ([]Process, error) {
	snap, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return nil, err
	}

	defer windows.CloseHandle(snap)

	var entry windows.ProcessEntry32
	entry.Size = uint32(unsafe.Sizeof(entry))

	if err = windows.Process32First(snap, &entry); err != nil {
		return nil, err
	}

	var processes []Process

	for {
		processes = append(processes, Process{
			PID:  entry.ProcessID,
			Name: windows.UTF16ToString(entry.ExeFile[:]),
		})

		if err = windows.Process32Next(snap, &entry); err != nil {
			break
		}
	}

	return processes, nil
}