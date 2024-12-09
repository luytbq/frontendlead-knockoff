package api

import (
	"net/http"
)

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/submit":
		handleSubmission(w, r)
	case "/api/questions":
		handleGetQuestionDetail(w, r)
	default:
		notFound(w, r)
	}
}

func notFound(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "Not found", http.StatusNotFound)
}
