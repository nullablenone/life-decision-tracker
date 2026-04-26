package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nullablenone/life-decision-tracker/domain/activity"
	"github.com/nullablenone/life-decision-tracker/domain/decision"
)

type Home struct {
	service         activity.Service
	decisionService decision.Service
}

func NewHome(service activity.Service, decisionService decision.Service) *Home {
	return &Home{service, decisionService}
}

func (h *Home) HomePage(c *gin.Context) {

	categories, err := h.decisionService.GetDecisionCategories()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "home.html", gin.H{

		"categories": categories,
	})
}
