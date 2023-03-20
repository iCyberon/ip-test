package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "IP: %q", getip2())
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}

type IP struct {
    Query string
}

func getip2() string {
    req, err := http.Get("http://ip-api.com/json/")
    if err != nil {
        return err.Error()
    }
    defer req.Body.Close()

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        return err.Error()
    }

    var ip IP
    json.Unmarshal(body, &ip)

    return ip.Query
}
