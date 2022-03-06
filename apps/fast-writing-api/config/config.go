package config

import (
	"errors"
	"github.com/spf13/viper"
	"strconv"
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

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.New("cannot read from " + path + ": " + err.Error())
	}
	return &Config{
		Database: Database{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetInt("database.port"),
			User:     viper.GetString("database.user"),
			Password: "system",
			Schema:   viper.GetString("database.schema"),
			Timeout:  viper.GetInt("database.timeout"),
			Debug:    viper.GetBool("database.debug"),
		}}, nil
}

func (c *Config) GetSQLConnection() string {
	return c.Database.User + ":" + c.Database.Password + "@tcp(" + c.Database.Host + ":" + strconv.Itoa(c.Database.Port) + ")/" + c.Database.Schema + "?timeout=" + strconv.Itoa(c.Database.Timeout) + "s" + "&parseTime=true"
}
