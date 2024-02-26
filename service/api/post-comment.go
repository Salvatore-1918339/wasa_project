package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	photoOwnerId, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("comment-photo: error converting id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	photo_id, err := strconv.Atoi(ps.ByName("photo_id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("comment-photo: error converting photo_id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ! login
	requestingUserId_str, err := extractBearerToken(r, w)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(err).Error("comment-Photo: Error")
		return
	}
	requestingUserId, err := strconv.Atoi(requestingUserId_str)
	if err != nil {
		ctx.Logger.WithError(err).Error("comment-photo: error converting requestingUser_id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ! Controllo se l'utente è bloccato
	banned, err := rt.db.CheckBan(
		User{User_id: requestingUserId}.toDataBase(),
		User{User_id: photoOwnerId}.toDataBase())
	if err != nil {
		ctx.Logger.WithError(err).Error("post-comment: Error check ban query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		ctx.Logger.WithError(errors.New("the user is banned")).Error("post-comment: Error")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// ! Controllo se il commento è della lunghezza valida
	var comment Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("post-comment/Decode: failed to decode request body json")
		return
	}
	if !validCommentString(comment) {
		// ? Dimensioni della foto non idonee
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ! Aggiungo il commento nel DB
	currentTime := time.Now()
	datetime := currentTime.Format("2006-01-02 15:04:05")

	comment_id, err := rt.db.CreateComment(
		User{User_id: requestingUserId}.toDataBase(),
		Photo_id{Photo_id: photo_id}.toDataBase(),
		comment.Comment_string,
		datetime)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("post-Comment: failed executing query CreateComment ")
		return
	}

	w.WriteHeader(http.StatusCreated)
	// ! RITORNO L'ID DELL'COMMENTO
	err = json.NewEncoder(w).Encode(Comment_id{Comment_id: comment_id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("post-comment/Encode: failed convert photo_id to int64")
		return
	}
}
