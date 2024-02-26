package api

import (
	"errors"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	user_id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("get-photo: error converting id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ! Login
	requestingUserId_str, err := extractBearerToken(r, w)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(err).Error("get-photo: Error")
		return
	}
	requestingUserId, err := strconv.Atoi(requestingUserId_str)
	if err != nil {
		ctx.Logger.WithError(err).Error("get-photo: error converting requesingUserId")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ! Controllo se bannato
	banned, err := rt.db.CheckBan(
		User{User_id: requestingUserId}.toDataBase(),
		User{User_id: user_id}.toDataBase())
	if err != nil {
		ctx.Logger.WithError(err).Error("get-photo: Error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		ctx.Logger.WithError(errors.New("l'utente è bloccato")).Error("get-profile: Impossibile eseguire l'operazione")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	file := ps.ByName("photo_id") + ".jpg"
	path := filepath.Join("/tmp", "media", ps.ByName("id"), "photos", file)
	http.ServeFile(w, r, path)
	w.WriteHeader(http.StatusOK)

}
