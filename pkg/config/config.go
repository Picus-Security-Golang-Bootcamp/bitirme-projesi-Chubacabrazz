package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

// Config
type Config struct {
	ServerConfig ServerConfig `yaml:"ServerConfig"`
	JWTConfig    JWTConfig    `yaml:"JWTConfig"`
	DBConfig     DBConfig     `yaml:"DBConfig"`
	Logger       Logger       `yaml:"Logger"`
	CartConfig   CartConfig   `yaml:"CartConfig"`
}

// ServerConfig
type ServerConfig struct {
	AppVersion       string `yaml:"AppVersion"`
	Mode             string `yaml:"Mode"`
	RoutePrefix      string `yaml:"RoutePrefix"`
	Debug            bool   `yaml:"Debug"`
	Port             int    `yaml:"Port"`
	TimeoutSecs      int    `yaml:"TimeoutSecs"`
	ReadTimeoutSecs  int    `yaml:"ReadTimeoutSecs"`
	WriteTimeoutSecs int    `yaml:"WriteTimeoutSecs"`
}

// JWTConfig
type JWTConfig struct {
	SessionTime int    `yaml:"SessionTime"`
	SecretKey   string `yaml:"SecretKey"`
}

// DBConfig
type DBConfig struct {
	MigrationFolder string `yaml:"MigrationFolder"`
	DataSourceName  string `yaml:"DataSourceName"`
	Name            string `yaml:"Name"`
	MaxOpen         int    `yaml:"MaxOpen"`
	MaxIdle         int    `yaml:"MaxIdle"`
	MaxLifetime     int    `yaml:"MaxLifetime"`
}

// Logger
type Logger struct {
	Development bool   `yaml:"Development"`
	Encoding    string `yaml:"Encoding"`
	Level       string `yaml:"Level"`
}

// CartConfig
type CartConfig struct {
	MaxAllowedForBasket     int     `yaml:"MaxAllowedForBasket"`
	MaxAllowedQtyPerProduct int     `yaml:"MaxAllowedQtyPerProduct"`
	MinCartAmountForOrder   float64 `yaml:"MinCartAmountForOrder"`
}

// LoadConfig: Load config file from given path
func LoadConfig(filename string) (*Config, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil

}
