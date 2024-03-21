package api

import (
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
	/*
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
	*/

	photo_id_str := ps.ByName("photo_id")
	photo_id, err := strconv.Atoi(photo_id_str)
	if err != nil {
		ctx.Logger.WithError(err).Error("get-photo : error converting photo_id_str")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	path := filepath.Join("/tmp", "media", ps.ByName("id"), "photos", photo_id_str+".png")

	exist, err := rt.db.CheckPhoto(user_id, photo_id)
	if err != nil {
		ctx.Logger.WithError(err).Error("get-photo : err in checkPhoto query")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !exist {
		ctx.Logger.WithError(err).Error("get-photo : err in checkPhoto query")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, path)

}
