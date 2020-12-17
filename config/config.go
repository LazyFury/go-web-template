package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
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
