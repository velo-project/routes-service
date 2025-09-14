package services

import (
	"time"

	"gitlab.com/velo-company/services/routes-service/internal/core/domain"
	"gitlab.com/velo-company/services/routes-service/internal/core/ports"
)

type CreateTrackServiceInput struct {
	UserID          *int              `json:"user_id"`
	InitialLocation string            `json:"initial_location"`
	FinalLocation   string            `json:"final_location"`
	VisitedAt       *time.Time        `json:"visited_at"`
	Track           []domain.Location `json:"track"`
}

type CreateTrackServiceOutput struct {
	Message    string `json:"message"`
	ID         *int   `json:"track_id"`
	StatusCode int    `json:"status_code"`
}

type CreateTrackService interface {
	Execute(input CreateTrackServiceInput, email string) *CreateTrackServiceOutput
}

type createTrackService struct {
	SavePort             ports.SaveTrackPort
	GetUserIdByEmailPort ports.GetUserIdByEmailPort
}

func NewCreateTrackService(savePort ports.SaveTrackPort, userService ports.GetUserIdByEmailPort) CreateTrackService {
	return &createTrackService{
		SavePort:             savePort,
		GetUserIdByEmailPort: userService,
	}
}

func (c *createTrackService) Execute(input CreateTrackServiceInput, email string) *CreateTrackServiceOutput {
	userId := c.GetUserIdByEmailPort.Execute(email)

	if userId == nil {
		return &CreateTrackServiceOutput{
			Message:    "Este usuário não existe",
			ID:         nil,
			StatusCode: 404,
		}
	}

	track := domain.NewTrack(nil, userId, input.InitialLocation, input.FinalLocation, nil, input.Track)
	track.ID = c.SavePort.Execute(track)

	return &CreateTrackServiceOutput{
		Message:    "Criado",
		ID:         track.ID,
		StatusCode: 201,
	}
}
