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
	boardID := c.Param("id")

	// VALIDASI: Kalau orang ngetik ID ngarang, buang dia ke halaman Home
	if !h.service.IsBoardValid(boardID) {
		c.Redirect(http.StatusFound, "/")
		return
	}

	activities, err := h.service.GetActivities(boardID)
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
		"board_id":   boardID,
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

	boardID := c.Param("id")

	// VALIDASI: Kalau orang ngetik ID ngarang
	if !h.service.IsBoardValid(boardID) {
		c.Redirect(http.StatusFound, "/")
		return
	}

	err = h.service.AddActivity(boardID, title, decisionID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/board/"+boardID+"/activities")
}

func (h *Handler) StartBoard(c *gin.Context) {
	boardID, err := h.service.GenerateNewBoard()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Gagal membuat sesi baru"})
		return
	}

	// Set Cookie, Expired 30 Hari
	c.SetCookie("recent_board_id", boardID, 2592000, "/", "", false, true)

	c.Redirect(http.StatusFound, "/board/"+boardID+"/activities")
}
