package routers

import (
	_ "GinWeb/docs"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetupRouters(engine *gin.Engine) {
	engine.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	LoadKnowledgeRouters(engine)
}
