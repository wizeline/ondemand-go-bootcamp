package configuration

import (
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/logger"
)

var (
	cfgInstance = newConfig()
	once        sync.Once
)

type Config struct {
	AppVersion SemanticVersion
	HTTP       HTTP
	Database   Database
}

func GetInstance() Config {
	return cfgInstance
}

func setDefaultConfig() {
	viper.SetDefault("http.address", ":8080")
	viper.SetDefault("http.shutdown.timeout", time.Second*15)
	viper.SetDefault("application.version", "v0.0.0")
	viper.SetDefault("database.driver", "csv")
	viper.SetDefault("database.csv.file_name", "fruits.csv")
	viper.SetDefault("database.csv.data_dir", "./data")
}

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
				address:         viper.GetString("http.address"),
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
