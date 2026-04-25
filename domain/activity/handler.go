package activity

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nullablenone/life-decision-tracker/domain/decision"
)

type Handler struct {
	service         Service
	decisionService decision.Service
}

func NewHandler(service Service, decisionService decision.Service) *Handler {
	return &Handler{service, decisionService}
}

func (h *Handler) Index(c *gin.Context) {
	activities, err := h.service.GetActivities()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}

	categories, err := h.decisionService.GetDecisionCategories()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "activities.html", gin.H{
		"activities": activities,
		"categories": categories,
	})
}

func (h *Handler) Store(c *gin.Context) {
	title := c.PostForm("title")
	decisionIDStr := c.PostForm("decision_id")

	decisionID, err := strconv.Atoi(decisionIDStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "Invalid decision ID"})
		return
	}

	err = h.service.AddActivity(title, decisionID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/activities")
}
