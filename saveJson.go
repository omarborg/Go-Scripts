package main

import(
  "fmt"
  "net"
  "os"
  "encoding/json"
  "bytes"
  "io"
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

  user := User{
    Name: UserInfo{userName: "TestUser"},
    Email: []Email{Email{emailAddress: "testuser@test.com"}
      }
    }
    if len(os.Args) != 2 {
      fmt.Println("Usage: ", os.Args[0], "host:port")
      os.Exit(1)
    }
    service: os.Args[1]

    conn, err := net.Dial("tcp", service)
    checkError(err)

    encoder := json.NewEncoder(conn)
    decoder := json.NewDecoder(conn)

    encoder.Encode(user)
    var newUser User
    decoder.Decode(&newUser)
    fmt.Println(newUser.String())
    os.Exit(0)
  }
  func checkError(err error) {
    if err != nil {
      fmt.Println("Error", err.Error())
      os.Exit(1)
    }
  }
