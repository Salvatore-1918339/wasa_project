package database

func (db *appdbimpl) BanUser(banner User, banned User) error {

	_, err1 := db.c.Exec("INSERT INTO banned (banner,banned) VALUES (?, ?)", banner.User_id, banned.User_id)
	if err1 != nil {
		return err1
	}
	return nil
}
