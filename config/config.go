package config

import (
	"fmt"
)

var configuration = map[string]string{
	"mysql_host":     "localhost",
	"mysql_port":     "3306",
	"mysql_db":       "userdb",
	"mysql_user":     "root",
	"mysql_password": "main",
	"ACCESS_SECRET":  "232791",
	"redis_host":     "localhost",
	"redis_port":     "6379",
	"redis_password": "",
}

func LoadConfig() map[string]string {
	return configuration
}

func GetValue(key string) (string, error) {
	if v, ok := configuration[key]; ok {
		return v, nil
	}
	return "", fmt.Errorf("configuration variable %s not found", key)
}
