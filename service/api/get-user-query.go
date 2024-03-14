package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

/*
	http codes che può tornare:
	- 200 : FOUND
	- 403 : Forbidden
	- 404 : Not Found
	- 500 : StatusInternalServerError
*/

func (rt *_router) searchUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	queryStr := r.URL.Query().Get("user_query_id") // Prendo il Nickname

	fmt.Print("QueryStr: ", queryStr)

	// ! Controllo che abbia fatto l'accesso
	identifier_str, err := extractBearerToken(r, w)
	if err != nil {
		ctx.Logger.WithError(err).Error("get-user-query: Authentication Error")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	identifier, _ := strconv.Atoi(identifier_str)

	// ! Ricerco nel DB
	values, err := rt.db.SearchUser(queryStr)
	if err != nil {
		ctx.Logger.WithError(err).Error("get-user-query: Error executing query")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// ! Manda l'output all'utente
	if len(values) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// ? Controllo per ogni utente se hanno Bannato l'ID dell'richiedente, se non l'hanno bannato questi vengono inseriti nella lista
	var users []User
	for i := 0; i < len(values); i++ {
		banned, err := rt.db.CheckBan(
			User{User_id: identifier}.toDataBase(),
			User{User_id: values[i].User_id}.toDataBase())
		if err != nil {
			ctx.Logger.WithError(err).Error("get-user-query: Error in checkban")
			return
		}
		if !banned {
			users = append(users, User{User_id: values[i].User_id, Nickname: values[i].Nickname})
		}
	}

	// Ritorno nell'Header STATUS OK
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(users)

}
