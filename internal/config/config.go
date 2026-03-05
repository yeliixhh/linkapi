package config

import (
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
)

type Config struct {
	ServerConf *ServerConfig `json:"server" yaml:"server"`
	DBConf     *DBConfig     `json:"db" yaml:"db"`
}

type ServerConfig struct {
	Port string `json:"port" yaml:"port"`
	Host string `json:"host" yaml:"host"`
}

type DBConfig struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	DBName   string `json:"database" yaml:"dbname"`
}

// 创建配置文件
func NewConfig() (*Config, error) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	var conf Config

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.New("配置文件加载失败")
	}

	fileContent, err := os.ReadFile(viper.ConfigFileUsed())
	if err != nil {
		return nil, errors.New("配置文件加载失败")
	}
	compile := regexp.MustCompile(`\${([^}]+)}`)

	replaceContent := compile.ReplaceAllStringFunc(string(fileContent), func(match string) string {
		envVar := match[2 : len(match)-1]
		if getenv := os.Getenv(envVar); getenv != "" {
			return getenv
		}
		return match
	})

	err = viper.ReadConfig(strings.NewReader(replaceContent))
	if err != nil {
		return nil, errors.New("配置文件加载失败")
	}

	if err := viper.Unmarshal(&conf, func(dc *mapstructure.DecoderConfig) {
		dc.TagName = "yaml"
	}); err != nil {
		return nil, errors.New("配置文件加载失败")
	}

	return &conf, nil
}
