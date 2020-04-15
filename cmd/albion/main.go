package main

import (
	"github.com/mvaude/albion/internal/commands"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	commands.Execute()
}
