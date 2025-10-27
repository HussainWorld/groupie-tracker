package main

import (
    "log"
    "net/http"
)

type ErrorView struct {
    StatusCode int
    Error      string
}

func renderError(w http.ResponseWriter, status int, publicMsg string, logErr error) {
    if logErr != nil {
        log.Printf("HTTP %d: %s | err=%v", status, publicMsg, logErr)
    } else {
        log.Printf("HTTP %d: %s", status, publicMsg)
    }

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.WriteHeader(status)

    _ = tmpl.ExecuteTemplate(w, "error.html", ErrorView{
        StatusCode: status,
        Error:      publicMsg,
    })
}
