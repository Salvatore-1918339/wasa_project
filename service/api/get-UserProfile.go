package api

import (
	"database/sql"
	"encoding/json"
	"errors"
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
	requestingUserId_str, err := extractBearerToken(r, w)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(err).Error("get-UserProfile: Error during login")
		return
	}
	requestingUserId, err := strconv.Atoi(requestingUserId_str)
	if err != nil {
		ctx.Logger.WithError(err).Error("get-UserProfile: error converting requesingUserId")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ! Controllo se bannato
	banned, err := rt.db.CheckBan(
		User{User_id: requestingUserId}.toDataBase(),
		User{User_id: user_id}.toDataBase())
	if err != nil {
		ctx.Logger.WithError(err).Error("get-photo: Error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		ctx.Logger.WithError(errors.New("l'utente è bloccato")).Error("get-profile: Impossibile eseguire l'operazione")
		w.WriteHeader(http.StatusPartialContent)
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

	// ! Followers of the Users
	follower, err := rt.db.GetFollower(User{User_id: user_id}.toDataBase())
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		ctx.Logger.WithError(err).Error("get-UserProfile : error executing the follower Query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	profile.Followers = follower
	profile.N_Follower = len(profile.Followers)

	// ! Followed number
	num_followed, err := rt.db.GetNumberFollowed(profile.User_id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
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
