package main 

import (
	"fmt"
	"LunaStream/src/LunaHttp"
	"net/http"
)

func main() {
    http.HandleFunc("/", LunaHttp.HomeHandler)
    http.HandleFunc("/stream", LunaHttp.StreamHandler)
    fmt.Println("Server running on :8080")
    http.ListenAndServe(":8080", nil)
}
