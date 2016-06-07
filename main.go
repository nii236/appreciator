package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/nii236/appreciator/cmd"
)

func main() {
	f := &log.TextFormatter{}
	f.ForceColors = true
	if os.Getenv("DEBUG") != "" {
		log.SetLevel(log.DebugLevel)
	}
	log.SetFormatter(f)
	cmd.Execute()
}
