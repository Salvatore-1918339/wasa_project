package api

import (
	"errors"
	"net/http"
	"strings"
)

// verifica se l'id di un utente ha la lunghezza giusta
func validIdentifier(identifier string) bool {
	var nospace_id = strings.ReplaceAll(identifier, " ", "")
	//5-25 come def in API
	return len(identifier) >= 5 && len(identifier) <= 25 && nospace_id != "" && !strings.ContainsAny(nospace_id, "?")
}

func validCommentString(comment Comment) bool {
	return len(comment.Comment_string) >= 1 && len(comment.Comment_string) <= 300
}

// ! estrazione del TOKEN di Autenticazione
func extractBearerToken(req *http.Request, w http.ResponseWriter) (string, error) {
	// Ottieni l'intestazione Authorization dalla richiesta
	authHeader := req.Header.Get("Authorization")

	// Verifica se l'intestazione Authorization è vuota o mancante
	if authHeader == "" {
		w.WriteHeader(http.StatusForbidden)
		return "", errors.New("Intestazione Authorization mancante")
	}

	// Verifica se l'intestazione inizia con "Bearer "
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("Formato di Authorization non valido. Deve iniziare con 'Bearer '")
	}

	// Estrai il token Bearer dalla stringa dell'intestazione
	token := strings.TrimPrefix(authHeader, "Bearer ")

	return token, nil
}
