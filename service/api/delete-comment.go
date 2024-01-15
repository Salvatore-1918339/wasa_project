package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	user_id_str := ps.ByName("id")
	photo_id_str := ps.ByName("photo_id")
	comment_id_str := ps.ByName("comment_id")
	requestingUserId_str, err := extractBearerToken(r, w)

	// Controllo errore dall'estrazione del TOKEN
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(err).Error("uncommentPhoto: Error")
		return
	}

	// Trasformo in interi
	user_id, _ := strconv.Atoi(user_id_str)
	photo_id, _ := strconv.Atoi(photo_id_str)
	comment_id, _ := strconv.Atoi(comment_id_str)
	requestingUser_id, _ := strconv.Atoi(requestingUserId_str)

	// ! Controllo se l'Utente è bannato
	banned, err := rt.db.CheckBan(
		User{User_id: requestingUser_id}.toDataBase(), // ? User che potrebbe essere bannato
		User{User_id: photo_id}.toDataBase())          // ? Owner della photo

	if banned {
		ctx.Logger.WithError(errors.New("L'utente è bloccato")).Error("Impossibile eseguire l'operazione")
		w.WriteHeader(http.StatusForbidden) //errore 403
		return
	}

	// ! CONTROLLO - Il commento può essere tolto solo dall'owner della photo oppure dal creatore del commento
	owner, err := rt.db.FindCommentOwner(Comment_id{Comment_id: comment_id}.toDataBase())
	if err != nil {
		ctx.Logger.WithError(err).Error("delete-comment: Error in FindCommentOwner")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// TODO - RICONTROLLA QUESTO IF QUI SOTTO: E' SBAGLIATO D:
	if requestingUser_id != user_id || requestingUser_id != owner {
		ctx.Logger.WithError(err).Error("delete-comment: access denied")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// ! Eseguo la query per eliminare il commento
	err = rt.db.DeleteComment(Comment_id{Comment_id: comment_id}.toDataBase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("delete-comment: Impossibile eseguire la query")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
