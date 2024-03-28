package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
)

var (
	HttpPort string
	DB       *gorm.DB
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检测文件路径：", err)
	}
	LoadServer(file)

}

func LoadServer(file *ini.File) {
	HttpPort = file.Section("server").Key("HttpPort").MustString("3000")
}
