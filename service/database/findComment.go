package database

import "fmt"

func (db *appdbimpl) FindComment(Comment_id int) (Comment, error) {
	queryRes, err := db.c.Query("SELECT Comment.comment_id, Comment.photo_id, Comment.txt, Users.user_id, Users.nickname, Comment.timestamp	FROM Comment INNER JOIN Users ON Comment.owner = Users.user_id WHERE Comment.comment_id = ?;", Comment_id)
	var comment Comment
	if err != nil {
		return comment, err
	}

	for queryRes.Next() {
		err = queryRes.Scan(&comment.Comment_identifier, &comment.Photo_id, &comment.Comment_string, &comment.Owner.User_id, &comment.Owner.Nickname, &comment.Timestamp) // Prendiamo i risultati ottenuti e li inseriamo nel singolo utente
		if err != nil {
			return comment, err
		}
	}
	fmt.Print("COMMENTO", comment)
	return comment, nil

}
