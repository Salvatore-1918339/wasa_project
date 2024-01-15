package api

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//Prendo user_id e photo_id forniti in URL
	urlArry := strings.Split(r.URL.Path, "/")
	u_id_temp, p_id_temp := urlArry[2], urlArry[4]

	user_id, err1 := strconv.Atoi(u_id_temp)
	photo_id, err2 := strconv.Atoi(p_id_temp)

	switch {
	case err1 != nil:
		ctx.Logger.WithError(err1).Error("delete-photo: error with string conversion 1")
		return
		// Gestisci l'errore di user_id come preferisci

	case err2 != nil:
		ctx.Logger.WithError(err2).Error("delete-photo: error with string conversion 2")
		return
		// Gestisci l'errore di photo_id come preferisci

	default:
		// Utilizza gli interi ottenuti
		// fmt.Printf("user_id: %d\n", user_id)
		// fmt.Printf("photo_id: %d\n", photo_id)

	}

	// elimino la foto dal DB
	err := rt.db.DeletePhoto(user_id, photo_id)
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-photo: error deleting photo from DB ")
		return
	}

	//Elimino la foto dalla directory
	path := filepath.Join("/tmp", "media", u_id_temp, "photos") // /tmp/media/:id/photos
	filepath := filepath.Join(path, p_id_temp+".jpg")
	// fmt.Print(filepath)

	err = os.Remove(filepath)
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-photo: error deleting photo from OS ")

	}

}
