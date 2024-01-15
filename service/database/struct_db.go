package database

import (
	"time"
)

type Comment struct {
	Comment_string     string    `json:"comment_string"`
	Photo_id           int       `json:"photo_id"`
	Owner              string    `json:"owner"`
	Timestamp          time.Time `json:"timestamp"`
	Comment_identifier string    `json:"comment_identifier"`
}

/*
	Liker - Utente al quale piace una
*/
type Like struct {
	User_Liker int
	Photo_id   int
}

type User_id struct {
	User_id int `json:"user_id"`
}

//Struttura User
type User struct {
	User_id  int    `json:"user_id"`
	Nickname string `json:"nickname"`
}
type Photo_id struct {
	Photo_id int `json:"photo_id"`
}

type Comment_id struct {
	Comment_id int `json:"comment_identifier"`
}

//Struttura Photo
type Complete_Photo struct {
	Photo_id  int       `json:"photo_id"`
	Owner     int       `json:"owner"`
	Timestamp time.Time `json:"timestamp"`
	Comments  []Comment `json:"comments"`
	Likes     []User_id `json:"likes"`
}
