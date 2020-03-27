package main

import(
  "crypto/md5"
  "fmt"
  "io"
  "log"
  "os"
)

// Function main opens a txt file and checks if the file exists if not logs an errors.
// Creates a new md5 hash and copies the hash to the file and logs any errors
// Prints the hash sum and closes the file
func main() {

  file, err := os.Open("filename.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  hash := md5.new()

  if _,err := io.Copy(hash, file); err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%x", hash.sum(nil))
}
