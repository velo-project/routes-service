package ports

import "gitlab.com/velo-company/services/routes-service/internal/core/domain"

type SaveTrackPort interface {
	Execute(track *domain.Track) *int
}

type FindByUserIDPort interface {
	Execute(userId *int) []domain.Track
}
