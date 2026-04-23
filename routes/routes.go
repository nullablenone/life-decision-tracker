package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nullablenone/life-decision-tracker/domain/decision"
)

func SetRoutes(decisionHandler *decision.Handler) *gin.Engine {
	router := gin.Default()

	router.GET("/decision", decisionHandler.GetDecisionCategories)

	return router
}
