package config

import (
	"github.com/go-ini/ini"
	"log"
	"strings"
)

var (
	dbUser      string
	dbPassword  string
	dbTableName string

	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func Init() {
	IniFile, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Println("err = ", err)
		return
	}
	getMysqlValue(IniFile)
	getRedisValue(IniFile)
	dsn := strings.Join([]string{dbUser, ":", dbPassword, "@tcp(localhost:3306)/", dbTableName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	InitDB(dsn)
}

func getMysqlValue(IniFile *ini.File) {
	dbUser = IniFile.Section("database").Key("dbUser").String()
	dbPassword = IniFile.Section("database").Key("dbPassword").String()
	dbTableName = IniFile.Section("database").Key("dbTableName").String()
}

func getRedisValue(file *ini.File) {
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}
