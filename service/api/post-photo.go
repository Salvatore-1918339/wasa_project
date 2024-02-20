package api

import (
	"bytes"
	"errors"
	"fmt"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Salvatore-1918339/wasa_project/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) PostPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("Content-Type", "application/json")

	user_id_str := ps.ByName("id")
	requestingUserId_str, err := extractBearerToken(r, w)
	// Controllo errore dall'estrazione del TOKEN
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.WithError(err).Error("GetNIcknameHandler: Error")
		return
	}

	user_id, _ := strconv.Atoi(user_id_str)
	requestingUserId, _ := strconv.Atoi(requestingUserId_str)

	// ! check Login
	if user_id != requestingUserId {
		w.WriteHeader(http.StatusForbidden) // 403
		ctx.Logger.WithError(errors.New("you aren't allowed to use this operation")).Error("post-photo: Error")
		return
	}

	// Contenuto del Body
	data, err := io.ReadAll(r.Body)
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-upload: error reading body content")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Controllo il Formato della Foto
	errFormatJpg := checkFormatJpg(io.NopCloser(bytes.NewBuffer(data))) //passo dei Bytes da consumare
	if errFormatJpg != nil {
		ctx.Logger.WithError(errFormatJpg).Error("photo-upload: Error not supported format")
	}

	// Mi serve un ID UNIVOCO per la foto
	currentTime := time.Now()
	datetime := currentTime.Format("2006-01-02 15:04:05")
	id_user, _ := strconv.Atoi(ps.ByName("id"))

	PhotoId, err := rt.db.CreatePhoto(id_user, datetime) //Prendo l'id della Photo
	if err != nil {
		ctx.Logger.WithError(err).Error("photo-upload: Error in CreatePhoto DB")
		return
	}

	strPhotoId := strconv.Itoa(PhotoId)

	// Prendo l'PATH delle PHOTO dell'User
	path := filepath.Join("/tmp", "media", ps.ByName("id"), "photos") // /tmp/media/:id/photos
	fmt.Print(filepath.Join("\n\nUpload photo: ", path, strPhotoId+".jpg"))
	// Creo un nuovo file dentro la cartella foto del Utente
	outputFile, err := os.Create(filepath.Join(path, strPhotoId+".jpg"))
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

}

func checkFormatJpg(body io.ReadCloser) error {
	//Decode consuma body. Ecco perché inseriamo body e newReader
	_, errJpg := jpeg.Decode(body)
	if errJpg != nil {
		return errJpg
	}
	return nil
}
