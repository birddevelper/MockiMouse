package config

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type configResolver struct {
	config Config
}

var ConfigResolver configResolver

func init() {
	ConfigResolver.loadConfig()
}
func (resolver configResolver) loadConfig() {
	filename, _ := filepath.Abs("./config.yml")
	configFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configFile, &ConfigResolver.config)
	if err != nil {
		panic(err)
	}
}

func (resolver configResolver) GetPort() int {

	return resolver.config.MockServer.Port
}

func (resolver configResolver) GetContextPath() string {

	return resolver.config.MockServer.ContextPath
}

func (resolver configResolver) GetEndPoints() []EndPoint {

	return resolver.config.MockServer.Endpoints
}
