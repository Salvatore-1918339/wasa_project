package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	requestingUserId_str, err := extractBearerToken(r, w)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("put-like: Error extract token")
		return
	}

	// Converto i Valori ID in Interi
	user_id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("put-like: Error conversion")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	photo_id, err := strconv.Atoi(ps.ByName("photo_id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("put-like: Error conversion")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	/*
		like_id, _ := strconv.Atoi(ps.ByName("like_id"))
		if err != nil {
			ctx.Logger.WithError(err).Error("put-like: Error conversion")
			w.WriteHeader(http.StatusBadRequest)
			return
		}*/
	requestingUserId, _ := strconv.Atoi(requestingUserId_str)

	// Controllo se l'utente è bloccato
	banned, err := rt.db.CheckBan(
		User{User_id: requestingUserId}.toDataBase(),
		User{User_id: user_id}.toDataBase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-like: Error check ban query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		ctx.Logger.WithError(errors.New("the User is banned")).Error("put-like: Error")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// ! Metto like
	err = rt.db.PutLike(
		Photo_id{Photo_id: photo_id}.toDataBase(),
		User{User_id: requestingUserId}.toDataBase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-like: error executing query ")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
