package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/CrowdStrike/fortio"
	"gopkg.in/yaml.v2"
)

type ExampleConfig struct {
	Timeout  fortio.Duration  `config:"env=TIMEOUT;default=100ms;usage=Timeout for service" json:"timeout"`
	Name     string           `config:"default=test;usage=Name of service" json:"name"`
	Registry Registry         `config:";default=./registry.json;usage="`
	Map      fortio.MapObject `config:";default={\"a\":\"b\"};usage="`
}

// Validates assigned config values
func (ec *ExampleConfig) Validate() error {
	if ec.Timeout.Duration > time.Duration(100)*time.Millisecond {
		return errors.New("timeout can't be greater than 100ms")
	}
	if ec.Name == "" {
		return errors.New("name can't be empty")
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
		panic(err)
	}

	fmt.Printf("Loaded config: %+v\n", config)

	// validate configs loaded
	err = config.Validate()
	if err != nil {
		// handle error
	}
}

type Registry struct {
	RegType string
	Name    string
	Age     int
}

func (r *Registry) String() string {
	return "fortio.registry"
}

func (r *Registry) Set(s string) error {
	reg, err := LoadRegistryFromFile(s)
	if err != nil {
		return err
	}
	*r = *reg
	return nil
}

func (r *Registry) ParseString(s string) error {
	return r.Set(s)
}

func (r *Registry) Type() string {
	return "fortio.registry"
}

func LoadRegistryFromFile(path string) (*Registry, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := &Registry{}
	err = json.Unmarshal(raw, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
