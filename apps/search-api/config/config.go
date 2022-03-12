package config

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	Elasticsearch Elasticsearch
}

type Elasticsearch struct {
	Host     string
	Port     int
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
		Elasticsearch: Elasticsearch{
			Host:     viper.GetString("elasticsearch.host"),
			Port:     viper.GetInt("elasticsearch.port"),
		}}, nil
}
