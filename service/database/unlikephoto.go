package database

func (db *appdbimpl) DeleteLike(photo_id Photo_id, user User) error {
	_, err := db.c.Exec("DELETE FROM Likes WHERE photo_id=? AND user_id=?", photo_id.Photo_id, user.User_id)
	if err != nil {
		return err
	}
	return nil
}
