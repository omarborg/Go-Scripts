package main

import(
  "fmt"
  "net"
  "os"
  "encoding/json"
)

//Defining json structs
type User struct{
  Name UserInfo
  Email []Email
}

type UserInfo struct {
  userName string
}

type Email struct {
  emailAddress string
}

//Function to create a user using User struct
func (u User) String() string {

  userString := u.UserInfo.userName

  for _, emailData := range u.Email {
    userString += "\n" + emailData.emailAddress
  }
  return userString
}

//Function main establishes a connection to the localhost on port 1200 and listens and accepts any incoming connections,
// decodes and encodes the user info, prints the user info and closes the connection
func main() {

  address := "127.0.0.1:1200"
  tcpAddr, err := net.ResolveTCPAddr("tcp", address)
  checkError(err)

  listener, err := net.ListenTCP("tcp", tcpAddr)
  checkError(err)

  for {
    conn, err := listener.Accept()
    if err != nil {
      continue
    }

    encoder := json.NewEncoder(conn)
    decoder := json.NewDecoder(conn)

    var user User
    decoder.Decode(&user)
    fmt.Println(user.String())
    encoder.Encode(user)
  }
  conn.Close()
}

//Function to check for errors and print them if any
func checkError(err error) {

  if err != nil {
    fmt.Println("Error", err.Error())
    os.Exit(1)
  }

}

/* NOTE: same script can be used to store gob serialised data

    1- In import change "encoding/json" to "encoding/gob"
    2- In the main function change user.json to user.gob also
    3- In the laod funciton change the encoder and decoder values to gob.NewEncoder and gob.NewDecoder

 */
