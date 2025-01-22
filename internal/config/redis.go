package config

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"post"`
	Password string `mapstructure:"password"`
}
