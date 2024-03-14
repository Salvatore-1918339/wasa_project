package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	user_id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("put-nickname: error converting id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ! Login
	requestingUserId_str, err := extractBearerToken(r, w)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(err).Error("put-nickname: Error")
		return
	}
	requestingUserId, err := strconv.Atoi(requestingUserId_str)
	if err != nil {
		ctx.Logger.WithError(err).Error("put-nickname: error converting requesingUserId")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user_id != requestingUserId {
		w.WriteHeader(http.StatusUnauthorized)
		ctx.Logger.WithError(errors.New("you aren't allowed to use this operation")).Error("put-Nickname: Error")
		return
	}

	// Prendo il Body
	var newNickname Nickname
	err = json.NewDecoder(r.Body).Decode(&newNickname)
	if err != nil {
		ctx.Logger.WithError(err).Error("put-nickname: error decoding json")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !validIdentifier(newNickname.Nickname) {
		ctx.Logger.WithError(errors.New("nickname provided does not meet the necessary conditions")).Error("put-nickname: Error in Nickname section")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// ! Controllo che il nickname sia libero
	exist, err := rt.db.SearchNickname(newNickname.Nickname)
	if err != nil {
		ctx.Logger.WithError(err).Error("put-nickname: Error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if exist {
		// ? L'utente esiste già
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ! richiamo changeNickname del DB
	err = rt.db.ChangeNickname(User{User_id: user_id}.toDataBase(), newNickname.Nickname)
	if err != nil {
		ctx.Logger.WithError(err).Error("put-nickname: Error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
