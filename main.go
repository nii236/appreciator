package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/nii236/appreciator/cmd"
)

func main() {
	f := &log.TextFormatter{}
	f.ForceColors = true
	log.SetFormatter(f)
	cmd.Execute()
}
