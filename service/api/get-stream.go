package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/Salvatore-1918339/wasa_project/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	profileOwnerId_str := ps.ByName("id")
	requestingUserId_str, err := extractBearerToken(r, w)

	// Controllo errore dall'estrazione del TOKEN
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(err).Error("getMyStream: Error")
		return
	}
	// Converto in Interi gli ID
	profileOwnerId, _ := strconv.Atoi(profileOwnerId_str)
	requestingUserId, _ := strconv.Atoi(requestingUserId_str)

	// ! Controllo se l'Utente può effettuare l'operazione - Non è il proprietario del profilo
	if profileOwnerId != requestingUserId {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(errors.New("you cannot do this")).Error("getMyStream: Error")
		return
	}

	// ! Prendo tutti i Followed by requestingUserId
	users_following, err := rt.db.GetFollowing(User{User_id: requestingUserId}.toDataBase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("getMyStream: Error executing GetFollowing ")
		return
	}
	fmt.Print("\nEcco gli Utenti che Segui: ", users_following)

	// ! Prendo tutte le photo degli utenti che seguo
	var stream_photos []database.Complete_Photo
	for i := 0; i < len(users_following); i++ {
		photos, err := rt.db.FindPhotos(User{User_id: users_following[i]}.toDataBase()) // Trovo tutte le photo dell'User che Seguo
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			ctx.Logger.WithError(err).Error("getMyStream: Error in FindPhotos")
			return
		}
		// Aggiungi gli elementi di photos a stream_photos
		stream_photos = append(stream_photos, photos...)

	}

	// ! Ora raccolgo tutti i commenti di ogni singola photo

	for i := 0; i < len(stream_photos); i++ {
		comments, err := rt.db.FindComment(Photo_id{Photo_id: stream_photos[i].Photo_id}.toDataBase())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("getMyStream: Error in FindComment")
			return
		}
		stream_photos[i].Comments = comments // Inseriamo i Comments Ottenuti per la foto nella Struct
	}

	// ! Ora raccolgo i Like delle singole Photo
	for i := 0; i < len(stream_photos); i++ {
		fmt.Print("stream di Photo:", stream_photos[i].Photo_id)
		likes, err := rt.db.FindLikes(Photo_id{Photo_id: stream_photos[i].Photo_id}.toDataBase())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("getMyStream: Error in FindLikes")
			return
		}
		stream_photos[i].Likes = likes
	}
	fmt.Print("\n\nSTREAM", stream_photos)
	w.WriteHeader(http.StatusOK)
	// manda l'output all'utente.
	_ = json.NewEncoder(w).Encode(stream_photos)
}
