package services

import (
	"gitlab.com/velo-company/services/routes-service/internal/core/ports"
)

type DeleteTrack interface {
	Execute(input DeleteTrackInput) DeleteTrackOutput
}

type deleteTrack struct {
	DeleteTrackPort    ports.DeleteTrackPort
	UserExistsByIdPort ports.UserExistsByIdPort
}

type DeleteTrackInput struct {
	UserId  int
	TrackId int
}

type DeleteTrackOutput struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func NewDeleteTrack(dt ports.DeleteTrackPort, fu ports.UserExistsByIdPort) DeleteTrack {
	return &deleteTrack{
		DeleteTrackPort:    dt,
		UserExistsByIdPort: fu,
	}
}

func (d deleteTrack) Execute(input DeleteTrackInput) DeleteTrackOutput {
	exists, err := d.UserExistsByIdPort.Execute(input.UserId)
	if err != nil {
		return DeleteTrackOutput{
			Message:    "Estamos enfrentando problemas no momento. Tente novamento mais tarde",
			StatusCode: 502,
		}
	}
	if !exists {
		return DeleteTrackOutput{
			Message:    "Este usuário não existe",
			StatusCode: 404,
		}
	}

	err = d.DeleteTrackPort.Execute(&input.TrackId)

	if err != nil {
		return DeleteTrackOutput{
			Message:    "Não conseguimos apagar essa rota. Tente novamente mais tarde",
			StatusCode: 500,
		}
	}

	return DeleteTrackOutput{
		Message:    "OK",
		StatusCode: 200,
	}
}
