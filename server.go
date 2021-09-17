package server

import "net/http"

http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))