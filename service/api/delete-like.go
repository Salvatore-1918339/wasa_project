package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userId_str := ps.ByName("id")
	photoId_str := ps.ByName("photo_id")
	likeId_str := ps.ByName("like_id")
	requestingUserId_str, err := extractBearerToken(r, w)

	// Controllo errore dall'estrazione del TOKEN
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(err).Error("commentPhoto: Error")
		return
	}

	// Converto i Valori ID in Interi
	user_id, _ := strconv.Atoi(userId_str)
	photo_id, _ := strconv.Atoi(photoId_str)
	like_id, _ := strconv.Atoi(likeId_str)
	requestingUserId, _ := strconv.Atoi(requestingUserId_str)

	// Controllo che l'utente che mette like sia quello auth
	if like_id != requestingUserId {
		ctx.Logger.WithError(errors.New("Non sei autorizzato ad eseguire questa operazione")).Error("Impossibile eseguire l'operazione")
		w.WriteHeader(http.StatusForbidden) //errore 403
		return
	}

	// Controllo se l'utente è bloccato
	banned, err := rt.db.CheckBan(
		User{User_id: requestingUserId}.toDataBase(),
		User{User_id: user_id}.toDataBase())

	if banned {
		ctx.Logger.WithError(errors.New("L'utente è bloccato")).Error("Impossibile eseguire l'operazione")
		w.WriteHeader(http.StatusForbidden) //errore 403
		return
	}

	// ! Tolgo like
	err = rt.db.DeleteLike(
		Photo_id{Photo_id: photo_id}.toDataBase(),
		User{User_id: requestingUserId}.toDataBase())
	if err != nil {
		ctx.Logger.WithError(err).Error("Error in DeleteLIke: error executing query ")
		w.WriteHeader(http.StatusInternalServerError) //errore 403
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204
}
