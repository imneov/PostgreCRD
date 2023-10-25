package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"log"
)

func Init() *Config {
	// viper.SetConfigName("config1") // 读取yaml配置文件
	viper.SetConfigName("config") // 读取json配置文件
	//viper.AddConfigPath("/etc/appname/")   //设置配置文件的搜索目录
	//viper.AddConfigPath("$HOME/.appname")  // 设置配置文件的搜索目录
	viper.AddConfigPath(".") // 设置配置文件和可执行二进制文件在用一个目录
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
		log.Fatal(err) // 读取配置文件失败致命错误
	}

	var c Config = Config{
		Databases: []Database{
			{
				PGServer: PGServer{DATABASE: "aaaaa"},
				TableMapping: map[string]TableMapping{
					"iam": TableMapping{
						Type: "",
						Rules: map[string]Rule{
							"name": {Source: "metadata.name"},
						},
					},
				},
			},
		},
	}
	viper.Unmarshal(&c)
	byt, _ := yaml.Marshal(c)
	fmt.Println(string(byt))
	return &c
}
