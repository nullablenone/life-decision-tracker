package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nullablenone/life-decision-tracker/domain/activity"
	"github.com/nullablenone/life-decision-tracker/domain/decision"
	"github.com/nullablenone/life-decision-tracker/domain/home"
	"github.com/nullablenone/life-decision-tracker/middleware"
)

func SetRoutes(decisionHandler *decision.Handler, activityHandler *activity.Handler, homeHandler *home.Home) *gin.Engine {
	router := gin.Default()

	// API
	router.GET("/api/decision", decisionHandler.GetDecisionCategories)

	// Load templates
	router.LoadHTMLGlob("templates/*")

	// Home route
	router.GET("/", homeHandler.HomePage)

	// Route baru buat generate Magic URL
	router.GET("/start", middleware.IPRateLimit(), activityHandler.StartBoard)

	// Activity routes diubah jadi pakai parameter :id
	router.GET("/board/:id/activities", activityHandler.Index)
	router.POST("/board/:id/activities", activityHandler.Store)

	return router
}
