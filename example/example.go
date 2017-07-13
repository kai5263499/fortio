package main

import (
	"encoding/json"
	"errors"

	"github.com/CrowdStrike/fortio"
	"gopkg.in/yaml.v2"
)

type ExampleConfig struct {
	Timeout int64  `config:"env=TIMEOUT;default=100000;usage=Timeout for service" json:"timeout"`
	Name    string `config:"default=;usage=Name of service" json:"name"`
}

// Validates assigned config values
func (ec *ExampleConfig) Validate() error {
	if ec.Timeout > 100000 {
		return errors.New("Timeout can't be greater than 100ms")
	}
	if ec.Name == "" {
		return errors.New("Name can't be empty")
	}
	return nil
}

// DumpJSON will return JSON marshalled string of config
func (ec *ExampleConfig) DumpJSON() (string, error) {
	v, err := json.Marshal(ec)
	return string(v), err
}

// DumpYAML will return YAML marshalled string of config
func (ec *ExampleConfig) DumpYAML() (string, error) {
	v, err := yaml.Marshal(ec)
	return string(v), err
}

func main() {
	// Initialize empty config as pointer
	config := &ExampleConfig{}
	// Initialize config manager
	cm := fortio.NewConfigManager("fortio-test", "My Fortio example")
	// Pass config pointer to be loaded from env variables
	err := cm.Load(config)
	if err != nil {
		// handle error
	}

	// validate configs loaded
	err = config.Validate()
	if err != nil {
		// handle error
	}
}
