package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// getContextReply is an example of HTTP endpoint that returns "Hello World!" as a plain text. The signature of this
// handler accepts a reqcontext.RequestContext (see httpRouterHandler).
func (rt *_router) doLoginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("Content-Type", "application/json")

	// Prende Utente passato in r
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !validIdentifier(user.Nickname) {
		ctx.Logger.WithError(err).Error("session: Can't Create a User. User nickname not Valid. <<")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ! Controlla che l'utente esista
	temp_user, err := rt.db.CheckUser(user.toDataBase())
	if err != nil && err != sql.ErrNoRows {
		ctx.Logger.WithError(err).Error("session: Error in CheckUser")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if temp_user.User_id != 0 { // ? Ho trovato l'user quindi assegno i valori
		user.User_id = temp_user.User_id
		user.Nickname = temp_user.Nickname

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("session: can't create response json")
			return
		}
		return
	}

	// ! Controlla se l'utente è già presebnte nel DB
	err = rt.db.CreateUser(user.toDataBase())
	if err != nil {
		w.WriteHeader(http.StatusOK) // L'utente esisteva già nel DB
		_ = json.NewEncoder(w).Encode(user)
		return
	}

	// ! prendo l'ID Salvato nel DB
	id, err_fu := rt.db.FindUserId(user.toDataBase())
	if err_fu != nil {
		ctx.Logger.WithError(err_fu).Error("session: Error in FindUserId()")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.User_id = id // Cambio l'ID dentro la variabile user

	// ! Creo la cartella del nuovo Utente
	CreateFolder(strconv.Itoa(id), ctx)

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("session: can't create response json")
		return
	}
}

func CreateFolder(id string, ctx reqcontext.RequestContext) {
	err := os.Mkdir("/tmp/media", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		ctx.Logger.WithError(err).Error("Session : Error creating /tmp/media")
		return
	}
	path := filepath.Join("/tmp/media", id)
	err = os.Mkdir(path, os.ModePerm)
	if err != nil && os.IsExist(err) {
		ctx.Logger.WithError(err).Error("Session : /tmp/media/{id} it already exists")
	} else if err != nil {
		ctx.Logger.WithError(err).Error("Session : Error creating /tmp/media/{id}")
		return
	}
	path = filepath.Join(path, "photos")
	err = os.Mkdir(path, os.ModePerm)
	if err != nil && os.IsExist(err) {
		ctx.Logger.WithError(err).Error("Session : /tmp/media/{id}/photo/ it already exists")
	} else if err != nil {
		ctx.Logger.WithError(err).Error("Session : Error creating /tmp/media/{id}/photo/")
		return
	}
}
