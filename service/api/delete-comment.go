package api

import (
	"net/http"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	user_id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("uncomment-photo: error converting id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	comment_id, err := strconv.Atoi(ps.ByName("comment_id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("uncomment-photo: error converting comment_id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ! Login
	requestingUserId_str, err := extractBearerToken(r, w)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(err).Error("uncommentPhoto: Error")
		return
	}
	requestingUser_id, err := strconv.Atoi(requestingUserId_str)
	if err != nil {
		ctx.Logger.WithError(err).Error("comment-photo: error converting requestingUserId")
		w.WriteHeader(http.StatusBadRequest)
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
	if requestingUser_id != user_id && requestingUser_id != owner {
		ctx.Logger.WithError(err).Error("delete-comment: access denied")
		w.WriteHeader(http.StatusUnauthorized)
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
