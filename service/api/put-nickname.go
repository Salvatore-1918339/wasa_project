package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) PutNickname(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	user_id_str := ps.ByName("id")
	requestingUserId_str, err := extractBearerToken(r, w)

	// Controllo errore dall'estrazione del TOKEN
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(err).Error("GetNIcknameHandler: Error")
		return
	}

	user_id, _ := strconv.Atoi(user_id_str)
	requestingUserId, _ := strconv.Atoi(requestingUserId_str)

	// ! Non è l'utente Loggato che sta cercando di effetturare questa operazione
	if user_id != requestingUserId {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(errors.New("you aren't allowed to  use this operation")).Error("GetNIcknameHandler: Error")
		return
	}

	// Prendo il Body
	var newNickname Nickname
	err = json.NewDecoder(r.Body).Decode(&newNickname)
	if err != nil {
		ctx.Logger.WithError(err).Error("update-nickname: error decoding json")
		w.WriteHeader(http.StatusBadRequest) // ! Errore 400
		return
	} else if !validIdentifier(newNickname.Nickname) {
		ctx.Logger.WithError(errors.New("nickname provided does not meet the necessary conditions")).Error("update-nickname: Error in Nickname section")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ! richiamo changeNickname del DB
	err = rt.db.ChangeNickname(User{User_id: user_id}.toDataBase(), newNickname.Nickname)
	if err != nil {
		ctx.Logger.WithError(err).Error("update-nickname: Error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// risponde con 204 http status
	w.WriteHeader(http.StatusNoContent)
}
