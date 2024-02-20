package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userId_str := ps.ByName("id")
	userToUnFollowId_str := ps.ByName("follower_id")

	// ? trasformo in interi
	userId, _ := strconv.Atoi(userId_str)
	userToUnFollowId, _ := strconv.Atoi(userToUnFollowId_str)

	// ? controllo se sono uguali
	if userId == userToUnFollowId {
		ctx.Logger.WithError(errors.New("the user cannot follow himself")).Error("followUser: Error")
		return
	}

	// ? Eseguo la query del deleteFollow
	err := rt.db.UnFollowUser(
		User{User_id: userId}.toDataBase(),
		User{User_id: userToUnFollowId}.toDataBase())
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-follow: error executing DELETE query")
		return
	}
}
