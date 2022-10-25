package route

import (
	loginAuth "github.com/053steve/gin-boilerplate/controllers/auth/login"
	registerAuth "github.com/053steve/gin-boilerplate/controllers/auth/register"
	handlerLogin "github.com/053steve/gin-boilerplate/handlers/auth-handlers/login"
	handlerRegister "github.com/053steve/gin-boilerplate/handlers/auth-handlers/register"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Auth
	*/
	LoginRepository := loginAuth.NewRepositoryLogin(db)
	loginService := loginAuth.NewServiceLogin(LoginRepository)
	loginHandler := handlerLogin.NewHandlerLogin(loginService)

	registerRepository := registerAuth.NewRepositoryRegister(db)
	registerService := registerAuth.NewServiceRegister(registerRepository)
	registerHandler := handlerRegister.NewHandlerRegister(registerService)

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/login", loginHandler.LoginHandler)

}
