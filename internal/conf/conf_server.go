package conf

type ServerConfig struct {
	Port         int      `mapstructure:"port" yaml:"port"`
	Mode         string   `mapstructure:"mode" yaml:"mode"` // debug / release
	AllowOrigins []string `mapstructure:"allow_origins" yaml:"allow_origins"`
	Version      string   `mapstructure:"mode" yaml:"version"`
}
