package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) putBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	banner, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("putBan: error conversion")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	banned, err := strconv.Atoi(ps.ByName("banned_user_id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("putBan: error conversion")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Controllo se gli ID sono uguali
	if banner == banned {
		ctx.Logger.WithError(errors.New("the user cannot ban himself")).Error("putBan: error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ! Eseguo il Ban
	err = rt.db.BanUser(
		User{User_id: banner}.toDataBase(),
		User{User_id: banned}.toDataBase())

	if err != nil {
		ctx.Logger.WithError(err).Error("putBan: Error executing the query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// ? TOLGO IL FOLLOW DA ENTRAMBI
	err = rt.db.UnFollowUser(
		User{User_id: banner}.toDataBase(),
		User{User_id: banned}.toDataBase())

	if err != nil {
		ctx.Logger.WithError(err).Error("putBan: error executing the unfollow query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
