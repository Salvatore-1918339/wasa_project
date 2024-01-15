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

	w.Header().Set("Content-Type", "application/json")
	photoOwnerId_str := ps.ByName("id")
	photoId_str := ps.ByName("photo_id")
	requestingUserId_str, err := extractBearerToken(r, w)

	// Controllo errore dall'estrazione del TOKEN
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(err).Error("commentPhoto: Error")
		return
	}

	// Converto i Valori ID in Interi
	photoOwnerId, _ := strconv.Atoi(photoOwnerId_str)
	photo_id, _ := strconv.Atoi(photoId_str)
	requestingUserId, _ := strconv.Atoi(requestingUserId_str)

	// Controllo se l'utente è bloccato
	banned, err := rt.db.CheckBan(
		User{User_id: requestingUserId}.toDataBase(),
		User{User_id: photoOwnerId}.toDataBase())

	if banned {
		ctx.Logger.WithError(errors.New("L'utente è bloccato")).Error("Impossibile eseguire l'operazione")
		w.WriteHeader(http.StatusForbidden) //errore 403
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