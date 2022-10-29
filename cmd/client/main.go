package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/alexandrevilain/teleinfo-timescaledb/pkg/log"
	teleinfo "github.com/j-vizcaino/goteleinfo"
	"github.com/spf13/pflag"
)

func main() {
	err := startClient()
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}
}

func startClient() error {
	logger, err := log.CreateLogger()
	if err != nil {
		return err
	}
	logger.Info("Starting client")

	f := pflag.NewFlagSet("", pflag.ContinueOnError)
	f.Usage = func() {
		fmt.Println(f.FlagUsages())
		os.Exit(0)
	}
	cfg := &Config{}
	cfg.AddFlags(f)

	if err := f.Parse(os.Args[1:]); err != nil {
		return err
	}

	if err := cfg.Validate(); err != nil {
		return err
	}

	logger.Info("Starting teleinfo reader", "mode", cfg.Teleinfo.Mode, "device", cfg.SerialDevice)
	port, err := teleinfo.OpenPort(cfg.SerialDevice, cfg.Teleinfo.Mode)
	if err != nil {
		return err
	}
	defer port.Close()

	reader := teleinfo.NewReader(port, &cfg.Teleinfo.Mode)
	for {
		logger.V(1).Info("Reading frames")
		frame, err := reader.ReadFrame()
		if err != nil {
			logger.Error(err, "Error reading teleinfo frame")
			continue
		}

		data, err := json.Marshal(frame)
		if err != nil {
			logger.Error(err, "Error marshaling teleinfo frame")
			continue
		}

		u := url.URL{
			Scheme: cfg.Server.Scheme,
			Host:   cfg.Server.Host,
			Path:   cfg.Server.Scheme,
		}

		resp, err := http.Post(u.String(), "application/json", bytes.NewBuffer(data))
		if err != nil {
			logger.Error(err, "Error sending teleinfo frame")
			continue
		}

		result, err := httputil.DumpResponse(resp, true)
		if err != nil {
			logger.Error(err, "Error dumping teleinfo frame response")
			continue
		}

		logger.Info(string(result))
	}
}
