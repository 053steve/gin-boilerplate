package main

import (
	"log"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"

	config "github.com/053steve/gin-boilerplate/configs"
	util "github.com/053steve/gin-boilerplate/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.Run()

	/**
	@description Setup Server
	*/
	router := SetupRouter()
	/**
	@description Run Server
	*/
	log.Fatal(router.Run(":" + util.GodotEnv("GO_PORT")))
}

func SetupRouter() *gin.Engine {
	/**
	  @description Setup Database Connection
	*/

	db := config.Connection()
	/**
	  @description Init Router
	*/
	router := gin.Default()
	/**
	  @description Setup Mode Application
	*/
	if util.GodotEnv("GO_ENV") != "production" && util.GodotEnv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if util.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	/**
	  @description Setup Middleware
	*/
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))
	/**
	  @description Init All Route
	*/
	route.InitAuthRoutes(db, router)
	// route.InitStudentRoutes(db, router)

	return router
}
