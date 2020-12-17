package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type (
	// BaseConfig BaseConfig
	BaseConfig struct {
		TablePrefix string `json:"table_prefix"` //数据库表前缀
		BaseURL     string `json:"base_url"`     // 网站根目录
		Port        int    `json:"port"`         //端口

		IV  string `json:"iv"`
		KEY string `json:"key"`
	}
)

// ReadConfig ReadConfig
func ReadConfig(data interface{}, path string) interface{} {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		log.Fatalln("打开配置文件错误，请创建!")
		panic(io.EOF)
	}
	if err = json.NewDecoder(f).Decode(data); err != nil {
		panic(err)
	}
	return data
}
