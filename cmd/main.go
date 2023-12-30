package main

import (
	"github.com/AliceDiNunno/hid"
	"github.com/sirupsen/logrus"
	hub "godeck/src/adapters/events/hub"
	"godeck/src/core/Framework"
	"godeck/src/core/OS"
	"godeck/src/core/connector"
	"os"
)

func main() {
	log := logrus.New()

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown host"
	}
	log.WithFields(logrus.Fields{
		"hostname": hostname,
	}).Println("Framework V0.1")

	var eventHub = hub.NewHub()

	isSupported := hid.Supported()

	if !isSupported {
		log.Errorln("HID is not supported on this platform")
		os.Exit(1)
	}

	var framework connector.FrameworkConnector = nil

	osbuilder := func() connector.OSConnector {
		prodos := OS.StartOS(framework, eventHub)
		return prodos
	}

	proos := Framework.NewProdeckFramework(eventHub, osbuilder)

	framework = proos

	proos.Start()

	exit := make(chan string)

	for {
		select {
		case <-exit:
			os.Exit(0)
		}
	}
}
