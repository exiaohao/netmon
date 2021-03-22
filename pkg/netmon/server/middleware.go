package server

import "github.com/gin-gonic/gin"

func AgentWrap() gin.HandlerFunc {
	return func(c *gin.Context) {
		agentID := c.GetHeader("X-NETMON-AGENT")
		// todo: check agent is valid

		c.Set("agentID", agentID)
		c.Next()
	}
}