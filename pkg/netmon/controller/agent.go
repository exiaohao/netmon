package controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

func GetTasks(c *gin.Context) {
	log.Printf("Req AgentID=%s", c.Get("AgentID"))
}

func SubmitTaskResult(c *gin.Context) {

}