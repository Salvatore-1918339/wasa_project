package database

func (db *appdbimpl) DeleteComment(comment Comment_id) error {

	_, err := db.c.Exec("DELETE FROM Comment WHERE comment_id=?", comment.Comment_id)
	if err != nil {
		return err
	}
	return nil
}
