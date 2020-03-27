package main

//Imports libraries
import (
  "fmt"
  "os"
  "net"
  "io/ioutils"
)

//checks the length of the input and prints a usage error in response to user input
func main() {

  if len(os.Args) != 2 {

     fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
     os.Exit(1)
  }

 //assigns service value to the appropriate user input
  service := os.Args[1]

 //Tries resolving tcp address using tcp4, the net library and checks for errors using the checkError function
  tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
  checkError(err)

 //Tries establishing a connection using tcp, the net library and checks for errors using the checkError funct
  conn, err := net.DialTCP("tcp", nil, tcpAddr)
  checkError(err)

  //Tries writing to the connection established using the conn.write and checks for errors
  -, err = conn.Write([]byte("Head / HTTP/1.0\r\n\r\n"))
  checkError(err)

  //Reads back the connection communication to the user and checks for errors
  result, err := ioutil.ReadAll(conn)
  checkError(err)

  fmt.Println(String(result))

  os.Exit(0)
}

//Function to check for errors
func checkError(err error) {

  if err != nil {
    fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
    os.Exit(1)
  }

}
