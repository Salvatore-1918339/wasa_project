package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userId_str := ps.ByName("id")
	userToFollowId_str := ps.ByName("follower_id")

	// TODO - DEVI CONTROLLARE L'AUTHENTICATION

	// ? trasformo in interi
	userId, _ := strconv.Atoi(userId_str)
	userToFollowId, _ := strconv.Atoi(userToFollowId_str)

	// ? controllo se sono uguali
	if userId == userToFollowId {
		ctx.Logger.WithError(errors.New("The user cannot follow himself")).Error("followUser: Error")
		return
	}

	// ? Eseguo la query di Follow
	err := rt.db.FollowUser(
		User{User_id: userToFollowId}.toDataBase(),
		User{User_id: userId}.toDataBase())
	if err != nil {
		ctx.Logger.WithError(err).Error("put-follow: error executing insert query")
		return
	}
}
