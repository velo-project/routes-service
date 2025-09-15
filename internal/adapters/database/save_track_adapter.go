package database

import (
	"database/sql"
	"log"

	"gitlab.com/velo-company/services/routes-service/internal/core/domain"
	"gitlab.com/velo-company/services/routes-service/internal/core/ports"
)

type saveTrackAdapter struct {
	DB *sql.DB
}

func NewSaveTrackAdapter(db *sql.DB) ports.SaveTrackPort {
	return &saveTrackAdapter{DB: db}
}

const (
	saveTrackQuery         = `INSERT INTO tbl_tracks (fk_user, tx_initial_location, tx_final_location, tx_visited_at) VALUES ($1, $2, $3, $4) RETURNING id_track;`
	saveTrackLocationQuery = `INSERT INTO tbl_locations (fk_track, tx_lat, tx_lng) VALUES ($1, $2, $3);`
)

func (a *saveTrackAdapter) Execute(track *domain.Track) *int {
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
