package api

import (
	"time"

	"github.com/Salvatore-1918339/wasa_project/service/database"
)

type Comment_id struct {
	Comment_id int `json:"comment_identifier"`
}

type Comment struct {
	Comment_string     string    `json:"comment_string"`
	Photo_id           int       `json:"photo_id"`
	Owner              int       `json:"comment_owner"`
	Timestamp          time.Time `json:"comment_timespamp"`
	Comment_identifier int       `json:"comment_identifier"`
}

type Nickname struct {
	Nickname string `json:"nickname"`
}

type id_user struct {
	User_id string `json:"user_id"`
}

/*
	Liker - Utente al quale piace una
*/
type Like struct {
	User_Liker int
	Photo_id   int
}

//Struttura User
type User struct {
	User_id  int    `json:"user_id"`
	Nickname string `json:"nickname"`
}

//Struttura Photo
type Complete_Photo struct {
	Photo_id  int                `json:"photo_id"`
	Owner     User               `json:"owner"`
	Timestamp time.Time          `json:"timestamp"`
	Comments  []Comment          `json:"comments"`
	Likes     []database.User_id `json:"likes"`
}

type Profile struct {
	User_id    int                       `json:"user_id"`
	Nickname   string                    `json:"nickname"`
	N_Follower int                       `json:"n_follower"`
	Followers  []int                     `json:"followers"`
	Following  int                       `json:"n_following"`
	Posts      []database.Complete_Photo `json:"posts"`
}

type Photo_id struct {
	Photo_id int `json:"photo_id"`
}

func (u User) toDataBase() database.User {
	return database.User{
		User_id:  u.User_id,
		Nickname: u.Nickname,
	}
}

func (p Photo_id) toDataBase() database.Photo_id {
	return database.Photo_id{
		Photo_id: p.Photo_id,
	}
}

func (p Comment_id) toDataBase() database.Comment_id {
	return database.Comment_id{
		Comment_id: p.Comment_id,
	}
}
