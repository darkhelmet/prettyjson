package main

import (
    "encoding/json"
    "flag"
    "log"
    "os"
    "strings"
)

var indent = flag.Int("indent", 4, "the number of spaces to indent")

func main() {
    flag.Parse()
    decoder := json.NewDecoder(os.Stdin)
    var v interface{}
    err := decoder.Decode(&v)
    if err != nil {
        log.Fatalf("failed decoding: %s", err)
    }
    data, err := json.MarshalIndent(v, "", strings.Repeat(" ", *indent))
    if err != nil {
        log.Fatalf("failed marshalling: %s", err)
    }
    os.Stdout.Write(data)
    os.Stdout.Write([]byte{'\n'})
}
