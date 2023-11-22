package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/LaQuannT/shark-api/auth"
	"github.com/LaQuannT/shark-api/repository"
	"github.com/LaQuannT/shark-api/types"
)

type SharkHandler struct {
	Repo   *repository.SharkRepo
	Logger *logrus.Logger
}

type Search struct {
	Name string `form:"name"`
}

func (s *Search) Stringf() {
	c := cases.Title(language.English)
	s.Name = c.String(strings.ToLower(s.Name))
}

func RegisterSharkHandlers(
	router *gin.Engine,
	repo *repository.SharkRepo,
	uRepo *repository.UserRepo,
	logger *logrus.Logger,
) {
	h := &SharkHandler{Repo: repo, Logger: logger}

	router.GET("/sharks", auth.ValidateApiKey(uRepo, logger), h.HandleGetSharkByName)
	router.POST("/sharks", h.HandleCreateShark)
	router.GET("/sharks/:sharkId", h.HandleGetSharkById)
	router.PUT("/sharks/:sharkId", h.HandleUpdateShark)
	router.DELETE("/sharks/:sharkId", h.HandleDeleteShark)
}

func (h *SharkHandler) HandleCreateShark(c *gin.Context) {
	var shark types.Shark

	if err := c.ShouldBindJSON(&shark); err != nil {
		h.Logger.Errorf("SharkHandler/HandleCreateShark/ShouldBindJSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.Repo.CreateShark(shark)
	if err != nil {
		h.Logger.Errorf("SharkHandler/HandleCreateShark/CreateShark: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s, err := h.Repo.GetSharkById(id)
	if err != nil {
		h.Logger.Errorf("SharkHandler/HandleCreateShark/GetSharkById: %v", err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "shark created - error returing shark data"},
		)
		return
	}

	c.JSON(http.StatusCreated, s)
}

func (h *SharkHandler) HandleGetSharkById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("sharkId"))
	if err != nil {
		h.Logger.Errorf("SharkHandler/handlerGetSharkById/Atoi: %v", err)
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Invalid shark id param - must be a valid number"},
		)
		return
	}

	s, err := h.Repo.GetSharkById(id)
	if err != nil {
		h.Logger.Errorf("SharkHandler/HandleGetSharkById/GetSharkById: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, s)
}

func (h *SharkHandler) HandleGetSharkByName(c *gin.Context) {
	var search Search

	if err := c.ShouldBind(&search); err != nil {
		h.Logger.Errorf("SharkHandler/HandleGetSharkByName/ShouldBind: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	search.Stringf()

	shark, err := h.Repo.GetSharkByName(search.Name)
	if err != nil {
		h.Logger.Errorf("SharkHandler/HandleGetSharkByName/GetSharkByName: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shark)
}

func (h *SharkHandler) HandleUpdateShark(c *gin.Context) {
	var shark types.Shark
	id, err := strconv.Atoi(c.Param("sharkId"))
	if err != nil {
		h.Logger.Errorf("SharkHandler/HandleUpdateShark/Atoi: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&shark); err != nil {
		h.Logger.Errorf("SharkHandler/HandleUpdateShark/ShouldBindJSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shark.Id = id

	if err := h.Repo.UpdateShark(shark); err != nil {
		h.Logger.Errorf("SharkHandler/HandleUpdateShark/UpdateShark: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s, err := h.Repo.GetSharkById(id)
	if err != nil {
		h.Logger.Errorf("SharkHandler/HandleUpdateShark/GetSharkById: %v", err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Shark Updated - error returnin shark data"},
		)
		return
	}

	c.JSON(http.StatusOK, s)
}

func (h *SharkHandler) HandleDeleteShark(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("sharkId"))
	if err != nil {
		h.Logger.Errorf("SharkHandler/HandleDeleteShark/Atoi: %v", err)
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Invalid shark id param - must be a valid number"},
		)
		return
	}

	if err := h.Repo.DeleteShark(id); err != nil {
		h.Logger.Errorf("SharkHandler/HandleDeleteShark/DeleteShark: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Shark %d - successfully deleted", id)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}
