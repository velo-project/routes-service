package database

import (
	"database/sql"
	"log"

	"gitlab.com/velo-company/services/routes-service/internal/core/domain"
)

type FindByUserIDAdapter struct {
	DB *sql.DB
}

func NewFindByUserIDAdapter(db *sql.DB) *FindByUserIDAdapter {
	return &FindByUserIDAdapter{DB: db}
}

const (
	findTracksByUserIDQuery = `SELECT id, initial_location, final_location, visited_at FROM tracks WHERE user_id = $1;`
	findTrackLocationsQuery = `SELECT lat, lng FROM track_locations WHERE track_id = $1;`
)

func (a *FindByUserIDAdapter) Execute(userId *int) []domain.Track {
	rows, err := a.DB.Query(findTracksByUserIDQuery, userId)
	if err != nil {
		log.Printf("ERROR: could not find tracks by user id: %v", err)
		return nil
	}
	defer rows.Close()

	var tracks []domain.Track
	for rows.Next() {
		var track domain.Track
		var trackID int
		track.UserID = userId

		if err := rows.Scan(&trackID, &track.InitialLocation, &track.FinalLocation, &track.VisitedAt); err != nil {
			log.Printf("ERROR: could not scan track: %v", err)
			continue
		}
		track.ID = &trackID

		locations, err := a.fetchLocationsForTrack(trackID)
		if err != nil {
			log.Printf("ERROR: could not fetch locations for track %d: %v", trackID, err)
		}
		track.Track = locations
		tracks = append(tracks, track)
	}

	return tracks
}

func (a *FindByUserIDAdapter) fetchLocationsForTrack(trackID int) ([]domain.Location, error) {
	locationRows, err := a.DB.Query(findTrackLocationsQuery, trackID)
	if err != nil {
		return nil, err
	}
	defer locationRows.Close()

	var locations []domain.Location
	for locationRows.Next() {
		var loc domain.Location
		if err := locationRows.Scan(&loc.Lat, &loc.Lng); err != nil {
			return nil, err
		}
		locations = append(locations, loc)
	}
	return locations, nil
}
