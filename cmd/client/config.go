package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

// Config is the client configuration struct.
type Config struct {
	SerialDevice string
	Teleinfo     struct {
		Mode string
	}
	Server struct {
		Host   string
		Scheme string
		Path   string
	}
}

// AddFlags adds flags related to config to the specified FlagSet.
func (c *Config) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.SerialDevice, "device", "/dev/ttyUSB0", "Serial port to read frames from")
	fs.StringVar(&c.Teleinfo.Mode, "teleinfo.mode", "historic", "Teleinfo mode standard or historic")
	fs.StringVar(&c.Server.Host, "server.host", "localhost:9000", "HTTP service host to send frames to")
	fs.StringVar(&c.Server.Scheme, "server.scheme", "https", "HTTP service scheme to send frames to")
	fs.StringVar(&c.Server.Path, "server.path", "/v1/frames", "HTTP service path to send frames to")
}

// Validate checks validation of Config.
func (c *Config) Validate() error {
	if c.Teleinfo.Mode != "historic" && c.Teleinfo.Mode != "standard" {
		return fmt.Errorf("unknown teleinfo.mode: '%s' (supported: historic, standard)", c.Teleinfo.Mode)
	}
	return nil
}
