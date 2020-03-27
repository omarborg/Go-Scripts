package main

import(
  "crypto/rsa"
  "encoding/gob"
  "fmt"
  "os"
)

//Function main loads an rsa key using the loadRsaKey function.
//Prints loaded private key primes and exponent
//Prints loaded publicKey modulus and exponent
func main() {

  var key rsa.PrivateKey
  loadRsaKey("private.key", &key)

  fmt.Println("Private key primes: ", key.Primes[0].String(), key.Primes[1].String())
  fmt.Println("Private key exponent: ", key.D.String())

  var publicKey rsa.PublicKey
  loadRsaKey("public.key", &publicKey)

  fmt.Println("Public key mod: ", publicKey.N.String())
  fmt.Println("Public key exponent: ", publicKey.E)

}

//Function loads a file with rsa keys and desocdes its content
func loadRsaKey(fileName string, key interface{}) {
  inFile, err := os.Open(fileName)
  checkError(err)
  decoder := gob.NewDecoder(inFile)
  err = decoder.Decode(key)
  checkError(err)
  inFile.Close()
}

//Function checks for errors and prints them if any
func checkError(err error) {

  if err != nil {
    fmt.Println("Error", err.Error())
    os.Exit(1)
  }

}
