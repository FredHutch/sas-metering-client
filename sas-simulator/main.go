package main

import (
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	s := spinner.New(spinner.CharSets[28], 500*time.Millisecond)
	s.Prefix = "sas.exe running: "
	s.Start()
	defer s.Stop()
	for {
		time.Sleep(time.Second * 5)
	}
}
