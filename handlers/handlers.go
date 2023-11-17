package handlers

import (
	"database/sql"

	"github.com/LaQuannT/shark-api/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitHandlers(router *gin.Engine, db *sql.DB, log *logrus.Logger) {
	ur := repository.NewUserRepo(db)
	RegisterUserHandlers(router,ur,log)
}