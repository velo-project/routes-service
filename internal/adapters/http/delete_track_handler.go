package http

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/velo-company/services/routes-service/internal/adapters/database"
	"gitlab.com/velo-company/services/routes-service/internal/adapters/grpc"
	"gitlab.com/velo-company/services/routes-service/internal/core/services"
	ggrpc "google.golang.org/grpc"
)

func DeleteTrackHandler(c *gin.Context, DB *sql.DB, grpcConn *ggrpc.ClientConn) {
	userExistsByIdPort := grpc.NewUserExistsByIdAdapter(grpcConn)
	deleteTrackPort := database.NewDeleteTrackAdapter(DB)
	service := services.NewDeleteTrack(deleteTrackPort, userExistsByIdPort)

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

	trackIdStr := c.Param("id")

	if trackIdStr == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Track Id is empty"})
	}

	trackId, err := strconv.Atoi(trackIdStr)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Track Id is invalid"})
	}

	result := service.Execute(services.DeleteTrackInput{
		UserId:  userId,
		TrackId: trackId,
	})

	c.JSON(result.StatusCode, result)
}
