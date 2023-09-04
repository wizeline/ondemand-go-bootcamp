package configuration

import (
	"strings"
	"sync"
	"time"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/logger"

	"github.com/spf13/viper"
)

var (
	cfgInstance = newConfig()
	once        sync.Once
)

// Config is the representation of the application's configuration.
type Config struct {
	AppVersion SemanticVersion
	HTTP       HTTP
	Database   Database
}

// GetInstance returns instance of the implemented configuration.
func GetInstance() Config {
	return cfgInstance
}

func setDefaultConfig() {
	viper.SetDefault("http.host", "0.0.0.0")
	viper.SetDefault("http.port", 8080)
	viper.SetDefault("http.shutdown.timeout", time.Second*15)
	viper.SetDefault("application.version", "v0.0.0")
	viper.SetDefault("database.driver", "csv")
	viper.SetDefault("database.csv.file_name", "fruits.csv")
	viper.SetDefault("database.csv.data_dir", "./data")
}

// newConfig creates a new Config instance of type singleton.
// The configuration instance can be configured only once.
func newConfig() Config {
	var cfg Config
	once.Do(func() {
		setDefaultConfig()
		viper.SetEnvPrefix("capstone")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		cfg = Config{
			AppVersion: SemanticVersion(viper.GetString("application.version")),
			HTTP: HTTP{
				host:            viper.GetString("http.host"),
				port:            viper.GetInt("http.port"),
				shutdownTimeout: viper.GetDuration("http.shutdown.timeout"),
			},
			Database: Database{
				driver: viper.GetString("database.driver"),
				CSV: NewCsvDB(
					viper.GetString("database.csv.file_name"),
					viper.GetString("database.csv.data_dir"),
				),
			},
		}
		logger.Log().Debug().
			Str("app_version", cfg.AppVersion.String()).
			Msg("configuration loaded")
	})
	return cfg
}
