package services

import (
	"gitlab.com/velo-company/services/routes-service/internal/core/domain"
	"gitlab.com/velo-company/services/routes-service/internal/core/ports"
)

type CreateTrackServiceInput struct {
	InitialLocation string            `json:"initial_location" binding:"required,min=3"`
	FinalLocation   string            `json:"final_location" binding:"required,min=3"`
	Track           []domain.Location `json:"track" binding:"required,min=1,dive"`
}

type CreateTrackServiceOutput struct {
	Message    string `json:"message"`
	ID         *int   `json:"track_id"`
	StatusCode int    `json:"status_code"`
}

type CreateTrackService interface {
	Execute(input CreateTrackServiceInput, userId int) *CreateTrackServiceOutput
}

type createTrackService struct {
	SavePort           ports.SaveTrackPort
	UserExistsByIdPort ports.UserExistsByIdPort
}

func NewCreateTrackService(savePort ports.SaveTrackPort, userExistsByIdPort ports.UserExistsByIdPort) CreateTrackService {
	return &createTrackService{
		SavePort:           savePort,
		UserExistsByIdPort: userExistsByIdPort,
	}
}

func (c *createTrackService) Execute(input CreateTrackServiceInput, userId int) *CreateTrackServiceOutput {
	exists, err := c.UserExistsByIdPort.Execute(userId)

	if err != nil {
		return &CreateTrackServiceOutput{
			Message:    "Estamos enfrentando problemas no momento. Tente novamento mais tarde",
			ID:         nil,
			StatusCode: 502,
		}
	}

	if exists == false {
		return &CreateTrackServiceOutput{
			Message:    "Este usuário não existe",
			ID:         nil,
			StatusCode: 404,
		}
	}

	track := domain.NewTrack(nil, &userId, input.InitialLocation, input.FinalLocation, nil, input.Track)
	track.ID = c.SavePort.Execute(track)

	return &CreateTrackServiceOutput{
		Message:    "Criado",
		ID:         track.ID,
		StatusCode: 201,
	}
}
