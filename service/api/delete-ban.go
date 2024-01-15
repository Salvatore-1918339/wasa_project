package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	//prendiamo gli ID
	urlArry := strings.Split(r.URL.Path, "/")
	pathbanner, pathbanned := urlArry[2], urlArry[4]

	//Convertiamo
	banner, err := strconv.Atoi(pathbanner)
	if err != nil {
		ctx.Logger.WithError(err).Error("deleteBan: error with string conversion 1")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	banned, err2 := strconv.Atoi(pathbanned)
	if err2 != nil {
		ctx.Logger.WithError(err2).Error("deleteBan: error with string conversion 2")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Eseguiamo la query
	err = rt.db.UnBanUser(banner, banned)
	if err != nil {
		ctx.Logger.WithError(err).Error("deleteBan: Error executing the query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Print("\nUser: ", banner, " unban:", banned)

}
