package database

import (
	"database/sql"
	"log"

	"gitlab.com/velo-company/services/routes-service/internal/core/ports"
)

type deleteTrackAdapter struct {
	DB *sql.DB
}

func NewDeleteTrackAdapter(DB *sql.DB) ports.DeleteTrackPort {
	return &deleteTrackAdapter{
		DB: DB,
	}
}

const (
	deleteTrackQuery         = `DELETE FROM tbl_tracks WHERE id_track = $1`
	deleteTrackLocationQuery = `DELETE FROM tbl_locations WHERE fk_track = $1`
)

func (d deleteTrackAdapter) Execute(trackID *int) error {
	tx, err := d.DB.Begin()
	if err != nil {
		log.Printf("ERROR: could not begin transaction: %v", err)
		return err
	}

	_, err = tx.Exec(deleteTrackLocationQuery, *trackID)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(deleteTrackQuery, *trackID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("ERROR: could not commit transaction: %v", err)
		return err
	}

	return nil
}
