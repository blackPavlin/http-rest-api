package store

// Config ...
type Config struct {
	DarabaseURL string `toml:"database_url"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config {

	}
}