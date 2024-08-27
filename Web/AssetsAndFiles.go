package main

import "net/http"

func main() {
  fs := http.FileServer(http.Dir("assets/"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  http.ListenandServe(":8080", nil)
}
