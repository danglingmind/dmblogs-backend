package config

import "errors"

var configuration = map[string]string{
	"MysqlHost":     "localhost",
	"MysqlPort":     "3306",
	"MysqlDbName":   "userdb",
	"MysqlUser":     "root",
	"Mysqlpassword": "main",
}

func LoadConfig() map[string]string {
	return configuration
}

func GetValue(key string) (string, error) {
	if v, ok := configuration[key]; ok {
		return v, nil
	}
	return "", errors.New("given variable not found")
}
