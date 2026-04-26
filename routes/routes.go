package routes

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nullablenone/life-decision-tracker/domain/activity"
	"github.com/nullablenone/life-decision-tracker/domain/decision"
	"github.com/nullablenone/life-decision-tracker/domain/home"
)

func generateBoardID() string {
	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err != nil {
		return "default"
	}
	return hex.EncodeToString(bytes)
}

func SetRoutes(decisionHandler *decision.Handler, activityHandler *activity.Handler, homeHandler *home.Home) *gin.Engine {
	router := gin.Default()

	// API
	router.GET("/api/decision", decisionHandler.GetDecisionCategories)

	// Load templates
	router.LoadHTMLGlob("templates/*")

	// Home route
	router.GET("/", homeHandler.HomePage)

	// Route baru buat generate Magic URL
	router.GET("/start", func(c *gin.Context) {
		boardID := generateBoardID()
		c.Redirect(http.StatusFound, "/board/"+boardID+"/activities")
	})

	// Activity routes diubah jadi pakai parameter :id
	router.GET("/board/:id/activities", activityHandler.Index)
	router.POST("/board/:id/activities", activityHandler.Store)

	return router
}
