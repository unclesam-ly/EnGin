package conf

type ServerConfig struct {
	Port int    `mapstructure:"port" yaml:"port"`
	Mode string `mapstructure:"mode" yaml:"mode"` // debug / release
}
