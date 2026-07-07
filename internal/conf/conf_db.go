package conf

import "fmt"

// DatabaseConfig 定义数据库连接配置结构体
type DatabaseConfig struct {
	Driver   string `mapstructure:"driver" yaml:"driver"`
	Host     string `mapstructure:"host" yaml:"host"`
	Port     int    `mapstructure:"port" yaml:"port"`
	User     string `mapstructure:"user" yaml:"user"`
	Password string `mapstructure:"password" yaml:"password"`
	Dbname   string `mapstructure:"dbname" yaml:"dbname"`
}

// Dsn 根据 driver 动态生成 DSN 字符串
func (db DatabaseConfig) Dsn() string {
	switch db.Driver {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			db.User, db.Password, db.Host, db.Port, db.Dbname)
	case "postgres", "postgresql":
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			db.Host, db.Port, db.User, db.Password, db.Dbname)
	default:
		return ""
	}
}
