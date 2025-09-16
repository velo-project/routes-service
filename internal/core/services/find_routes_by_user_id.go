package services

import (
	"gitlab.com/velo-company/services/routes-service/internal/core/domain"
	"gitlab.com/velo-company/services/routes-service/internal/core/ports"
)

type FindRoutesByUserId interface {
	Execute(input FindRoutesByUserIdInput) FindRoutesByUserIdOutput
}

type findRoutesByUserId struct {
	FindRoutesByUserIdPort ports.FindByUserIDPort
	UserExistsByIdPort     ports.UserExistsByIdPort
}

type FindRoutesByUserIdInput struct {
	UserId int
}

type FindRoutesByUserIdOutput struct {
	Message    string         `json:"message"`
	Tracks     []domain.Track `json:"tracks"`
	StatusCode int            `json:"status_code"`
}

func NewFindRoutesByUserId(findRoutePort ports.FindByUserIDPort, verifyUserPort ports.UserExistsByIdPort) FindRoutesByUserId {
	return &findRoutesByUserId{
		FindRoutesByUserIdPort: findRoutePort,
		UserExistsByIdPort:     verifyUserPort,
	}
}

func (f findRoutesByUserId) Execute(input FindRoutesByUserIdInput) FindRoutesByUserIdOutput {
	exists, err := f.UserExistsByIdPort.Execute(input.UserId)
	if err != nil {
		return FindRoutesByUserIdOutput{
			Message:    "Estamos enfrentando problemas no momento. Tente novamento mais tarde",
			Tracks:     []domain.Track{},
			StatusCode: 502,
		}
	}
	if !exists {
		return FindRoutesByUserIdOutput{
			Message:    "Este usuário não existe",
			Tracks:     []domain.Track{},
			StatusCode: 404,
		}
	}

	tracks := f.FindRoutesByUserIdPort.Execute(&input.UserId)

	if len(tracks) == 0 {
		return FindRoutesByUserIdOutput{
			Message:    "Este usuário não tem rotas salvas",
			Tracks:     []domain.Track{},
			StatusCode: 404,
		}
	}

	return FindRoutesByUserIdOutput{
		Message:    "OK",
		Tracks:     tracks,
		StatusCode: 200,
	}
}
