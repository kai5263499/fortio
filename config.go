package fortio

// Config defines a interface that needs to be implemented
// by all config struct's to be used with ConfigManager
type Config interface {
	// Validates assigned config values
	Validate() error

	// DumpJSON will return JSON marshalled string of config
	DumpJSON() (string, error)

	// DumpYAML will return YAML marshalled string of config
	DumpYAML() (string, error)
}
