package config

import (
	"github.com/ogier/pflag"
	"github.com/spf13/viper"

	"github.com/1ch0/gin-template/pkg/server/infrastructure/datastore"
)

// Config config for server
type Config struct {
	Server Server

	// Datastore config
	Datastore datastore.Config
}

type Server struct {
	// api server bind address
	BindAddr string
	// monitor metric path
	MetricPath string
	// PprofAddr the address for pprof to use while exporting profiling results.
	PprofAddr string
}

// ReadConfig config for server
func ReadConfig(path, name, configType string) *Config {
	config := &Config{}
	vip := viper.New()
	vip.AddConfigPath(path)
	vip.SetConfigName(name)
	vip.SetConfigType(configType)

	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := vip.Unmarshal(&config); err != nil {
		panic(err)
	}

	SetDefault(config)
	return config
}

func SetDefault(config *Config) {
	if config.Server.BindAddr == "" {
		config.Server.BindAddr = "0.0.0.0:8000"
	}
	if config.Server.MetricPath == "" {
		config.Server.MetricPath = "/metrics"
	}
	if config.Server.PprofAddr == "" {
		config.Server.PprofAddr = ""
	}

	if config.Datastore.Type == "" {
		config.Datastore.Type = "mongodb"
	}
	if config.Datastore.Database == "" {
		config.Datastore.Database = "go-restful-template"
	}
}

// AddFlags adds flags to the specified FlagSet
func (s *Config) AddFlags(fs *pflag.FlagSet, c *Config) {
	fs.StringVar(&s.Server.BindAddr, "bind-addr", c.Server.BindAddr, "The bind address used to serve the http APIs.")
	fs.StringVar(&s.Server.MetricPath, "metrics-path", c.Server.MetricPath, "The path to expose the metrics.")
	fs.StringVar(&s.Datastore.Database, "datastore-database", c.Datastore.Database, "Metadata storage database name, takes effect when the storage driver is mongodb.")
	fs.StringVar(&s.Datastore.URL, "datastore-url", c.Datastore.URL, "Metadata storage database url,takes effect when the storage driver is mongodb.")
	fs.StringVar(&s.Server.PprofAddr, "pprof-addr", c.Server.PprofAddr, "The address for pprof to use while exporting profiling results. The default value is empty which means do not expose it. Set it to address like :6666 to expose it.")
}
