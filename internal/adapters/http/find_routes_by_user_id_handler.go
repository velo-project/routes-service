package http

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/velo-company/services/routes-service/internal/adapters/database"
	"gitlab.com/velo-company/services/routes-service/internal/adapters/grpc"
	"gitlab.com/velo-company/services/routes-service/internal/core/services"
	ggrpc "google.golang.org/grpc"
)

func FindRoutesByUserIdHandler(c *gin.Context, DB *sql.DB, grpcConn *ggrpc.ClientConn) {
	userExistsByIdPort := grpc.NewUserExistsByIdAdapter(grpcConn)
	findByUserIdPort := database.NewFindByUserIDAdapter(DB)
	service := services.NewFindRoutesByUserId(
		findByUserIdPort, userExistsByIdPort)

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

	result := service.Execute(services.FindRoutesByUserIdInput{
		UserId: userId,
	})

	c.JSON(result.StatusCode, result)
}
