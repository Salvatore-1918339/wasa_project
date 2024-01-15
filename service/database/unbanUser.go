package database

func (db *appdbimpl) UnBanUser(banner int, banned int) error {

	_, err := db.c.Exec("DELETE FROM Banned WHERE banner=? AND banned=?", banner, banned)
	if err != nil {
		return err
	}
	return nil
}
