package main

import (
    "log"
    "net/http"
    "encoding/json"
    "io/ioutil"
)

func Write404ResponseHeader(w http.ResponseWriter, r *http.Request) { 
    http.NotFound(w, r)
}

func Write200ResponseHeader(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func SendAsJson(w http.ResponseWriter, t Todo) {
    json.NewEncoder(w).Encode(t)
}

func check(err error) {
    if err != nil {
        log.Fatalf("error: %s", err)
        panic(err)
    }
}

func checkWithMsg(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func GetFileContents(filePath string) string {
    data, err := ioutil.ReadFile(filePath)
	check(err)

    return string(data)
}
