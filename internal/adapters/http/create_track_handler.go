package http

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/velo-company/services/routes-service/internal/adapters/database"
	"gitlab.com/velo-company/services/routes-service/internal/adapters/grpc"
	"gitlab.com/velo-company/services/routes-service/internal/core/services"
)

func CreateTrackHandler(c *gin.Context, DB *sql.DB) {
	var savePort = database.NewSaveTrackAdapter(DB)
	var userExistsByIdAdapter = grpc.NewUserExistsByIdAdapter(nil)
	var useCase = services.NewCreateTrackService(
		savePort,
		userExistsByIdAdapter)

	value, exists := c.Get("userId")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	userId, ok := value.(int)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "UserId invalid type"})
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

	result := useCase.Execute(input, userId)

	c.JSON(result.StatusCode, result)
}
