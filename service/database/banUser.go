package database

func (db *appdbimpl) BanUser(banner User, banned User) error {

	_, err1 := db.c.Exec("INSERT INTO banned (banner,user_id) VALUES (?, ?)", banner.User_id, banned.User_id)
	return err1
}
