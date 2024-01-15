package database

func (db *appdbimpl) GetFollowing(user User) ([]int, error) {

	// follower : ossia User è un follower degli utenti che tornano in queryRes
	queryRes, err := db.c.Query("SELECT * FROM Followers WHERE follower=? ", user.User_id)
	if err != nil {
		return nil, err
	}

	var users_followed []int // Creo un'array di Utenti
	for queryRes.Next() {
		var user_follower_id, user_id int
		err = queryRes.Scan(&user_follower_id, &user_id) // Prendiamo i risultati ottenuti e li inseriamo nel singolo utente
		if err != nil {
			return nil, err
		}
		users_followed = append(users_followed, user_id) // Appendo gli User seguiti dal follower
	}
	return users_followed, nil
}
