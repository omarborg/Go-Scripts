package main

import(
  "fmt"
  "os"
  "net"
)


//Function main uses localhost port 1300 to listen and accpet connections
//uses a go routine to handle client connections
func main() {

  address := "127.0.0.1:1300"
  tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
  checkError(err)

  listener, err := net.ListenTCP("tcp", tcpAddr)
  checkError(err)

  for {
    conn, err := listener.Accept()
    if err != nil {
      continue
    }

    go handleClient(conn)
  }
}

//Function handleClient closes connection on exit and reads a buffer of max 512 bytes and writes that buffer back
func handleClient(conn net.Conn) {

  defer conn.Close()

  var buffer [512] byte
  for {

    nBytes , err := conn.Read(buffer[0:])
    if err != nil {
      return
    }

    _, err2 := conn.Write(buffer[0:nBytes])
    if err2 != nil {
      return
    }
  }
}

//Function to check errors and print if any
func checkError(err error){
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
    os.Exit(1)
  }
}
