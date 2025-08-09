package main

import (
    "fmt"
    "net/http"
    "log"
)

func main() {
    http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain")

        fmt.Fprint(w, "REPLY TANGIN A MO")
    })

    fmt.Println("Go server listening on port 3000")

    log.Fatal(http.ListenAndServe(":3000", nil))
}
