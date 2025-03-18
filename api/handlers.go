package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Handler handles HTTP requests
type Handler struct {
	client *Client
}

// NewHandler creates a new Handler
func NewHandler(client *Client) *Handler {
	return &Handler{client: client}
}

// respondWithJSON writes a JSON response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal Server Error"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondWithError writes an error response
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// GetPosts handles GET /api/posts
func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.client.GetPosts()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, posts)
}

// GetPost handles GET /api/posts/{id}
func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid post ID")
		return
	}

	post, err := h.client.GetPost(id)
	if err != nil {
		if strings.Contains(err.Error(), "unexpected status code: 404") {
			respondWithError(w, http.StatusNotFound, "Post not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, post)
}

// GetComments handles GET /api/comments
func (h *Handler) GetComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.client.GetComments()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, comments)
}

// GetCommentsForPost handles GET /api/posts/{id}/comments
func (h *Handler) GetCommentsForPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid post ID")
		return
	}

	comments, err := h.client.GetCommentsForPost(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, comments)
}

// GetUsers handles GET /api/users
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.client.GetUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

// GetUser handles GET /api/users/{id}
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.client.GetUser(id)
	if err != nil {
		if strings.Contains(err.Error(), "unexpected status code: 404") {
			respondWithError(w, http.StatusNotFound, "User not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

// GetAlbums handles GET /api/albums
func (h *Handler) GetAlbums(w http.ResponseWriter, r *http.Request) {
	albums, err := h.client.GetAlbums()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, albums)
}

// GetPhotos handles GET /api/photos
func (h *Handler) GetPhotos(w http.ResponseWriter, r *http.Request) {
	photos, err := h.client.GetPhotos()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, photos)
}

// GetTodos handles GET /api/todos
func (h *Handler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.client.GetTodos()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, todos)
}
