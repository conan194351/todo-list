package config

type Server struct {
	Server     string `mapstructure:"host_port"`
	CorsDomain string `mapstructure:"cors_domain"`
	SecretKey  string `mapstructure:"secret_key"`
	GinMode    string `mapstructure:"gin_mode"`
}
