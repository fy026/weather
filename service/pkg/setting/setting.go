package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string
	ENV     string

	Server_Name  string
	HTTPHost     string
	HTTPPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	RegisterUrl string

	EndServerName string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadRegister()

}

func LoadBase() {
	RunMode = Cfg.Section("").Key("ENV").MustString("")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	Server_Name = sec.Key("SERVER_NAME").MustString("ts")
	HTTPHost = sec.Key("HTTP_HOST").MustString("")
	HTTPPort = sec.Key("HTTP_PORT").MustString("50001")
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

}

func LoadRegister() {
	sec, err := Cfg.GetSection("register")
	if err != nil {
		log.Fatalf("Fail to get section 'register': %v", err)
	}

	RegisterUrl = sec.Key("REGISTER_RUL").MustString("http://127.0.0.1:2379")

}
