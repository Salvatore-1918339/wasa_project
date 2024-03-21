package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	user_id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		ctx.Logger.WithError(err).Error("Post-photo: error converting id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ! Login
	requestingUserId_str, err := extractBearerToken(r, w)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(err).Error("Post-Photo: Error")
		return
	}
	requestingUserId, err := strconv.Atoi(requestingUserId_str)
	if err != nil {
		ctx.Logger.WithError(err).Error("Post-photo: error converting id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user_id != requestingUserId {
		w.WriteHeader(http.StatusUnauthorized)
		ctx.Logger.WithError(errors.New("you aren't allowed to use this operation")).Error("post-photo: Error")
		return
	}

	// Contenuto del Body
	data, err := io.ReadAll(r.Body)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-upload: error reading body content")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Controllo il Formato della Foto
	errFormatJpg := checkFormatPng(io.NopCloser(bytes.NewBuffer(data))) // passo dei Bytes da consumare
	if errFormatJpg != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(errFormatJpg).Error("photo-upload: Error not supported format")
		return
	}

	// Mi serve un ID UNIVOCO per la foto
	datetime := time.Now()
	id_user, _ := strconv.Atoi(ps.ByName("id"))

	PhotoId, err := rt.db.CreatePhoto(id_user, datetime) // Prendo l'id della Photo
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-upload: Error in CreatePhoto DB")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	strPhotoId := strconv.Itoa(PhotoId)

	// Prendo l'PATH delle PHOTO dell'User
	path := filepath.Join("/tmp", "media", ps.ByName("id"), "photos") // /tmp/media/:id/photos

	// Creo un nuovo file dentro la cartella foto del Utente
	outputFile, err := os.Create(filepath.Join(path, strPhotoId+".png"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("photo-upload: error creating local photo file")
		return
	}

	_, err = io.Copy(outputFile, io.NopCloser(bytes.NewBuffer(data)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("photo-upload: error copying body content into file photo")
		return
	}
	outputFile.Close()

	photo, err := rt.db.FindPhoto(PhotoId)
	if err != nil {
		ctx.Logger.WithError(err).Error("post-photo: error executing FindPhoto")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	photo.Comments = nil
	photo.Likes = nil
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(photo)

}

func checkFormatPng(body io.ReadCloser) error {
	//Decode consuma body. Ecco perché inseriamo body e newReader
	_, errPng := png.Decode(body)
	if errPng != nil {
		return errPng
	}
	return nil
}
