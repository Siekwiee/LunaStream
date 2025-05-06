package LunaHttp

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"bytes"
	"image/jpeg"
	"github.com/kbinani/screenshot"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    w.Header().Set("Content-Type", "text/html")
    http.ServeFile(w, r, "index.html")
}

func StreamHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=frame")

    flusher, ok := w.(http.Flusher)
    if !ok {
        http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
        return
    }

    counter := 0
    for {
				img, err := screenshot.CaptureDisplay(0)
				if err != nil {
			     log.Printf("Error captureing screen_ %v", err)
					 continue
				}

				// Encode to JPEG
        var buf bytes.Buffer
		err = jpeg.Encode(&buf, img, &jpeg.Options{Quality:25})
        if err != nil {
            log.Printf("Error encoding image: %v", err)
            continue
        }

			  _, err = fmt.Fprintf(w, "--frame\r\nContent-Type: image/jpeg\r\nContent-Length: %d\r\n\r\n", buf.Len())
        if err != nil {
            return // Client disconnected
        }
				        // Write the actual JPEG image data
        _, err = w.Write(buf.Bytes())
        if err != nil {
            return // Client disconnected
        }

        counter++
        message := fmt.Sprintf("data: test%d\n\n", counter)
        fmt.Fprint(w, message)
        flusher.Flush()
        time.Sleep(25 * time.Millisecond)
    }
}
