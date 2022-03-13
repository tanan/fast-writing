package config

import (
	"errors"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

type Config struct {
	Database Database
}

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Schema   string
	Timeout  int
	Debug    bool
}

const (
	FileName = "config"
	FileType = "yaml"
)

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName(FileName)
	viper.SetConfigType(FileType)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.New("cannot read from " + path + ": " + err.Error())
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.New("cannot unmarshal config: " + err.Error())
	}
	return &Config{
		Database: Database{
			Host:     cfg.Database.Host,
			Port:     cfg.Database.Port,
			User:     cfg.Database.User,
			Password: cfg.Database.Password,
			Schema:   cfg.Database.Schema,
			Timeout:  cfg.Database.Timeout,
			Debug:    cfg.Database.Debug,
		}}, nil
}

func (c *Config) GetSQLConnection() string {
	return c.Database.User + ":" + c.Database.Password + "@tcp(" + c.Database.Host + ":" + strconv.Itoa(c.Database.Port) + ")/" + c.Database.Schema + "?timeout=" + strconv.Itoa(c.Database.Timeout) + "s" + "&parseTime=true"
}
