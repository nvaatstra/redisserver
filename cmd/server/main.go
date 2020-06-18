package main

import (
	"flag"
	"fmt"
	"path"
	"runtime"

	"github.com/nvaatstra/redisserver/pkg/server"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Normal:
	log.SetLevel(log.InfoLevel)

	// Debug:
	// log.SetLevel(log.DebugLevel)

	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return "", fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
}

func main() {
	var network string
	var address string

	flag.StringVar(&network, "net", "tcp", "Listen network")
	flag.StringVar(&address, "addr", "127.0.0.1:1234", "Listen address")

	flag.Parse()

	log.Debugf("Listen network: %s", network)
	log.Debugf("Listen address: %s", address)

	s := server.New(network, address)

	s.Run()
}
