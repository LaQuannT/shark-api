package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/LaQuannT/shark-api/repository"
	"github.com/LaQuannT/shark-api/types"
)

type UserHandler struct {
	Repo   *repository.UserRepo
	Logger *logrus.Logger
}

func RegisterUserHandlers(router *gin.Engine, repo *repository.UserRepo, logger *logrus.Logger) {
	h := &UserHandler{Repo: repo, Logger: logger}

	router.POST("/users", h.HandleCreateUser)
	router.GET("/users/:userId", h.HandleGetUser)
	router.PUT("/users/:userId", h.HandleUpdateUser)
	router.DELETE("/user/:userId", h.HandleDeleteUser)
}

func (h *UserHandler) HandleCreateUser(c *gin.Context) {
	var user types.User

	if err := c.ShouldBindJSON(&user); err != nil {
		h.Logger.Errorf("UserHandler/HandleCreateUser/ShouldBindJSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key := uuid.New()
	user.ApiKey = key.String()

	if user.PermissionLevel != types.Admin {
		user.PermissionLevel = types.Base
	}

	id, err := h.Repo.CreateUser(user)
	if err != nil {
		h.Logger.Errorf("UserHandler/HandleCreateUser/CreateUser: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.Repo.GetUser(id)
	if err != nil {
		h.Logger.Errorf("UserHandler/HandleCreateUser/GetUser: %v", err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "user created - error returning user data"},
		)
	}

	c.JSON(http.StatusCreated, u)
}

func (h *UserHandler) HandleGetUser(c *gin.Context)    {}
func (h *UserHandler) HandleUpdateUser(c *gin.Context) {}
func (h *UserHandler) HandleDeleteUser(c *gin.Context) {}
