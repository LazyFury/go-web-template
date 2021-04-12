package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type (
	// BaseConfig BaseConfig
	BaseConfig struct {
		TablePrefix string `json:"table_prefix"` //数据库表前缀
		BaseURL     string `json:"base_url"`     // 网站根目录
		Port        int    `json:"port"`         //端口
	}
)

// writeConf writeConf
func writeConf(data interface{}, path string) error {
	byte, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("配置文件类型错误")
	}

	err = ioutil.WriteFile(path, byte, 0644)
	if err != nil {
		log.Fatalf("写入默认配置失败")
	}
	log.Printf("已生成默认配置文件")
	return err
}

// ReadConfig ReadConfig
func ReadConfig(data interface{}, path string) interface{} {
	f, err := os.Open(path)

	if err != nil {
		writeConf(data, path)
		log.Fatalln("打开配置文件错误，请创建!")
	}
	defer f.Close()

	if err = json.NewDecoder(f).Decode(data); err != nil {
		panic(err)
	}
	return data
}
