package config

import "github.com/spf13/viper"

type Config struct {
	Application struct {
		Company     string
		Project     string
		ServiceName string
		LogLevel    string
		Env         string
		Host        string
		Port        int
	}
	Postgresql struct {
		Host     string
		Port     int
		User     string
		DbName   string
		Password string
	}
}

// LoadConfig 加载配置文件
func LoadConfig(file ...string) (*Config, error) {
	fi := "./config.yaml"
	if len(file) > 0 {
		fi = file[0]
	}
	viper.SetConfigFile(fi)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	conf := &Config{}
	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
