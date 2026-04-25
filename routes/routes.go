package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nullablenone/life-decision-tracker/domain/activity"
	"github.com/nullablenone/life-decision-tracker/domain/decision"
	"github.com/nullablenone/life-decision-tracker/domain/home"
)

func SetRoutes(decisionHandler *decision.Handler, activityHandler *activity.Handler, homeHandler *home.Home) *gin.Engine {
	router := gin.Default()

	// API
	router.GET("/api/decision", decisionHandler.GetDecisionCategories)

	// Load templates
	router.LoadHTMLGlob("templates/*")

	// Home route
	router.GET("/", homeHandler.HomePage)

	// Activity routes
	router.GET("/activities", activityHandler.Index)
	router.POST("/activities", activityHandler.Store)

	return router
}
