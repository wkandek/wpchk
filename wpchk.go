package main

import (
  "crypto/sha256"
  "fmt"
  "io"
  "log"
  "net/http"
  "strings"
  _ "github.com/breml/rootcerts"
)

func pageCheckSum(url string) {
  // Make HTTP GET request
  response, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  defer response.Body.Close()

  b, err := io.ReadAll(response.Body)
  sha256Sum := sha256.Sum256([]byte(b))
  checksum := fmt.Sprintf("%x\n", sha256Sum)

  logurl := "http://msg2.kandek.com/log?" + url + "_" + checksum
  logurl = strings.TrimSuffix(logurl,"\n")
  log.Println("Checksum:", logurl)
  logresponse, err := http.Get(logurl)
  if err != nil {
    log.Fatal(err)
  }
  defer logresponse.Body.Close()
  b, err = io.ReadAll(logresponse.Body)
}


func main() {
  urlArray := [3]string{"https://en.wikipedia.org","https://el.wikipedia.org","https://de.wikipedia.org"}
  for index, element := range urlArray {
    fmt.Println(index, "=>", element)
    pageCheckSum(element)
  }
}
