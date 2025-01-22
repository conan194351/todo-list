package setup

import (
	"fmt"
	"github.com/conan194351/todo-list.git/internal/config"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDSN(env config.Database) string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		env.Host, env.Port, env.User, env.Pass, env.Name)
	return dsn
}

func InitDatabase() *gorm.DB {
	env := config.GetConfig().Database
	source := GetDSN(env)
	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  source,
				PreferSimpleProtocol: true,
			},
		), &gorm.Config{
			//Logger: gormlog.Default.LogMode(gormlog.Info),
		},
	)
	if err != nil {
		return nil
	}
	return db
}

var DBConnection = fx.Provide(InitDatabase)
