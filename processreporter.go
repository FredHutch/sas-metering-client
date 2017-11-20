package main

import (
	"fmt"
	"log"

	"github.com/StackExchange/wmi"
)

//Win32Process to hold properties returned from Win32_Process WMI call
type Win32_Process struct {
	Name                string
	ProcessId           int
	CommandLine         string
	ExecutablePath      string
	WindowsVersion      string
	ThreadCount         int
	UserModeTime        uint64
	KernelModeTime      uint64
	WorkingSetSize      uint64
	WriteOperationCount uint64
	WriteTransferCount  uint64
	ReadOperationCount  uint64
	ReadTransferCount   uint64
}

//rename to main to run standalone
func checkrunning(name string) bool {

	var dst []Win32_Process
	//q := wmi.CreateQuery(&dst, "")
	queryString := fmt.Sprintf("WHERE name = '%s'", name)
	q := wmi.CreateQuery(&dst, queryString)
	//q := wmi.CreateQuery(&dst, "WHERE name = 'sas.exe'")
	err := wmi.Query(q, &dst)
	if err != nil {
		log.Fatal(err)
	}

	// convert 100 nanoscond time units to 1ms time units
	for idx := 0; idx < len(dst); idx++ {
		dst[idx].KernelModeTime = dst[idx].KernelModeTime / 10000
		dst[idx].UserModeTime = dst[idx].UserModeTime / 10000
	}

	if len(dst) > 0 {
		return true
	}

	return false

}
