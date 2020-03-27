package main

import(
  "fmt"
  "net"
  "os"
  "encoding/json"
)

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

func (u User) String() string {

  userString := u.UserInfo.userName

  for _, emailData := range u.Email {
    userString += "\n" + emailData.emailAddress
  }
  return userString
}

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

func checkError(err error) {
  if err != nil {
    fmt.Println("Error", err.Error())
    os.Exit(1)
  }
}
