package dbrepo

import (
	"errors"
	"github.com/KingKord/bookings/internal/models"
	"time"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into a database
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {

	// if the room id is 2, then fail; otherwise, pass
	if res.RoomID == 2 {
		return 0, errors.New("some error)")
	}
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	if r.RoomID == 1000 {
		return errors.New("some error)")

	}
	return nil
}

// SearchAvailabilityByDates returns true if availability exist for roomID, and false if no availability
func (m *testDBRepo) SearchAvailabilityByDates(start, end time.Time, roomID int) (bool, error) {
	if start.Year() == 3000 {
		return true, nil
	} else if roomID > 2 {
		return false, errors.New("some error)")
	}
	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms, if any, for given date range
func (m testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {

	var rooms []models.Room
	if start.Year() == 3000 {
		room := models.Room{
			ID:        1,
			RoomName:  "General's Quarters",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		rooms = append(rooms, room)
		return rooms, nil
	} else if start.Year() == 3002 {
		return rooms, errors.New("some error)")
	}
	return rooms, nil

}

// GetRoomByID gets a room by ID
func (m testDBRepo) GetRoomByID(id int) (models.Room, error) {

	var room models.Room
	if id > 2 {
		return room, errors.New("some error")
	}

	return room, nil
}

func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *testDBRepo) UpdateUser(u models.User) error {
	//TODO implement me
	panic("implement me")
}

func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	//TODO implement me
	panic("implement me")
}
