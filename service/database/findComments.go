package database

func (db *appdbimpl) FindComments(photo Photo_id) ([]Comment, error) {
	queryRes, err := db.c.Query("SELECT Comment.comment_id, Comment.photo_id, Comment.txt, Users.user_id, Users.nickname, Comment.timestamp	FROM Comment INNER JOIN Users ON Comment.owner = Users.user_id WHERE Comment.photo_id= ?;", photo.Photo_id)
	// ! Controllo degli errori
	if err != nil {
		queryRes.Close()
		return nil, err
	}
	if queryRes.Err() != nil {
		return nil, queryRes.Err()
	}

	var comments []Comment // Creo un'array di Commenti
	for queryRes.Next() {
		var comment Comment
		err = queryRes.Scan(&comment.Comment_identifier, &comment.Photo_id, &comment.Comment_string, &comment.Owner.User_id, &comment.Owner.Nickname, &comment.Timestamp) // Prendiamo i risultati ottenuti e li inseriamo nel singolo utente
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment) // Appendo gli comment
	}
	return comments, nil
}
