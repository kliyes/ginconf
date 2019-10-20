package config

import (
	"log"

	"github.com/go-ini/ini"
)

// Base defines basic configuration struct
type Base struct {
	ListenAddr string `ini:"listen_address"`
	Mode       string `ini:"mode"`
}

// Database defines database configuration struct
type Database struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Name     string `ini:"name"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Charset  string `ini:"charset"`
	Locale   string `ini:"locale"`
}

// Initialize parses configuration files into given config struct
func Initialize(config interface{}, files ...interface{}) {
	if err := ini.MapTo(config, defaultConfig, files...); err != nil {
		log.Fatalf("configuration file(s) parsed error: %v", err)
	}
	log.Println(config)
}
