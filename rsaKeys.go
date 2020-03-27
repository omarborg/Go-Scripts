package main

import(
  "crypto/rand"
  "crypto/rsa"
  "crypto/x509"
  "encoding/gob"
  "encoding/pem"
  "fmt"
  "os"
)

//Function main generates an rsa key with a size of 512 bytes, chekcs for errors.
//Prints genrated private key primes and exponent
//Prints genrated publicKey modulus and exponent
//Saves the private and public key to an outFile using the saveGobKey function
func main() {

  reader := rand.Reader
  bitSize := 512
  key, err := rsa.GenerateKey(reader, bitSize)
  checkError(err)

  fmt.Println("Private key primes: ", key.Primes[0].String(), key.Primes[1].String())
  fmt.Println("Private key exponent: ", key.D.String())

  publicKey := key.PublicKey
  fmt.Println("Public key mod: ", publicKey.N.String())
  fmt.Println("Public key exponent: ", publicKey.E)

  saveGobKey("private.key", key)
  saveGobKey("public.key", publicKey)

  savePemKey("private.pem", key)
}

//Function creates a file and encodes its content
func saveGobKey(fileName string, key interface{}) {
  outFile, err := os.Create(fileName)
  checkError(err)
  encoder := gob.NewEncoder(outFile)
  err = encoder.Encode(key)
  checkError(err)
  outFile.Close()
}

//Function creates a file and encodes its content
func savePemKey(fileName string, key *rsa.PrivateKey) {
  outFile, err := os.Create(fileName)
  checkError(err)

  var privateKey = &pem.Block{Type: "RSA Private Key",
      Bytes: x509.MarshalPKCS1PrivateKey(key)}

  pem.Encode(outFile, privateKey)

  outFile.Close()
}

//Function checks for errors and prints them if any
func checkError(err error) {

  if err != nil {
    fmt.Println("Error", err.Error())
    os.Exit(1)
  }

}
