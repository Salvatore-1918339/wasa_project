package api

import (
	"errors"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	user_id, _ := strconv.Atoi(ps.ByName("id"))
	requestingUserId_str, err := extractBearerToken(r, w)
	// Controllo errore dall'estrazione del TOKEN
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(err).Error("get-photo: Error")
		return
	}

	requestingUserId, _ := strconv.Atoi(requestingUserId_str)

	// ! Controllo se bannato
	banned, err := rt.db.CheckBan(
		User{User_id: requestingUserId}.toDataBase(),
		User{User_id: user_id}.toDataBase())
	if err != nil {
		ctx.Logger.WithError(err).Error("get-photo: Error")
		w.WriteHeader(http.StatusForbidden) //errore 403
		return
	}
	if banned {
		ctx.Logger.WithError(errors.New("l'utente è bloccato")).Error("get-profile: Impossibile eseguire l'operazione")
		w.WriteHeader(http.StatusForbidden) //errore 403
		return
	}

	file := ps.ByName("photo_id") + ".jpg"
	path := filepath.Join("/tmp", "media", ps.ByName("id"), "photos", file)
	http.ServeFile(w, r, path)

}
