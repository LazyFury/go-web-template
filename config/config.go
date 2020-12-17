package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// ReadConfig ReadConfig
func ReadConfig(data interface{}, path string) (err error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		log.Fatalln("打开配置文件错误，请创建!")
		return io.EOF
	}
	if err = json.NewDecoder(f).Decode(data); err != nil {
		return err
	}
	return nil
}
