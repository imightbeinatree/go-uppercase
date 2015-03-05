package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  uppercase := strings.ToUpper(os.Args[1])
  fmt.Println("{\"string\":\""+uppercase+"\"}")
}


