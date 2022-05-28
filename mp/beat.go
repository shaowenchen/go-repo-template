package mp

import (
	"net/http"
	"time"
)

func BeatVqdata() {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		http.Get("https://vqdata.chenshaowen.com/")
	}
}
