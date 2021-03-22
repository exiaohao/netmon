package server

import (
	"github.com/exiaohao/netmon/pkg/netmon/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	// For local api requests
	//apiGroup := r.Group("/api")
	// For Agents
	agentGroup := r.Group("/agent")
	agentGroup.Use(AgentWrap())
	agentGroup.GET("/task", controller.GetTasks)
	agentGroup.POST("/task", controller.SubmitTaskResult)
}