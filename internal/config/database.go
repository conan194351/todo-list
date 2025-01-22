package config

type Database struct {
	Host string `mapstructure:"host" `
	Port string `mapstructure:"port"`
	User string `mapstructure:"user" `
	Pass string `mapstructure:"pass"`
	Name string `mapstructure:"name" `
}
