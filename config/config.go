package config

import (
	"io/ioutil"

	"github.com/dowerx/goSchedule/ec"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Address     string
	Tasks       string
	TaskConfigs string `yaml:"taskconfigs"`
}

func LoadConfig(path string) Config {
	data, err := ioutil.ReadFile(path)
	ec.Check(err)
	conf := Config{}
	err = yaml.Unmarshal(data, &conf)
	ec.Check(err)
	return conf
}
