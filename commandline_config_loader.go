package fortio

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// CmdLineConfigLoader is config loader that makes the given config fields
// available as a command line flags for overwriting.
type CmdLineConfigLoader struct {
	types []interface{}
}

// NewCmdLineConfigLoader will return new CmdLineConfigLoader with capability to read given types
func NewCmdLineConfigLoader(types ...interface{}) *CmdLineConfigLoader {
	types = append(types, &Duration{})
	return &CmdLineConfigLoader{
		types: types,
	}
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
		parsed := false
		if sp, ok := dest.Interface().(StringParsable); ok {
			val := viper.GetString(name)
			sp.ParseString(val)
			parsed = true
		}
		if len(cmd.types) > 0 {
			for _, t := range cmd.types {
				if dest.Elem().Type().String() == reflect.TypeOf(t).Elem().String() {
					val := viper.GetString(name)
					iface, ok := t.(pflag.Value)
					if !ok {
						msg := fmt.Sprintf("%s is not of type %T", reflect.TypeOf(t).Elem().String(), iface)
						return errors.New(msg)
					}
					err := iface.Set(val)
					if err != nil {
						return err
					}

					dest.Elem().Set(reflect.ValueOf(iface).Elem())
					parsed = true
				}
			}
		}
		if !parsed {
			for i := 0; i < dest.Elem().Type().NumField(); i++ {
				fieldStruct := dest.Elem().Type().Field(i)
				cmd.loadValue(dest.Elem().Field(i).Addr(), fieldStruct.Name)
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
		// TODO: support for map/slice types
		msg := fmt.Sprintf("unsupported type: %s", dest.Elem().Type())
		return errors.New(msg)
	}
	return nil
}
