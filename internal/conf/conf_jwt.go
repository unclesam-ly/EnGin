package conf

type JwtConfig struct {
	AccessExpire  int    `mapstructure:"access_expire" yaml:"access_expire"`
	RefreshExpire int    `mapstructure:"refresh_expire" yaml:"refresh_expire"`
	AccessSecret  string `mapstructure:"access_secret" yaml:"access_secret"`
	RefreshSecret string `mapstructure:"refresh_secret" yaml:"refresh_secret"`
	Issuer        string `mapstructure:"issuer" yaml:"issuer"`
}
