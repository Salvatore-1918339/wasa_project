package api

import (
	"encoding/json"
	"fmt"
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
	fmt.Print("\n#!# UTENTE PASSATO : ", user.Nickname)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !validIdentifier(user.Nickname) {
		ctx.Logger.WithError(err).Error("session: Can't Create a User. User nickname not Valid. <<")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ! Controlla se l'utente è già presebnte nel DB
	err = rt.db.CreateUser(user.toDataBase())

	// ! prendo l'ID Salvato nel DB
	id, err_fu := rt.db.FindUserId(user.toDataBase())
	if err_fu != nil {
		ctx.Logger.WithError(err_fu).Error("session: Error in FindUserId()")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.User_id = id // Cambio l'ID dentro la variabile user

	if err != nil { // ? Se err di CreateUser() != nil
		w.WriteHeader(http.StatusCreated) // L'utente esisteva già nel DB
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("session: can't create response json")
		}
		return
	}

	// ! Creo la cartella del nuovo Utente
	CreateFolder(strconv.Itoa(id), ctx)

	// ! 201
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("session: can't create response json")
		return
	}
}

func CreateFolder(id string, ctx reqcontext.RequestContext) error {
	// VEDI DI TROVARE UNA SOLUZIONE MIGLIORE PER RISOLVERE QUESTO PROBLEMA
	err := os.Mkdir("/tmp/media", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		// Gestisci altri errori
		fmt.Printf("Errore durante la creazione della directory: %v\n", err)
		return err
	}
	path := filepath.Join("/tmp/media", id)
	err = os.Mkdir(path, os.ModePerm)
	if err != nil && os.IsExist(err) {
		fmt.Print("La directory esiste già. SEZ[1]\n")
	} else if err != nil {
		// Gestisci altri errori
		fmt.Printf("Errore durante la creazione della directory: %v\n", err)
		return err
	}
	path = filepath.Join(path, "photos")
	err = os.Mkdir(path, os.ModePerm)
	if err != nil && os.IsExist(err) {
		fmt.Print("La directory esiste già. SEZ[2]\n")
	} else if err != nil {
		// Gestisci altri errori
		fmt.Printf("Errore durante la creazione della directory: %v\n", err)
		return err
	}

	return nil

}
