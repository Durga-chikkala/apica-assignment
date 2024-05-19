package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Durga-chikkala/apica-assignment/models"
	"github.com/Durga-chikkala/apica-assignment/service"
	"github.com/Durga-chikkala/apica-assignment/sockets"
)

type Handler struct {
	s      service.LRUCache
	socket *sockets.Manager
}

func New(lruCache service.LRUCache, s *sockets.Manager) Handler {
	return Handler{s: lruCache, socket: s}
}

// Get handles HTTP GET requests for retrieving a value associated with a key. It responds with errors if the key is missing or no value is found.
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	if key == "" {
		models.SetError(w, http.StatusBadRequest, "Parameter key is required")

		return
	}

	value := h.s.Get(key)
	if value == "" {
		models.SetError(w, http.StatusNotFound, "No key found")

		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.Success{Data: value})
}

// Set handles HTTP POST/PUT requests to set a key-value pair. It responds with success or error messages based on the request body validity.
func (h *Handler) Set(w http.ResponseWriter, r *http.Request) {
	cache, err := io.ReadAll(r.Body)
	if err != nil {
		models.SetError(w, http.StatusBadRequest, "Invalid request body")

		return
	}

	var reqData models.CacheData

	err = json.Unmarshal(cache, &reqData)
	if err != nil {
		models.SetError(w, http.StatusBadRequest, "Invalid Body")

		return
	}

	if reqData.Key == "" {
		models.SetError(w, http.StatusBadRequest, "Key is required")

		return
	}

	if reqData.Value == "" {
		models.SetError(w, http.StatusBadRequest, "Value is required")

		return
	}

	if reqData.Expiration == 0 {
		models.SetError(w, http.StatusBadRequest, "time is required")

		return
	}

	cacheData := h.s.Set(&reqData)

	h.socket.Broadcast(cacheData)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.Success{Data: "Successfully inserted"})
}

// Delete handles HTTP DELETE request removes the cache entry associated with the specified key
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	if key == "" {
		models.SetError(w, http.StatusBadRequest, "Parameter key is required")

		return
	}

	cacheData := h.s.Delete(key)

	h.socket.Broadcast(cacheData)

	w.WriteHeader(http.StatusNoContent)

	json.NewEncoder(w).Encode(models.Success{Data: "Deleted successfully"})
}

// GetAllKeys handles HTTP GET request returns a map containing all keys in the cache along with their associated cache data.
func (h *Handler) GetALLKeys(w http.ResponseWriter, r *http.Request) {
	keys := h.s.GetAllKeys()
	h.socket.UpgradeHandler(w, r)
	h.socket.Broadcast(keys)
}
