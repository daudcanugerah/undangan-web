package config

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"braces.dev/errtrace"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// Sqlite ...
type Sqlite struct {
	DBFile string `mapstructure:"db_file"`
}

// RedisConfig ...
type RedisConfig struct {
	DSN      string `mapstructure:"dsn"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// AppConfig ...
type AppConfig struct {
	Timezone string `mapstructure:"timezone"`
	Proxy    string `mapstructure:"proxy_url"`
}

// NatsConfig ...
type NatsConfig struct {
	DSN string `mapstructure:"dsn"`
}

// Config ...
type Config struct {
	Otel     OtelConfig `mapstructure:"otel"`
	DBSqlite Sqlite     `mapstructure:"sqlite"`
	Nats     NatsConfig `mapstructure:"nats"`
	App      AppConfig  `mapstructure:"app"`
}

// SetUpTimezone ...
func SetUpTimezone(tz string) error {
	if tz != "" {
		var err error
		time.Local, err = time.LoadLocation(tz)
		if err != nil {
			return errtrace.Wrap(fmt.Errorf("error loading location '%s': %v", tz, err))
		}
	}

	return nil
}

func SetProxy(proxy string) (*http.Client, error) {
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return nil, fmt.Errorf("invalid proxy URL: %v", err)
	}

	httpClient := &http.Client{
		Timeout: time.Second * 60,
		Transport: &http.Transport{
			Proxy:           http.ProxyURL(proxyURL),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	return httpClient, nil
}

// InitConfig ...
func InitConfig(cfgFile string) (*Config, error) {
	var config Config

	viper.SetConfigType("toml")
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			return nil, fmt.Errorf("find home dir error: %v", err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, errtrace.Wrap(fmt.Errorf("read init config error: %v", err))
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
	// parse config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, errtrace.Wrap(fmt.Errorf("unmarshal config error: %v", err))
	}

	return &config, nil
}
