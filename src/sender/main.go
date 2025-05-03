package sender

import (
	"fmt"
	"github.com/siekwiee/"
)

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/stream", streamHandler)
    fmt.Println("Server running on :8080")
    http.ListenAndServe(":8080", nil)
}
