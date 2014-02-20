package main

import (
  "net/http"
  "code.google.com/p/go.net/websocket"
  "time"
  "math/rand"
  "runtime"
  "./game"
)
//godoc.org/code.google.com/p/go.net/websocket

func serveFile(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "index.html")
}

func main() {
  rand.Seed( time.Now().UTC().UnixNano() )
  runtime.GOMAXPROCS( runtime.NumCPU() )

  http.HandleFunc("/", serveFile)
  http.Handle("/socket/", websocket.Handler(game.Game) )
  http.ListenAndServe(":8080", nil)
}