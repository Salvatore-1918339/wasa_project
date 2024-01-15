package database

func (db *appdbimpl) FindComment(photo Photo_id) ([]Comment, error) {
	queryRes, err := db.c.Query("SELECT * FROM Comment WHERE photo_id=?", photo.Photo_id)
	if err != nil {
		return nil, err
	}

	var comments []Comment // Creo un'array di Commenti
	for queryRes.Next() {
		var comment Comment
		err = queryRes.Scan(&comment.Comment_identifier, &comment.Photo_id, &comment.Owner, &comment.Comment_string, &comment.Timestamp) // Prendiamo i risultati ottenuti e li inseriamo nel singolo utente
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment) // Appendo gli comment
	}
	return comments, nil
}
