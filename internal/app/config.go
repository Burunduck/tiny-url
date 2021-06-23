package config

import (
	"github.com/joho/godotenv"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	ServiceHost     string
	ServiceGrpcPort int

	LogLevel string
}

func NewConfig() (*Config, error) {
	configName := "config"
	_ = godotenv.Load()
	if os.Getenv("CONFIG_NAME") != "" {
		configName = os.Getenv("CONFIG_NAME")
	}

	viper.SetConfigName(configName)
	viper.SetConfigType("toml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
	viper.WatchConfig()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	cfg.setLogLevel(cfg.LogLevel)

	jww.INFO.Println("Config parsed")
	return cfg, nil
}

func (c *Config) setLogLevel(logLevel string) {
	foundLogLevel, ok := logLevelMap[logLevel]
	if !ok {
		jww.ERROR.Printf("incorrect log level %s\n", logLevel)
		return
	}
	jww.SetLogThreshold(foundLogLevel)
	jww.SetStdoutThreshold(foundLogLevel)
}

var logLevelMap = map[string]jww.Threshold{
	"DEBUG": jww.LevelDebug,
	"INFO":  jww.LevelInfo,
	"WARN":  jww.LevelWarn,
	"ERROR": jww.LevelError,
	"FATAL": jww.LevelFatal,
}
