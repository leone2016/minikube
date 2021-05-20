package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
)


type HandsOn struct{
    Time time.Time `json:"tiempo"`
    HostName string `json:"host_name"`
}
type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    resp := HandsOn{
        Time:   time.Now(),
        HostName: os.Getenv("HOSTNAME"),

    }
    jsonRes, err := json.Marshal(&resp)
    if err != nil {
        fmt.Println("error: ",err)
    }
    w.Write([]byte(jsonRes))
}

func main() {
    s := &server{}
    http.Handle("/", s)
    log.Fatal(http.ListenAndServe(":9092", nil))
}
