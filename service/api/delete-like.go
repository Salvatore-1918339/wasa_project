package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	user_id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-like: Error conversion")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	photo_id, err := strconv.Atoi(ps.ByName("photo_id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-like: Error conversion")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	like_id, err := strconv.Atoi(ps.ByName("like_id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-like: Error conversion")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ! Login
	requestingUserId_str, err := extractBearerToken(r, w)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(err).Error("commentPhoto: Error")
		return
	}
	requestingUserId, _ := strconv.Atoi(requestingUserId_str)
	// ? Controllo che l'utente che toglie like sia quello auth
	if like_id != requestingUserId {
		ctx.Logger.WithError(errors.New("non sei autorizzato ad eseguire questa operazione")).Error("Impossibile eseguire l'operazione")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// ! Controllo se l'utente è bloccato
	banned, err := rt.db.CheckBan(
		User{User_id: requestingUserId}.toDataBase(),
		User{User_id: user_id}.toDataBase())
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-comment: Error in CheckBan")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		ctx.Logger.WithError(errors.New("l'utente è bloccato")).Error("impossibile eseguire l'operazione")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// ! Tolgo like
	err = rt.db.DeleteLike(
		Photo_id{Photo_id: photo_id}.toDataBase(),
		User{User_id: requestingUserId}.toDataBase())
	if err != nil {
		ctx.Logger.WithError(err).Error("error in DeleteLIke: error executing query ")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
