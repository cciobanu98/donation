package app

import (
	"log"

	handler "github.com/cciobanu98/donation/app/handlers"
	"github.com/cciobanu98/donation/app/models/project"
	"github.com/cciobanu98/donation/config"
	docs "github.com/cciobanu98/donation/docs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Package sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// App struct
type App struct {
	DB 		*gorm.DB
	Engine 	*gin.Engine
}

// Initialize app
func (a *App) Initialize(config config.Config) {
	// dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
	// 	config.DB.Username,
	// 	config.DB.Password,
	// 	config.DB.Host,
	// 	config.DB.Port,
	// 	config.DB.Name,
	// 	config.DB.Charset)
	docs.SwaggerInfo.BasePath = "/api"
	db, err := gorm.Open(config.DB.Dialect, config.DB.DbURI)
	if err != nil {
		log.Fatal("Could not connect databse")
	}
	db.AutoMigrate(&project.Project{})
	a.Engine = gin.New()
	a.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	a.setProjectRoutes()
	repo := project.NewRepository(db)
	handler.NewHandler(repo)
	a.DB = db
}

func (a *App) setProjectRoutes() {
	a.Engine.GET("api/project", handler.GetProjects)
	a.Engine.POST("api/project", handler.CreateProject)
	a.Engine.GET("api/project/:id", handler.GetProject)
	a.Engine.DELETE("api/project/:id", handler.Delete)
}

// Run app on host
func (a *App) Run(host string) {
	err := a.Engine.Run(host)
	if err != nil {
		log.Fatalf("Error on start %v", err)
	}
}