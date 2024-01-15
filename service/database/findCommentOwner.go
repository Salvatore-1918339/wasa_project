package database

func (db *appdbimpl) FindCommentOwner(comment Comment_id) (int, error) {
	var owner int
	err := db.c.QueryRow("SELECT owner FROM Comment WHERE comment_id=?", comment.Comment_id).Scan(&owner)
	if err != nil {
		return -1, err
	}
	return owner, nil
}
