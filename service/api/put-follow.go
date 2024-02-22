package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// TODO - DEVI CONTROLLARE L'AUTHENTICATION

	// ? trasformo in interi
	userId, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("put-follow: error conversion id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userToFollowId, _ := strconv.Atoi(ps.ByName("follower_id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("put-follow: error conversion follower_id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ! Login
	requestingUserId_str, err := extractBearerToken(r, w)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("put-follow: Error")
		return
	}
	requestingUserId, _ := strconv.Atoi(requestingUserId_str)
	if userId != requestingUserId {
		ctx.Logger.WithError(errors.New("unauthorized")).Error("put-follow: error")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// ! controllo se sono uguali
	if userId == userToFollowId {
		ctx.Logger.WithError(errors.New("the user cannot follow himself")).Error("put-follow: Error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ! Eseguo la query di Follow
	err = rt.db.FollowUser(
		User{User_id: userId}.toDataBase(),
		User{User_id: userToFollowId}.toDataBase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-follow: error executing insert query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
