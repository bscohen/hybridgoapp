// +build !appengine

package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.ListenAndServe("localhost:8080", nil)
}