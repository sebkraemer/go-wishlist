package handlers

import (
	"net/http"
)

func SessionLoadAndSave(next http.Handler) http.Handler {
	return Repo.App.Session.LoadAndSave(next)
}
