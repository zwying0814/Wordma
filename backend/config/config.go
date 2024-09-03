package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppKey          string
	NeedModeration  bool
	IPDataPath      string
	DatabasePath    string
	Port            string
	DevelopMode     bool
	DisableLog      bool
	LogPath         string
	NeedFilter      bool
	FilterPath      string
	Cors            string
	SupportMarkdown bool
	EmojiPaths      []string
)

// InitConfigFile 初始化
func InitConfigFile() {
	file, err := ini.Load("./data/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：", err.Error())
	}
	LoadSiteConfig(file)
	LoadCommentConfig(file)
	LoadSystemConfig(file)
}

func LoadSiteConfig(file *ini.File) {
	Port = file.Section("site").Key("Port").MustString("36580")
	AppKey = file.Section("site").Key("AppKey").MustString("F=G*[(m_1*m_2)/(r^2)]")
	IPDataPath = file.Section("site").Key("IPDataPath").MustString("/data/ip2region.xdb")
	DatabasePath = file.Section("site").Key("DatabasePath").MustString("/data/database.bin")
	Cors = file.Section("site").Key("Cors").MustString("*")
}

func LoadCommentConfig(file *ini.File) {
	NeedModeration = file.Section("comment").Key("NeedModeration").MustBool(false)
	NeedFilter = file.Section("comment").Key("NeedFilter").MustBool(true)
	FilterPath = file.Section("comment").Key("FilterPath").MustString("/data/filter.txt")
	SupportMarkdown = file.Section("comment").Key("SupportMarkdown").MustBool(true)
	EmojiPaths = file.Section("comment").Key("EmojiPaths").Strings("|")
}

func LoadSystemConfig(file *ini.File) {
	DevelopMode = file.Section("system").Key("DevelopMode").MustBool(false)
	DisableLog = file.Section("system").Key("DisableLog").MustBool(false)
	LogPath = file.Section("system").Key("LogPath").MustString("/data/log.txt")
}
