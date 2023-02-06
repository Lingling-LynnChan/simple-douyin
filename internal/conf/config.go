package conf

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Data      *Data      `yaml:"data"`      // 数据源相关配置
	Server    *Server    `yaml:"server"`    // 服务器相关配置
	Bootstrap *Bootstrap `yaml:"bootstrap"` // 启动器相关配置
}

type Bootstrap struct {
	AppName  string `yaml:"app_name"`  // App 名字
	AppOwner string `yaml:"app_owner"` // App 所有者
	Addr     string `yaml:"addr"`      // 开启端口
}

type Server struct {
	StaticDir string `yaml:"static_dir"` // 静态资源路径配置
	StaticUrl string `yaml:"static_url"` // 静态资源 url
	IsLogger  bool   `yaml:"is_logger"`  // 是否需要日志
	IsRecover bool   `yaml:"is_recover"` // 是否需要恢复程序
}

type Data struct {
	Redis    *Redis    `yaml:"redis"`
	Database *Database `yaml:"database"`
}

type Database struct {
	Driver          string `yaml:"driver"`
	Source          string `yaml:"source"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}

type Redis struct {
	Addr         string  `yaml:"addr"`
	Network      string  `yaml:"network"`
	ReadTimeout  float32 `yaml:"read_timeout"`
	WriteTimeout float32 `yaml:"write_timeout"`
}

func GetConf(path string) *Config {
	var config Config

	yamlFile, err := os.ReadFile(path)
	if err != nil {
		panic("reading config file failed: " + err.Error())
	}

	// err = yaml.UnmarshalStrict(yamlFile, bootstrap)
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic("unmarshaling config struct failed" + err.Error())
	}

	return &config
}
