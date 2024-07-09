package main

import (
	"math/rand"
	"net/http"
)

type URLShortener struct {
	urls map[string]string
}

func (object *URLShortener) HandleShorten(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "invalid request method", http.StatusMethodNotAllowed)
	}

	originalUrl := request.FormValue("url")
	if originalUrl == "" {
		http.Error(writer, "URL Parameter missing", http.StatusBadRequest)
		return
	}

	shortKey := generateShortKey()
	object.urls[shortKey] = originalUrl
}

func generateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyz1234567890"
	const keyLength = 6

	// New(NewSource(seed))
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}
