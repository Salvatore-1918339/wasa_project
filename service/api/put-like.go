package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// ! Login
	requestingUserId_str, err := extractBearerToken(r, w)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(err).Error("put-like: Error extract token")
		return
	}
	requestingUserId, err := strconv.Atoi(requestingUserId_str)
	if err != nil {
		ctx.Logger.WithError(err).Error("put-like: Error conversion")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ! Controllo se l'utente è bloccato
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

	// ! Controllo che non ci sia già un like
	exist, err := rt.db.Checklike(photo_id, requestingUserId)
	if err != nil {
		ctx.Logger.WithError(err).Error("put-like : error executing Checkfollow")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if exist {
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
