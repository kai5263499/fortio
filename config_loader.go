package fortio

// ConfigLoader defines a interface that needs to be implemented by
// a config loader to be able to plug into config manager
type ConfigLoader interface {
	// Load takes in a implementation of Config and populates field values
	Load(config Config) error
}
