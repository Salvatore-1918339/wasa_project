package database

func (db *appdbimpl) DeletePhoto(user_id int, photo_id int) error {
	//Elimino la Photo dal DB
	_, err := db.c.Exec("DELETE FROM Photo WHERE photo_id=?", photo_id)
	if err != nil {
		return err
	}

	return nil

}
