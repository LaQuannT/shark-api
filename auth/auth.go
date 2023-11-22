package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/LaQuannT/shark-api/repository"
)

func getApiKey(h http.Header) (string, error) {
	authStr := h.Get("Authorization")
	if authStr == "" {
		return "", errors.New("no authorization string found")
	}

	val := strings.Split(authStr, " ")
	if len(val) != 2 {
		return "", errors.New("invalid authorization header")
	}

	if val[0] != "ApiKey" {
		return "", errors.New("invalid first part authorization header")
	}

	return val[1], nil
}

func ValidateApiKey(repo *repository.UserRepo, logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		key, err := getApiKey(c.Request.Header)
		if err != nil {
			logger.Errorf("Auth/ValidateApiKey/getApiKey: %v", err)
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		isValid, err := repo.CheckApiKey(key)
		if err != nil {
			logger.Errorf("Auth/ValidateApiKey/User.Repository/CheckApiKey: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if !isValid {
			logger.Errorf("Auth/ValidateApiKey/User.Repository/CheckApiKey: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
