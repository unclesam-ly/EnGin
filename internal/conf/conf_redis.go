package conf

type RedisConfig struct {
	Addr     string `mapstructure:"addr" yaml:"addr"`
	Password string `mapstructure:"password" yaml:"password"`
	Db       int    `mapstructure:"db" yaml:"db"`
}
