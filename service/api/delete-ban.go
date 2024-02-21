package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	banner, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-ban: error with string conversion id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user_id, err := strconv.Atoi(ps.ByName("banned_user_id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-ban: error with string conversion banned_user_id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	requestingUserId_str, err := extractBearerToken(r, w)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("delete-ban: Error extract token")
		return
	}
	requestingUserId, _ := strconv.Atoi(requestingUserId_str)

	// ! Login
	if banner != requestingUserId {
		ctx.Logger.WithError(errors.New("unauthorized")).Error("delete-ban: error")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	//Eseguiamo la query
	err = rt.db.UnBanUser(banner, user_id)
	if err != nil {
		ctx.Logger.WithError(err).Error("deleteBan: Error executing the query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
