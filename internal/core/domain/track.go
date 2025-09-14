package domain

type Track struct {
	ID     *int       `json:"id,omitempty"`
	UserID *int       `json:"user_id"`
	Track  []Location `json:"track"`
}

func NewTrack(id, userId *int, track []Location) *Track {
	return &Track{
		ID:     id,
		UserID: userId,
		Track:  track,
	}
}
