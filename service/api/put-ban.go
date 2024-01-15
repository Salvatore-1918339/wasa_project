package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) putBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// prendiamo gli ID
	// ! TODO - Cambia questa parte e falla più elegante

	banner_str := ps.ByName("id")
	banned_str := ps.ByName("banned_user_id")

	// Convertiamo
	banner, _ := strconv.Atoi(banner_str)
	banned, _ := strconv.Atoi(banned_str)

	fmt.Print("Banner: ", banner, " banned:", banned)

	// Controllo se gli ID sono uguali
	if banner == banned {
		ctx.Logger.WithError(errors.New("The user cannot ban himself")).Error("putBan: error")
		return
	}

	// ESEGUO IL BAN
	err := rt.db.BanUser(
		User{User_id: banner}.toDataBase(),
		User{User_id: banned}.toDataBase())

	if err != nil {
		ctx.Logger.WithError(err).Error("putBan: Error executing the query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Print("\nUser:", banner, " banned: ", banned)

	// ? TOLGO IL FOLLOW DA ENTRAMBI
	err = rt.db.UnFollowUser(
		User{User_id: banner}.toDataBase(),
		User{User_id: banned}.toDataBase())

	if err != nil {
		ctx.Logger.WithError(err).Error("putBan: error executing the unfollow query")
		return
	}

}
