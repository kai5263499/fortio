package fortio

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

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
		// Don't enumerate fields if interface implements StringParsable
		if sp, ok := dest.Interface().(StringParsable); ok {
			val := viper.GetString(name)
			sp.ParseString(val)
		} else {
			for i := 0; i < dest.Elem().Type().NumField(); i++ {
				fieldStruct := dest.Elem().Type().Field(i)
				if err := cmd.loadValue(dest.Elem().Field(i).Addr(), lowerFirst(fieldStruct.Name)); err != nil {
					return err
				}
			}
		}
	case reflect.String:
		val := viper.GetString(name)
		if val != "" {
			dest.Elem().SetString(val)
		}
	case reflect.Slice:
		val := viper.GetString(name)

		sl := StringList{}
		sl.Set(val)

		dest.Elem().Set(reflect.ValueOf(sl))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val := viper.GetInt64(name)
		dest.Elem().SetInt(val)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val := uint64(viper.GetInt64(name))
		dest.Elem().SetUint(val)
	case reflect.Float32, reflect.Float64:
		val := viper.GetFloat64(name)
		dest.Elem().SetFloat(val)
	case reflect.Bool:
		val := viper.GetBool(name)
		dest.Elem().SetBool(val)
	case reflect.Interface:
		// Skip interface hints
	default:
		msg := fmt.Sprintf("unsupported type: %s", dest.Elem().Type())
		return errors.New(msg)
	}
	return nil
}
