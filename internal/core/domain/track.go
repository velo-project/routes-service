package domain

import "time"

type Track struct {
	ID              *int       `json:"id,omitempty"`
	UserID          *int       `json:"user_id"`
	InitialLocation string     `json:"initial_location"`
	FinalLocation   string     `json:"final_location"`
	VisitedAt       *time.Time `json:"visited_at"`
	Track           []Location `json:"track"`
}

func NewTrack(id, userId *int, initialLocation, finalLocation string, visitedAt *time.Time, track []Location) *Track {
	return &Track{
		ID:              id,
		UserID:          userId,
		InitialLocation: initialLocation,
		FinalLocation:   finalLocation,
		VisitedAt:       visitedAt,
		Track:           track,
	}
}
