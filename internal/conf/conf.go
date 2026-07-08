package conf

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server" yaml:"server"`
	Database DatabaseConfig `mapstructure:"database" yaml:"database"`
	Logger   LoggerConfig   `mapstructure:"logger" yaml:"logger"`
	Jwt      JwtConfig      `mapstructure:"jwt" yaml:"jwt"`
	Redis    RedisConfig    `mapstructure:"redis" yaml:"redis"`
}

// LoadConfig 读取并解析配置文件
func LoadConfig(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	cfg := &Config{}
	if err := v.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	// 监听配置变更
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("配置文件已更新: %s\n", e.Name)
		if err := v.Unmarshal(cfg); err != nil {
			fmt.Printf("重载配置文件失败: %s\n", err)
		} else {
			fmt.Println("🚀 配置热重载成功，新配置已即时生效！")
		}
	})

	return cfg, nil
}
