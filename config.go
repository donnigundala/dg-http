package dghttp

import "time"

// Config holds all configuration for an HTTP server.
type Config struct {
	Enabled      bool          `mapstructure:"enabled"`
	Addr         string        `mapstructure:"addr"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
	TLS          TLSConfig     `mapstructure:"tls"`
}

// TLSConfig holds the TLS-specific configuration.
type TLSConfig struct {
	Enabled    bool   `mapstructure:"enabled"`
	CertFile   string `mapstructure:"cert_file"`
	KeyFile    string `mapstructure:"key_file"`
	TLSVersion string `mapstructure:"tls_version"`
}
