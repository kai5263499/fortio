# Fortio 

## What is Fortio?
Fortio is a config auto wiring module at CrowdStrike. Fortio means cargo in Greek, configs are our cargo that needs to 
be loaded and carried along in our services. Fortio is used to auto wire configurations from multiple sources like 
environment variables, config files, HTTP endpoints etc and follows the precedence order as 
[viper](https://github.com/spf13/viper) into predefined defined config structs at startup. Fortio makes use of 
[cobra](https://github.com/spf13/cobra) and [viper](https://github.com/spf13/viper) to achieve some of these features 
and not reinvent the wheel, but Fortio is also designed in such a way that replacing cobra or viper to a different 
solution should be easier and seamless.

## Using Fortio
Get Fortio
```bash
go get github.com/CrowdStrike/fortio
```

Define your config that must be auto wired
```go
type ExampleConfig struct {
	Timeout int64  `config:"env=TIMEOUT,default=100000;usage=Timeout for service" json:"timeout"`
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
```

In your main file
```go
import "github.com/CrowdStrike/fortio"

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
```

Checkout above example from [example.go](https://github.com/CrowdStrike/fortio/example/example.go)

## Contributors

[Praveen Bathala](https://github.com/prvn)

[Wes Widner](https://github.com/kai5263499)

## Contribute

TODO: