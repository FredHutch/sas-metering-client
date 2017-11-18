package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/StackExchange/wmi"
)

//Win32Process to hold properties returned from Win32_Process WMI call
type Win32_Process struct {
	Name                string `json:"process_name"`
	ProcessId           int    `json:"process_id"`
	CommandLine         string `json:"command_line"`
	ExecutablePath      string `json:"executable_path"`
	WindowsVersion      string `json:"windows_os_version"`
	ThreadCount         int    `json:"thread_count"`
	UserModeTime        uint64 `json:"cpu_user_time_ms"`
	KernelModeTime      uint64 `json:"cpu_kernel_time_ms"`
	WorkingSetSize      uint64 `json:"memory_usage_bytes"`
	WriteOperationCount uint64 `json:"write_operations"`
	WriteTransferCount  uint64 `json:"write_bytes"`
	ReadOperationCount  uint64 `json:"read_operations"`
	ReadTransferCount   uint64 `json:"read_bytes"`
}

//rename to main to run standalone
func proc() {
	http.HandleFunc("/ps", psHandler)
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "")
}

func psHandler(w http.ResponseWriter, r *http.Request) {
	var dst []Win32_Process
	q := wmi.CreateQuery(&dst, "")
	//q := wmi.CreateQuery(&dst, "WHERE name = 'chrome.exe'")
	err := wmi.Query(q, &dst)
	if err != nil {
		log.Fatal(err)
	}

	// convert 100 nanoscond time units to 1ms time units
	for idx := 0; idx < len(dst); idx++ {
		dst[idx].KernelModeTime = dst[idx].KernelModeTime / 10000
		dst[idx].UserModeTime = dst[idx].UserModeTime / 10000
	}

	p, err := json.MarshalIndent(dst, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(p))
}
