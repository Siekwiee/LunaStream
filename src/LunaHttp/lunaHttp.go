package LunaHttp

import (
	"fmt"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    http.ServeFile(w, r, "index.html")
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/event-stream")
    w.Header().Set("Cache-Control", "no-cache")
    w.Header().Set("Connection", "keep-alive")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    flusher, ok := w.(http.Flusher)
    if !ok {
        http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
        return
    }

    counter := 0
    for {
        counter++
        message := fmt.Sprintf("data: test%d\n\n", counter)
        fmt.Fprint(w, message)
        flusher.Flush()
        time.Sleep(1 * time.Second)
    }
}
