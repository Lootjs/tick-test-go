package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  keyword := "go language"
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    if (scanner.Text() == keyword) {
       fmt.Println("phrases are equal")
    } else {
       fmt.Println("phrases are not equal")
    }
  }
}

