package grpc

import (
	"context"
	"log"

	"gitlab.com/velo-company/services/routes-service/internal/core/ports"
	"gitlab.com/velo-company/services/routes-service/proto/user"
	"google.golang.org/grpc"
)

type userExistsByIdAdapter struct {
	client user.UserServiceClient
}

func NewUserExistsByIdAdapter(connection *grpc.ClientConn) ports.UserExistsByIdPort {
	return &userExistsByIdAdapter{
		client: user.NewUserServiceClient(connection),
	}
}

func (u userExistsByIdAdapter) Execute(userId int) (bool, error) {
	res, err := u.client.UserExistsById(context.Background(), &user.UserExistsByIdRequest{
		Id: int32(userId),
	})

	if err != nil {
		log.Print(err.Error())
		return false, err
	}

	return res.Exists, nil
}

// JUST FOR DEVELOPMENT
/*func (u userExistsByIdAdapter) Execute(userId int) (bool, error) {
	return true, nil
}*/
