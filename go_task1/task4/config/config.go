package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

// Http web配置结构体
type Http struct {
	Port int `yaml:"port"`
	// Domain   string `yaml:"domain"`
	// Protocol string `yaml:"protocol"`
}

// MySQL 配置结构体
type MySQL struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"db_name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	// MaxIdleConns    int    `yaml:"max_idle_conns"`
	// MaxOpenConns    int    `yaml:"max_open_conns"`
	// ConnMaxLefeTime int    `yaml:"conn_max_lefe_time"`
	Charset string `yaml:"charset"`
	//Collation string `yaml:"collation"`
}

// LogConf 日志相关配置
type LogConf struct {
	Level string `yaml:"level"`
}

type JWTConf struct {
	Secret string `yaml:"secret"`
	Expire int    `yaml:"expire"`
}

type Config struct {
	Http  Http    `yaml:"http"`
	MySQL MySQL   `yaml:"mysql"`
	Log   LogConf `yaml:"log"`
	JWT   JWTConf `yaml:"jwt"`
}

var conf *Config

func GetConf() *Config {
	fmt.Printf("conf port is %v \n", conf.Http)
	return conf
}

// InitConfig 初始化配置
func InitConfig() error {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config/config.yaml"
	}
	fmt.Println("configPath is : ", configPath)
	err := getYamlConf(configPath, &conf)
	return err
}

// getYamlConf 解析yaml配置
func getYamlConf(filePath string, out interface{}) (err error) {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	return yaml.Unmarshal(yamlFile, out)
}
