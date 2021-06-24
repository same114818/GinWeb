package knowledgeController

import (
	"GinWeb/db/repositories/knowledgeRepository"
	"GinWeb/global"
	konwledgeOutputs "GinWeb/outputs/knowledgeOutputs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 接口1
// @Description xxxxxxxxxxxx
// @Tags 相关接口
// @Accept application/json
// @Produce application/json
// @Success 200
//@Router /ping [get]
func GetKnowledges(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// @Summary 接口2
// @Description yyyyyyyyyyyyyyy
// @Tags 相关接口
// @Accept application/json
// @Produce application/json
// @Success 200
//@Router /knowledge/createDemo [post]
func CreateKnowledgeMain(c *gin.Context) {
	err := knowledgeRepository.Create(global.DBEngine)
	if err != nil {
		return
	}
	result := konwledgeOutputs.UpdateStatusOutput{
		IsSuccess: true,
		Msg:       "success",
	}
	c.JSON(http.StatusOK, result)
}
