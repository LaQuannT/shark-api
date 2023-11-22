package handlers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/LaQuannT/shark-api/repository"
)

func InitHandlers(router *gin.Engine, db *sql.DB, log *logrus.Logger) {
	usr := repository.NewUserRepo(db)
	srk := repository.NewSharkRepo(db)

	RegisterUserHandlers(router, usr, log)
	RegisterSharkHandlers(router, srk, usr, log)
}
