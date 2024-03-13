package database

func (db *appdbimpl) SearchUser(queryStr string) ([]User, error) {

	queryRes, err := db.c.Query("SELECT * FROM Users WHERE nickname LIKE ? ", "%"+queryStr+"%")
	// ! Controllo degli errori
	if err != nil {
		queryRes.Close()
		return nil, err
	}
	if queryRes.Err() != nil {
		return nil, queryRes.Err()
	}

	var users []User // Creo un'array di Utenti
	for queryRes.Next() {
		var user User
		err = queryRes.Scan(&user.User_id, &user.Nickname) // Prendiamo i risultati ottenuti e li inseriamo nel singolo utente
		if err != nil {
			return nil, err
		}
		users = append(users, user) // Appendo gli User
	}
	return users, nil

}
