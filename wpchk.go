package main

import (
  "crypto/sha256"
  "fmt"
  "io"
  "log"
  "net/http"
  _ "github.com/breml/rootcerts"
)

// root certs import nmakes the binary independent from the OS CA certificate file

func pageCheckSum(url string) {
  // Make HTTP GET request
  response, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  defer response.Body.Close()

  // checksum the bytes in the body
  b, err := io.ReadAll(response.Body)
  sha256Sum := sha256.Sum256([]byte(b))
  checksum := fmt.Sprintf("%x", sha256Sum)

  // post to log website
  logurl := "http://msg2.kandek.com/log?" + url + "_" + checksum
  log.Println("Checksum:", logurl)
  response, err = http.Get(logurl)
  if err != nil {
    log.Fatal(err)
  }
  defer response.Body.Close()
}


func main() {
  urlArray := [3]string{"https://en.wikipedia.org","https://el.wikipedia.org","https://de.wikipedia.org"}
  for index, element := range urlArray {
    fmt.Println(index, "=>", element)
    pageCheckSum(element)
  }
}
