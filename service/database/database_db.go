/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	CreateUser(User) error
	CreatePhoto(int, string) (int, error)
	CreateComment(User, Photo_id, string, string) (int, error)
	ChangeNickname(User, string) error
	BanUser(User, User) error
	FollowUser(User, User) error

	//GET
	GetNickname(int) (string, error)
	GetNumberFollower(int) (int, error)
	GetNumberFollowed(int) (int, error)
	SearchUser(string) ([]User, error)

	//DELETE
	DeletePhoto(int, int) error
	UnBanUser(int, int) error
	UnFollowUser(User, User) error
	DeleteComment(Comment_id) error
	DeleteLike(Photo_id, User) error

	//SPECIALE
	FindUserId(User) (int, error) // Trova nel DB l'utente per farmi tornare il suo ID
	Checkfollow(int, int) (bool, error)
	Checklike(int, int) (bool, error)
	CheckBan(User, User) (bool, error)
	FindCommentOwner(Comment_id) (int, error)
	GetFollowing(User) ([]int, error)
	FindPhotos(User) ([]Complete_Photo, error)
	FindLikes(Photo_id) ([]User_id, error)
	FindComment(Photo_id) ([]Comment, error)
	PutLike(Photo_id, User) error

	// Ping check if the DB is available or not
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// attiva le foreign keys
	_, errPramga := db.Exec(`PRAGMA foreign_keys= ON`)
	if errPramga != nil {
		return nil, fmt.Errorf("error setting pragmas: %w", errPramga)
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='example_table';`).Scan(&tableName)

	if errors.Is(err, sql.ErrNoRows) {
		// Creazione DB per gli Users se non esiste
		users := `CREATE TABLE IF NOT EXISTS Users 
										(user_id INTEGER NOT NULL, 
										nickname VARCHAR(25) NOT NULL UNIQUE,
										PRIMARY KEY("user_id" AUTOINCREMENT));`
		_, err = db.Exec(users)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: Users %w", err)
		}
		//Creazione DB per le Photo degli Users
		photos := `CREATE TABLE IF NOT EXISTS Photo(
					photo_id	INTEGER NOT NULL UNIQUE,
					owner	INTEGER NOT NULL,
					timestamp	DATETIME NOT NULL,
					PRIMARY KEY(photo_id AUTOINCREMENT),
					FOREIGN KEY(owner) REFERENCES Users ( user_id ));`

		_, err = db.Exec(photos)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: Photo %w", err)
		}
		comments := `CREATE TABLE IF NOT EXISTS Comment (
						comment_id	INTEGER NOT NULL UNIQUE,
						photo_id	INTEGER NOT NULL,
						owner	INTEGER NOT NULL,
						txt	VARCHAR(300) NOT NULL,
						timestamp	DATETIME NOT NULL,
						FOREIGN KEY(photo_id) REFERENCES Photo (photo_id),
						FOREIGN KEY(owner) REFERENCES Users (user_id),
						PRIMARY KEY(comment_id AUTOINCREMENT));`
		_, err = db.Exec(comments)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: Comment %w", err)
		}

		// ? banner : Utente che ha bannato user_id
		banned := `CREATE TABLE IF NOT EXISTS Banned(
				banner INTEGER NOT NULL,
				user_id INTEGER NOT NULL,
				PRIMARY KEY(banner, user_id),
				FOREIGN KEY("banner") REFERENCES Users (user_id),
				FOREIGN KEY("user_id") REFERENCES Users (user_id));`

		_, err = db.Exec(banned)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: Banned %w", err)
		}

		// ? follower : Utenti che seguono User_id
		followers := `CREATE TABLE IF NOT EXISTS Followers(
						follower INTEGER NOT NULL,	
						user_id INTEGERE NOT NULL,
						PRIMARY KEY(follower, user_id),
						FOREIGN KEY("follower") REFERENCES Users (user_id),
						FOREIGN KEY("user_id") REFERENCES Users (user_id));`

		_, err = db.Exec(followers)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: Followers %w", err)
		}

		likes := `CREATE TABLE IF NOT EXISTS Likes(
				photo_id INTEGER NOT NULL,	
				user_id INTEGER NOT NULL,
				PRIMARY KEY(photo_id, user_id),
				FOREIGN KEY("photo_id") REFERENCES Photo (photo_id),
				FOREIGN KEY("user_id") REFERENCES Users (user_id));`
		_, err = db.Exec(likes)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: Likes %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
