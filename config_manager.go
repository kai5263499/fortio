package fortio

import (
	"fmt"

	"unicode"

	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const tagName = "config"

// ConfigManager auto wires the given config pointer with given set of
// config loaders to NewConfigManager API
type ConfigManager struct {
	rootCmd       *cobra.Command
	logger        Logger
	configLoaders []ConfigLoader
}

// NewConfigManager returns new instance of ConfigManager
func NewConfigManager(appName, description string, configLoaders ...ConfigLoader) *ConfigManager {
	rootCmd := &cobra.Command{
		Use:   appName,
		Short: description,
	}

	helpCmd := &cobra.Command{
		Use:   "help",
		Short: fmt.Sprintf("help for %s", appName),
		Run: func(cmd *cobra.Command, args []string) {
			rootCmd.Flags().BoolP("help", "h", true, fmt.Sprintf("help for %s", appName))
			cmd.Usage()
		},
	}

	rootCmd.SetHelpCommand(helpCmd)
	rootCmd.AddCommand(helpCmd)

	return &ConfigManager{
		logger:  NewStdLogger(3, log.Ldate|log.Ltime),
		rootCmd: rootCmd,
		configLoaders: configLoaders,
	}
}

// SetLogger will set given logger and uses it for logging
func (cm *ConfigManager) SetLogger(logger Logger) {
	cm.logger = logger
}

// Load will create command line flags for given config and loads values into
// it from environment variables
func (cm *ConfigManager) Load(config Config) error {
	err := cm.createCommandLineFlags(config)
	if err != nil {
		cm.logger.Errorf("Unable to load config - %v", err)
		return err
	}

	cm.rootCmd.Run = func(cmd *cobra.Command, args []string) {}

	if err := cm.rootCmd.Execute(); err != nil {
		cm.logger.Debugf("Command line args: %+v", os.Args)
		cm.logger.Errorf("Error executing rootCmd - %v", err)
		return err
	}

	for _, loader := range cm.configLoaders {
		err := loader.Load(config)
		if err != nil {
			return err
		}
	}

	helpIsSet, _ := cm.rootCmd.Flags().GetBool("help")
	if helpIsSet {
		os.Exit(0)
	}

	return nil
}

// createCommandLineFlags will create command line flags for given config via Cobra and Viper
// to support command line overriding of config values
func (cm *ConfigManager) createCommandLineFlags(config interface{}) error {
	fields := make(map[string]field)
	getAllFields(config, fields)
	for name, field := range fields {
		lFirst := lowerFirst(name)

		switch ptr := field.addr.(type) {
		case *string:
			viper.SetDefault(lFirst, field.defaultValue)
			cm.rootCmd.PersistentFlags().String(lFirst, *ptr, field.usage)
		case *int:
			val, err := strconv.Atoi(field.defaultValue)
			if err != nil {
				cm.logger.Fatalf("default specified for %s is not a int", field.name)
			}
			viper.SetDefault(lFirst, val)
			cm.rootCmd.PersistentFlags().Int(lFirst, *ptr, field.usage)
		case *int64:
			val, err := strconv.ParseInt(field.defaultValue, 10, 64)
			if err != nil {
				cm.logger.Fatalf("default specified for %s is not a int64 val=%v defaultValue=%v", field.name, val, field.defaultValue)
			}
			viper.SetDefault(lFirst, val)
			cm.rootCmd.PersistentFlags().Int64(lFirst, *ptr, field.usage)
		case *uint:
			val, err := strconv.Atoi(field.defaultValue)
			if err != nil {
				cm.logger.Fatalf("default specified for %s is not a uint", field.name)
			}
			viper.SetDefault(lFirst, val)
			cm.rootCmd.PersistentFlags().Uint(lFirst, *ptr, field.usage)
		case *uint64:
			val, err := strconv.ParseUint(field.defaultValue, 10, 64)
			if err != nil {
				cm.logger.Fatalf("default specified for %s is not a uint64", field.name)
			}
			viper.SetDefault(lFirst, val)
			cm.rootCmd.PersistentFlags().Uint64(lFirst, *ptr, field.usage)
		case *float64:
			val, err := strconv.ParseFloat(field.defaultValue, 64)
			if err != nil {
				cm.logger.Fatalf("default specified for %s is not a float64", field.name)
			}
			viper.SetDefault(lFirst, val)
			cm.rootCmd.PersistentFlags().Float64(lFirst, *ptr, field.usage)
		case *bool:
			val, err := strconv.ParseBool(field.defaultValue)
			if err != nil {
				cm.logger.Fatalf("default specified for %s is not a bool", field.name)
			}
			viper.SetDefault(lFirst, val)
			cm.rootCmd.PersistentFlags().Bool(lFirst, *ptr, field.usage)
		default:
			cm.logger.Warnf("unknown field %s type %v", field.name, reflect.TypeOf(field))
			// TODO: support for types implementing pflag.Value
		}

		underscoreField := camelCaseToUnderscore(name)
		if field.env != "" {
			viper.BindEnv(lFirst, field.env)
		} else {
			viper.BindEnv(lFirst, underscoreField)
		}
		viper.BindPFlag(lFirst, cm.rootCmd.PersistentFlags().Lookup(lFirst))
	}
	return nil
}

type field struct {
	addr         interface{}
	name         string
	defaultValue string
	usage        string
	env          string
	required     bool
}

func lowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

func getAllFields(obj interface{}, m map[string]field) {
	xv := reflect.ValueOf(obj).Elem() // Dereference into addressable value
	xt := xv.Type()

	for i := 0; i < xt.NumField(); i++ {
		f := xt.Field(i)
		if f.Anonymous {
			getAllFields(xv.Field(i).Addr().Interface(), m)
		} else {
			addr := xv.Field(i).Addr().Interface()

			fld := getField(xt.Field(i))
			fld.name = f.Name
			fld.addr = addr
			// add defaults to help, cobra/viper doesn't let us add this
			// and no clear example on how to use SetHelpTemplate
			fld.usage = fmt.Sprintf("%s [default: %v]", fld.usage, fld.defaultValue)
			m[f.Name] = fld
		}
	}
}

func getField(fld reflect.StructField) field {
	f := field{
		required: false,
	}
	tags := strings.Split(fld.Tag.Get(tagName), ";")
	for _, tag := range tags {
		t := strings.Split(tag, "=")
		if t[0] == "default" {
			f.defaultValue = t[1]
		} else if t[0] == "usage" {
			f.usage = t[1]
		} else if t[0] == "env" {
			f.env = t[1]
		} else if t[0] == "required" {
			f.required = true
		}

	}
	return f
}

func camelCaseToUnderscore(s string) string {
	runes := []rune(s)
	size := len(runes)

	var out []rune
	for i := 0; i < size; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < size && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToUpper(runes[i]))
	}

	return string(out)
}
