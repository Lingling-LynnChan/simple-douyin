package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Bootstrap struct {
	Data   *Data   `yaml:"data"`
	Server *Server `yaml:"server"`
}

type Server struct {
	Addr     string `yaml:"addr"`
	AppName  string `yaml:"appName"`
	AppOwner string `yaml:"appOwner"`
}

type Data struct {
	Redis    *Redis    `yaml:"redis"`
	Database *Database `yaml:"database"`
}

type Database struct {
	Driver          string `yaml:"driver"`
	Source          string `yaml:"source"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	ConnMaxLifetime int    `yaml:"connMaxLifetime"`
}

type Redis struct {
	Addr         string  `yaml:"addr"`
	Network      string  `yaml:"network"`
	ReadTimeout  float32 `yaml:"readTimeout"`
	WriteTimeout float32 `yaml:"writeTimeout"`
}

func GetConf(path string) *Bootstrap {
	var bootstrap Bootstrap

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		panic("reading config file failed: " + err.Error())
	}

	// err = yaml.UnmarshalStrict(yamlFile, bootstrap)
	err = yaml.Unmarshal(yamlFile, &bootstrap)

	if err != nil {
		panic("unmarshaling config struct failed" + err.Error())
	}

	return &bootstrap
}
