package routers

import (
	"GinWeb/controllers/knowledgeController"

	"github.com/gin-gonic/gin"
)

func LoadKnowledgeRouters(e *gin.Engine) {
	e.GET("ping", knowledgeController.GetKnowledges)
	e.POST("knowledge/createDemo", knowledgeController.CreateKnowledgeMain)
}
