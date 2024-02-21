package api

import (
	"database/sql"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	user_id, err1 := strconv.Atoi(ps.ByName("id"))
	if err1 != nil {
		ctx.Logger.WithError(err1).Error("delete-photo: error converting id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	photo_id, err2 := strconv.Atoi(ps.ByName("photo_id"))
	if err2 != nil {
		ctx.Logger.WithError(err2).Error("delete-photo: error converting photo_id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestingUserId_str, err := extractBearerToken(r, w)
	// Controllo errore dall'estrazione del TOKEN
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(err).Error("delete-photo: Error extractBearerToken")
		return
	}

	// ! login

	// elimino la foto dal DB
	err = rt.db.DeletePhoto(user_id, photo_id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.Logger.WithError(err).Error("delete-photo: error deleting photo from DB - not found")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		ctx.Logger.WithError(err).Error("delete-photo: error deleting photo from DB ")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Elimino la foto dalla directory
	path := filepath.Join("/tmp", "media", ps.ByName("id"), "photos") // /tmp/media/:id/photos
	filepath := filepath.Join(path, ps.ByName("photo_id")+".jpg")
	// fmt.Print(filepath)

	err = os.Remove(filepath)
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-photo: error deleting photo from OS ")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
