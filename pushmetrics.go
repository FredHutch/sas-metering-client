package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func pushMetrics() {

	/// Skip SSL Cert verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	/// End Skip

	hostname, err := os.Hostname()
	if err != nil {
		log.Print(err)
	}

	for {

		time.Sleep(time.Second * 60)

		var running int
		if checkrunning("sas.exe") == true {
			running = 1
		} else {
			running = 0
		}

		metric := fmt.Sprintf("# TYPE sas_running gauge\n# HELP sas_running Is sas.exe running or not (1=true, 0=false)\nsas_running %d\n", running)

		body := strings.NewReader(metric)
		URL := fmt.Sprintf("https://prometheus.fhcrc.org:9991/metrics/job/sas_desktops/instance/%s", hostname)
		req, err := http.NewRequest("POST", URL, body)
		if err != nil {
			log.Print(err)
			continue
		}
		req.SetBasicAuth("user", "pass")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		//resp, err := http.DefaultClient.Do(req)
		resp, err := client.Do(req) //replace with line above and comment out skip section to re-enable  SSL verfication
		if err != nil {
			log.Print(err)
			continue
		}
		defer resp.Body.Close()
	}
}
