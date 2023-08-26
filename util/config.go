package util

import (
	"time"

	"github.com/spf13/viper"
)

// config stores all configuration of the application
// the values are read by viper from a config file or environment file
type Config struct {
	DBDriver              string        `mapstructure:"DB_DRIVER"`
	DBSource              string        `mapstructure:"DB_SOURCE"`
	MigrationURL          string        `mapstructure:"MIGRATION_URL"`
	HTTPServerAddress     string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress     string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey     string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenExpiration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration  time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

// LoadConfig reads configurations from file or environment file
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") //json OR xml

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
