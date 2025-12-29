package main

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type URL struct{
	Code string
	LongUrl string
	CreatedAt time.Time
}

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	Code string `json:"code"`
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var url_setups []URL

func shortenHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "content type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if err := validateURL(req.URL); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	length := 6
	var sb strings.Builder
	sb.Grow(length)

	for i := 0; i < length; i++ {
		idx := rand.Intn(len(charset))
		sb.WriteByte(charset[idx]) 
	}

	url_setups = append(url_setups, URL{
		Code: sb.String(),
		LongUrl : req.URL,
	})

	resp := shortenResponse{Code:sb.String()}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(resp)
}

func validateURL(u string) error {
	if u == "" {
		return errors.New("url is required")
	}
	return nil
}

func expandHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	code := r.URL.Query().Get("code")
	if code == ""{
		http.Error(w, "code is required", http.StatusBadRequest)
		return
	}

	for _, url := range(url_setups){
		if url.Code == code{
			w.Write([]byte(url.LongUrl))
			return
		}
	}

	http.Error(w, "url not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/shorten",shortenHandler)
	http.HandleFunc("/expand",expandHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}