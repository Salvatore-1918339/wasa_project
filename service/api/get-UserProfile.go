package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var profile Profile
	user_id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("get-UserProfile : error conversion id")
		return
	}
	profile.User_id = user_id

	// ! Login
	_, err = extractBearerToken(r, w)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(err).Error("get-UserProfile: Error during login")
		return
	}

	// ! Getnickname
	nickname, err := rt.db.GetNickname(profile.User_id)
	if err != nil {
		ctx.Logger.WithError(err).Error("get-UserProfile : Error executing the getNickname query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	profile.Nickname = nickname

	// ! Followers number
	num_follower, err := rt.db.GetNumberFollower(profile.User_id)
	if err != nil && err != sql.ErrNoRows {
		ctx.Logger.WithError(err).Error("get-UserProfile : error executing the number follower Query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	profile.Follower = num_follower

	// ! Followed number
	num_followed, err := rt.db.GetNumberFollowed(profile.User_id)
	if err != nil && err != sql.ErrNoRows {
		ctx.Logger.WithError(err).Error("get-UserProfile : error executing the number follower Query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	profile.Following = num_followed

	// ! List of Photo
	profile_photos, err := rt.db.FindPhotos(User{User_id: user_id}.toDataBase())
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(err).Error("getMyStream: Error in FindPhotos")
		return
	}
	profile.Posts = profile_photos
	w.WriteHeader(http.StatusOK)
	// manda l'output all'utente.
	_ = json.NewEncoder(w).Encode(profile)

}
