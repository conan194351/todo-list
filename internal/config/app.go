package config

import "github.com/gin-gonic/gin"

type App struct {
	Env      string `mapstructure:"env" json:"env"`
	Timezone string `mapstructure:"timezone" json:"timezone"`
	LogPath  string `mapstructure:"log_path" json:"logPath"`
	LogLevel string `mapstructure:"log_level" json:"logLevel"`
	Version  string `mapstructure:"version" json:"version"`
}

func (app App) IsProduction() bool {
	return app.Env == "production"
}

func (app App) IsDevelopment() bool {
	return app.Env == "development"
}

func (app App) IsTest() bool {
	return app.Env == "test"
}

func (app App) GetMode() string {
	if app.IsProduction() {
		return gin.ReleaseMode
	}

	return gin.DebugMode
}
