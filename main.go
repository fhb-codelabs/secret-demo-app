package main

import (
    "fmt"
    "github.com/kelseyhightower/envconfig"
    "log"
    "net/http"
)


type AppConfig struct {
    Port         string
    Secret       string
}

var c AppConfig

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, your secret is %s ", c.Secret  )
}

func main() {
    err := envconfig.Process("app", &c)
    if err != nil {
      log.Fatal(err.Error())
    }

    if c.Port == "" {
        log.Println("No port is set, defaulting to 8080")
        c.Port = "8080"
    }

    if c.Secret == "" {
        log.Fatal("Secret to start this up is missing, the APP_SECRET environment variable might be unset")
    }

    http.HandleFunc("/", handler)
    log.Println("Listening on Port " + c.Port)
    log.Fatal(http.ListenAndServe(":" + c.Port, nil))
}
