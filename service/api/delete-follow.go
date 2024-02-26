package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// ? trasformo in interi
	userId, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("put-follow: error conversion id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userToUnFollowId, err := strconv.Atoi(ps.ByName("follower_id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("put-follow: error conversion follower_id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ! Login
	_, err = extractBearerToken(r, w)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(err).Error("put-follow: Error")
		return
	}

	// ? controllo se sono uguali
	if userId == userToUnFollowId {
		ctx.Logger.WithError(errors.New("the user cannot follow himself")).Error("followUser: Error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ? Eseguo la query del deleteFollow
	err = rt.db.UnFollowUser(
		User{User_id: userId}.toDataBase(),
		User{User_id: userToUnFollowId}.toDataBase())
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-follow: error executing DELETE query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
