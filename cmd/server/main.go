package main

import "net/http"

func main() { http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) { w.Write([]byte(`{"status":"ok"}`)) }); http.HandleFunc("/matches", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusAccepted) }); _ = http.ListenAndServe(":8084", nil) }
