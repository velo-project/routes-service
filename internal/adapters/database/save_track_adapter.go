package database

import (
	"database/sql"
	"log"

	"gitlab.com/velo-company/services/routes-service/internal/core/domain"
)

type SaveTrackAdapter struct {
	DB *sql.DB
}

func NewSaveTrackAdapter(db *sql.DB) *SaveTrackAdapter {
	return &SaveTrackAdapter{DB: db}
}

const (
	saveTrackQuery         = `INSERT INTO tracks (user_id, initial_location, final_location, visited_at) VALUES ($1, $2, $3, $4) RETURNING id;`
	saveTrackLocationQuery = `INSERT INTO track_locations (track_id, lat, lng) VALUES ($1, $2, $3);`
)

func (a *SaveTrackAdapter) Execute(track *domain.Track) *int {
	tx, err := a.DB.Begin()
	if err != nil {
		log.Printf("ERROR: could not begin transaction: %v", err)
		return nil
	}

	var trackID int
	err = tx.QueryRow(saveTrackQuery, track.UserID, track.InitialLocation, track.FinalLocation, track.VisitedAt).Scan(&trackID)
	if err != nil {
		tx.Rollback()
		log.Printf("ERROR: could not save track: %v", err)
		return nil
	}

	for _, loc := range track.Track {
		_, err := tx.Exec(saveTrackLocationQuery, trackID, loc.Lat, loc.Lng)
		if err != nil {
			tx.Rollback()
			log.Printf("ERROR: could not save track location: %v", err)
			return nil
		}
	}

	if err := tx.Commit(); err != nil {
		log.Printf("ERROR: could not commit transaction: %v", err)
		return nil
	}

	return &trackID
}
