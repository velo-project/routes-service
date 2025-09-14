package grpc

import (
	"context"

	"gitlab.com/velo-company/services/routes-service/internal/core/ports"
	"gitlab.com/velo-company/services/routes-service/proto/user"
	"google.golang.org/grpc"
)

type getUserIdByEmailAdapter struct {
	client user.UserServiceClient
}

func NewGetUserIdByEmailAdapter(connection *grpc.ClientConn) ports.GetUserIdByEmailPort {
	return &getUserIdByEmailAdapter{
		client: user.NewUserServiceClient(connection),
	}
}

func (a *getUserIdByEmailAdapter) Execute(email string) (*int, error) {
	res, err := a.client.GetUserIdByEmail(context.Background(), &user.GetUserIdByEmailRequest{
		Email: email,
	})

	if err != nil {
		return nil, err
	}

	userId := int(res.Id)
	return &userId, nil
}

var _ ports.GetUserIdByEmailPort = (*getUserIdByEmailAdapter)(nil)
