package http

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"gitlab.com/velo-company/services/routes-service/internal/adapters/database"
	"gitlab.com/velo-company/services/routes-service/internal/core/services"
)

func CreateTrackHandler(c *gin.Context, DB *sql.DB) {
	var savePort = database.NewSaveTrackAdapter(DB)
	var useCase = services.NewCreateTrackService(
		savePort,
		nil)

	emailValue, exists := c.Get("email")
	if !exists {
		c.JSON(401, gin.H{
			"message":     "Invalid token",
			"status_code": 401,
		})
		return
	}

	email, ok := emailValue.(string)
	if !ok {
		c.JSON(500, gin.H{
			"message":     "Invalid token",
			"status_code": 401,
		})
		return
	}

	var input services.CreateTrackServiceInput

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"message":     "Requisição invalida",
			"status_code": 400,
		})
		return
	}

	result := useCase.Execute(input, email)

	c.JSON(result.StatusCode, result)
}
