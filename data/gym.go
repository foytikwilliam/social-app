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
	Zipcode    string
	Email      string
	Website    string
	Created_At time.Time
}

type Review struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	GymId     int
	Rating    int
	Date      time.Time
	CreatedAt time.Time
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
	err = Db.QueryRow("SELECT id, uuid, name, address, city, state, zipcode, county, phone, email, website, created_at FROM gyms WHERE uuid = $1", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Name, &conv.Address, &conv.City, &conv.State, &conv.Zipcode, &conv.County, &conv.Phone, &conv.Email, &conv.Website, &conv.Created_At)

	return
}

func GymByID(id int) (Gym, error) {
	var gym Gym
	err := Db.QueryRow("SELECT id, uuid, name, address, city, state, county, phone, email, website, created_at FROM gyms WHERE id = $1", id).
		Scan(&gym.Id, &gym.Uuid, &gym.Name, &gym.Address, &gym.City, &gym.State, &gym.County, &gym.Phone, &gym.Email, &gym.Website, &gym.Created_At)

	if err != nil {
		return Gym{}, err
	}

	return gym, nil
}

func (gym *Gym) Reviews() (reviews []Review, err error) {

	rows, err := Db.Query("SELECT id, uuuid, body, user_id, gym_id, rating, date, created_at FROM reviews where gym_id", gym.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		review := Review{}
		if err = rows.Scan(&review.Id, &review.Uuid, &review.Body, &review.UserId, &review.GymId, &review.Rating, &review.Date, &review.CreatedAt); err != nil {
			return
		}
		reviews = append(reviews, review)
	}
	rows.Close()
	return
}

func (user *User) CreateReview(conv Gym, body string) (review Review, err error) {
	// Assuming gym_id is conv.Id
	statement := "INSERT INTO reviews (uuid, body, user_id, gym_id, rating, date, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, uuid, body, user_id, gym_id, rating, date, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	// Assuming you have the rating value
	rating := 0

	// Assuming you have the date value
	date := time.Now()

	// Use conv.Id as the gym_id
	err = stmt.QueryRow(createUUID(), body, user.Id, conv.Id, rating, date, time.Now()).
		Scan(&review.Id, &review.Uuid, &review.Body, &review.UserId, &review.GymId, &review.Rating, &review.Date, &review.CreatedAt)

	return
}

func ReviewByUUID(uuid string) (conv Review, err error) {
	conv = Review{}
	err = Db.QueryRow("SELECT id, uuid, body, user_id, gym_id, rating, date, created_at FROM reviews WHERE uuid = $1", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Body, &conv.UserId, &conv.CreatedAt)
	return
}

func ReviewsByGymID(GymId int) ([]Review, error) {
	rows, err := Db.Query("SELECT id, uuid, body, user_id, gym_id, rating, date, created_at FROM reviews WHERE gym_id = $1", GymId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []Review

	for rows.Next() {
		var review Review
		if err := rows.Scan(&review.Id, &review.Uuid, &review.Body, &review.UserId, &review.GymId, &review.Rating, &review.Date, &review.CreatedAt); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reviews, nil
}
