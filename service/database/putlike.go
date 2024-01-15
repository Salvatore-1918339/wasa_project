package database

func (db *appdbimpl) PutLike(photo_id Photo_id, user User) error {
	_, err := db.c.Exec("INSERT INTO Likes (photo_id, user_id) VALUES (?,?)", photo_id.Photo_id, user.User_id)
	if err != nil {
		return err
	}
	return nil
}
