package controllers

import (
	"encoding/json"
	"go-memcached/models"
	"net/http"
	"strconv"
	"time"
)

func BlogShow(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/blogs/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid blog id", http.StatusBadRequest)
		return
	}

	// Check cache

	data := models.CacheData("blog:"+idStr, 60, func() []byte {
		blog := models.BlogsFind(uint64(id))
		blogBytes, _ := json.Marshal(blog)

		// If not in cache, cache the result

		// Simulate delay
		time.Sleep(2 * time.Second)

		return blogBytes
	}())

	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(data); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
