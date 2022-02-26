package backend

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"notans/backend/config"
	"notans/backend/routes"
	"notans/backend/service"
	"os"
	"time"
)

type IApp struct {
	Router *gin.Engine
}

func (a *IApp) Initialize(config *config.Config) {
	var dialect gorm.Dialector = nil

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,
		},
	)

	if config.DB.Driver == "postgres" {
		dialect = postgres.Open(config.DB.Dsn)
	} else {
		dialect = sqlite.Open(config.DB.Dsn)
	}

	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("Could not connect database")
	}
	a.Router = gin.New()

	rtr := routes.IRoutes{
		Config: config,
	}
	rtr.Bind(a.Router)

	srv := service.IService{
		Db:     db,
		Config: config,
	}
	srv.InitializeService()
}
