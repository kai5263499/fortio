package fortio

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

// ConfigLoader defines a interface that needs to be implemented by
// a config loader to be able to plug into config manager
type ConfigLoader interface {
	// Load takes in a implementation of Config and populates field values
	Load(config Config) error
}

// CmdLineConfigLoader is config loader that makes the given config fields
// available as a command line flags for overwriting.
type CmdLineConfigLoader struct {
}

// Load will load the config field values as command line flags
func (cmd *CmdLineConfigLoader) Load(config Config) error {
	cfg := reflect.ValueOf(config)
	return cmd.loadValue(cfg, "")
}

func (cmd *CmdLineConfigLoader) loadValue(dest reflect.Value, name string) error {
	switch dest.Elem().Type().Kind() {
	case reflect.Struct:
		for i := 0; i < dest.Elem().Type().NumField(); i++ {
			fieldStruct := dest.Elem().Type().Field(i)
			cmd.loadValue(dest.Elem().Field(i).Addr(), fieldStruct.Name)
		}
	case reflect.String:
		val := viper.GetString(name)
		dest.Elem().SetString(val)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val := viper.GetInt64(name)
		dest.Elem().SetInt(val)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val := viper.GetInt64(name)
		dest.Elem().SetInt(val)
	case reflect.Float32, reflect.Float64:
		val := viper.GetFloat64(name)
		dest.Elem().SetFloat(val)
	case reflect.Bool:
		val := viper.GetBool(name)
		dest.Elem().SetBool(val)
	default:
		// TODO: support for map/slice types
		msg := fmt.Sprintf("unsupported type: %s", dest.Elem().Type())
		return errors.New(msg)
	}
	return nil
}
