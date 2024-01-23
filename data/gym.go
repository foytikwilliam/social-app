package data

import (
	"fmt"
	"time"
)

type Gym struct {
	Id         int
	Uuid       string
	Name       string
	Address    string
	City       string
	State      string
	County     string
	Phone      string
	Email      string
	Website    string
	Created_At time.Time
}

type Review struct {
	Id         int
	Uuid       string
	Body       string
	User_Id    int
	Gym_id     int
	Rating     int
	Date       string
	Created_At time.Time
}

func GetGyms() ([]Gym, error) {

	rows, err := Db.Query("SELECT id, uuid, name, address, city, state, county, phone, email, website, created_at FROM gyms")

	if err != nil {
		return nil, fmt.Errorf("failed to query gyms: %v", err)
	}

	var gyms []Gym

	for rows.Next() {
		var gym Gym
		err := rows.Scan(&gym.Id, &gym.Uuid, &gym.Name, &gym.Address, &gym.City, &gym.State, &gym.County, &gym.Phone, &gym.Email, &gym.Website, &gym.Created_At)
		if err != nil {
			return nil, fmt.Errorf("failed to scan gym row: %v", err)
		}
		gyms = append(gyms, gym)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during iteration through gym rows: %v", err)
	}

	return gyms, nil
}

// GetReviewsForGym fetches reviews for a specific gym from the database.
//func GetReviewsForGym(gymID int) ([]Review, error) {
// Implement logic to fetch reviews for the specified gym from the database
// Return a slice of Review objects
//}

func GymByUUID(uuid string) (conv Gym, err error) {
	conv = Gym{}
	err = Db.QueryRow("SELECT id, uuid, name, address, city, state, county, phone, email, website, created_at FROM gyms WHERE uuid = $1", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Name, &conv.Address, &conv.City, &conv.State, &conv.County, &conv.Phone, &conv.Email, &conv.Website, &conv.Created_At)
	return
}
