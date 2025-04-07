package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Query().Get("name")
        fmt.Fprintf(w, "Привет, %s!", name)
    })

    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Println("Сервер запущен на http://localhost:" + port)
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        fmt.Println("Ошибка запуска сервера:", err)
        os.Exit(1)
    }
}
