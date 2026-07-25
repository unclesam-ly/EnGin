package conf

type ServerConfig struct {
	Port         int      `yaml:"port"`
	Mode         string   `yaml:"mode"` // debug / release
	AllowOrigins []string `yaml:"allow_origins"`
	Version      string   `yaml:"version"`
}
