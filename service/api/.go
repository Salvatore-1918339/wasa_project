package api

import (
	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// getContextReply is an example of HTTP endpoint that returns "Hello World!" as a plain text. The signature of this
// handler accepts a reqcontext.RequestContext (see httpRouterHandler).
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "text/plain")	//Dico al client che La risposta che riceverà sarà del tipo ("content-type", "text/plain")
	_, _ = w.Write([]byte("Hello World!"))
	
	/*
	var user User
	err = rt.db.CreateUser(user.ToDatabase()) //Proviamo a creare un Utente
	if err != nil {

	}
	*/
	
}

